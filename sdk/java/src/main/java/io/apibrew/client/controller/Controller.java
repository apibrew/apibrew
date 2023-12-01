package io.apibrew.client.controller;

import io.apibrew.client.Client;
import io.apibrew.client.Config;
import io.apibrew.client.Entity;
import io.apibrew.client.Repository;
import io.apibrew.client.ext.Condition;
import io.apibrew.client.ext.Handler;
import io.apibrew.client.ext.impl.PollerExtensionService;
import io.apibrew.client.controller.model.ControllerInstance;
import lombok.extern.log4j.Log4j2;

import java.io.IOException;
import java.util.HashMap;
import java.util.Map;
import java.util.function.BiFunction;

@Log4j2
public class Controller<T extends Entity & ControllerInstance> {

    private final Class<T> instanceClass;

    private final BiFunction<Client, T, InstanceClient> newInstanceClient;
    private final String name;

    public Controller(String name, Class<T> instanceClass, BiFunction<Client, T, InstanceClient> newInstanceClient) {
        this.name = name;
        this.instanceClass = instanceClass;
        this.newInstanceClient = newInstanceClient;
    }

    private final Map<String, InstanceClient> instanceMap = new HashMap<>();

    public void startUpController(T.ServerConfig controller) throws IOException {
        log.info("Starting controller: " + controller.getHost());

        io.apibrew.client.Config.Server server = prepareConfigServer(controller);

        Client client = Client.newClientByServerConfig(server);

        Repository<T> instancesRepository = client.repository(instanceClass);

        log.info("Starting controller instances");
        instancesRepository.list().getContent().forEach(this::startUpInstance);

        log.info("Starting controller instance listener");

        PollerExtensionService extService = new PollerExtensionService(this.name + "-instance-poller", client, "storage-instance-poller");

        log.info("Started controller: " + controller.getHost());

        Handler<T> ControllerInstanceHandler = extService.handler(instanceClass);

        ControllerInstanceHandler.when(Condition.afterCreate())
                .when(Condition.async())
                .operate((event, instance) -> {
                    log.info("Creating instance: " + instance.getName());
                    startUpInstance(instance);
                    log.info("Created instance: " + instance.getName());

                    return instance;
                });

        ControllerInstanceHandler.when(Condition.afterUpdate())
                .when(Condition.async())
                .operate((event, instance) -> {
                    log.info("Updating instance: " + instance.getName());
                    destroyInstance(instance);
                    startUpInstance(instance);
                    log.info("Updated instance: " + instance.getName());

                    return instance;
                });

        ControllerInstanceHandler.when(Condition.beforeDelete())
                .when(Condition.async())
                .operate((event, instance) -> {
                    log.info("Deleting instance: " + instance.getName());
                    destroyInstance(instance);
                    log.info("Deleted instance: " + instance.getName());

                    return instance;
                });

        extService.run();
    }

    public void destroyInstance(T instance) {
        if (!instanceMap.containsKey(instance.getName())) {
            log.error("Instance not started: " + instance.getName());
            return;
        }

        instanceMap.get(instance.getName()).stop();
        instanceMap.remove(instance.getName());
    }

    private io.apibrew.client.Config.Server prepareConfigServer(T.ServerConfig controller) {
        io.apibrew.client.Config.Server controllerConfig = new io.apibrew.client.Config.Server();
        controllerConfig.setHost(controller.getHost());
        controllerConfig.setPort(controller.getPort());
        controllerConfig.setHttpPort(controller.getHttpPort());
        controllerConfig.setInsecure(controller.getInsecure());
        io.apibrew.client.Config.Authentication authentication = new io.apibrew.client.Config.Authentication();
        authentication.setUsername(controller.getAuthentication().getUsername());
        authentication.setPassword(controller.getAuthentication().getPassword());
        authentication.setToken(controller.getAuthentication().getToken());
        controllerConfig.setAuthentication(authentication);
        return controllerConfig;
    }

    public void startUpInstance(T instance) {
        if (instanceMap.containsKey(instance.getName())) {
            log.error("Instance already started: " + instance.getName());
            return;
        }

        log.info("Starting instance: " + instance.getName());

        Thread thread = new Thread(() -> {
            for (int i = 0; i < 100; i++) {
                try {
                    Client client = prepareClient(instance);

                    InstanceClient instanceClient = newInstanceClient.apply(client, instance);
                    instanceClient.init();
                    instanceMap.put(instance.getName(), instanceClient);
                    break;
                } catch (Exception e) {
                    log.error("Unable to start instance: " + instance.getName(), e);
                    try {
                        Thread.sleep(1000 * (i + 1));
                    } catch (InterruptedException interruptedException) {
                        interruptedException.printStackTrace();
                    }
                }
            }

            log.info("Started instance: " + instance.getName());
        });

        thread.setName("storage-instance-startup[" + instance.getName() + "]");

        thread.start();
    }

    private Client prepareClient(T instance) {
        Config.Server serverConfig = prepareConfigServer(instance.getServerConfig());

        return Client.newClientByServerConfig(serverConfig);
    }
}
