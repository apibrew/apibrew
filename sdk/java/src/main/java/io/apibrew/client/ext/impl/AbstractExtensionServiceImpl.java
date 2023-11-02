package io.apibrew.client.ext.impl;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import io.apibrew.client.Client;
import io.apibrew.client.Entity;
import io.apibrew.client.EntityInfo;
import io.apibrew.client.Repository;
import io.apibrew.client.ExtensionInfo;
import io.apibrew.client.ext.ExtensionService;
import io.apibrew.client.ext.GenericHandler;
import io.apibrew.client.ext.Handler;
import io.apibrew.client.impl.GenericHandlerImpl;
import io.apibrew.client.model.Extension;
import io.apibrew.client.model.Record;
import lombok.extern.log4j.Log4j2;

import java.io.IOException;
import java.util.*;
import java.util.function.BiFunction;

import static io.apibrew.client.helper.EventHelper.shortInfo;

@Log4j2
public abstract class AbstractExtensionServiceImpl implements ExtensionService {
    protected final Client client;
    private final Set<ExtensionInfo> extensionInfoSet = new HashSet<>();
    private final Set<ExtensionInfo> registeredExtensionInfoSet = new HashSet<>();
    private final Map<String, ExtensionInfo> extensionInfoIdMap = new HashMap<>();
    private final Map<ExtensionInfo, List<BiFunction<Extension.Event, Record, Record>>> extensionHandlerMap = new HashMap<>();
    private final Repository<Extension> extensionRepo;

    protected final ObjectMapper objectMapper = new ObjectMapper();
    private final String serviceName;
    Map<String, BiFunction<Extension.Event, Record, Record>> operatorMap = new HashMap<>();
    Map<String, ExtensionInfo> operatorIdExtensionInfoMap = new HashMap<>();


    protected AbstractExtensionServiceImpl(String serviceName, Client client) {
        this.serviceName = serviceName;
        this.client = client;
        this.extensionRepo = client.repo(Extension.class);

        objectMapper.registerModule(new JavaTimeModule());
    }

    @Override
    public <T extends Entity> Handler<T> handler(Class<T> entityClass) {
        return new HandlerImpl<T>(client, this, entityClass);
    }

    @Override
    public <T extends Entity> Handler<T> handler(EntityInfo<T> entityInfo) {
        return new HandlerImpl<T>(client, this, entityInfo);
    }

    @Override
    public GenericHandler genericHandler() {
        return new GenericHandlerImpl(client, this);
    }

    protected synchronized void registerExtensions() {
        log.debug("ExtensionService: {} / Registering extensions", serviceName);

        extensionInfoSet.forEach(extensionInfo -> {
            if (!registeredExtensionInfoSet.contains(extensionInfo)) {
                registeredExtensionInfoSet.add(extensionInfo);

                registerExtension(extensionInfo);
            }
        });

        log.debug("ExtensionService: {} / Registered extensions", serviceName);
    }

    private void registerExtension(ExtensionInfo extensionInfo) {
        log.debug("ExtensionService: {} / Registering extension: {}", serviceName, extensionInfo);

        Extension extension = extensionInfo.toExtension();
        extension.setCall(prepareExternalCall());
        extension.setName(serviceName + "/" + extension.getName());
        extension = extensionRepo.apply(extension);

        extensionInfoIdMap.put(extension.getId().toString(), extensionInfo);

        log.debug("ExtensionService: {} / Registered extension: {}", serviceName, extensionInfo);
    }

    protected abstract Extension.ExternalCall prepareExternalCall();

    protected Extension.Event processEvent(Extension.Event event) {
        log.debug("ExtensionService: {} / Begin processing event: {}", serviceName, shortInfo(event));
        // normalize event

        if (event.getAnnotations() == null) {
            event.setAnnotations(new HashMap<>());
        }

        String extensionId = event.getAnnotations().get("ExtensionId");
        ExtensionInfo extensionInfo = extensionInfoIdMap.get(extensionId);

        log.trace("ExtensionService: {} / Event ID: {} => Extension ID: {}", serviceName, event.getId(), extensionId);
        log.trace("ExtensionService: {} / ExtensionInfo: {}", serviceName, extensionInfo);

        if (extensionInfo == null) {
            log.warn("ExtensionInfo not found for event: {}", shortInfo(event));
            return event;
        }

        Extension.Event eventChain = processEvent(extensionInfo, event);

        log.debug("ExtensionService: {} / End processing event: {}", serviceName, shortInfo(event));

        return eventChain;
    }

    private Extension.Event processEvent(ExtensionInfo extensionInfo, Extension.Event eventChain) {
        List<BiFunction<Extension.Event, Record, Record>> handlers = extensionHandlerMap.get(extensionInfo);

        if (handlers != null) {
            for (BiFunction<Extension.Event, Record, Record> handler : handlers) {
                List<Record> records = eventChain.getRecords();
                boolean handlerHandled = false;
                if (records != null) {
                    List<Record> eventChainRecords = eventChain.getRecords();
                    List<Record> processedRecords = new ArrayList<>();

                    for (Record record : eventChainRecords) {
                        log.debug("ExtensionService: {} / Processing record: {}", serviceName, record.getId());
                        Record processedRecord = handler.apply(eventChain, record);
                        if (processedRecord != null) {
                            processedRecords.add(processedRecord);
                        }
                        handlerHandled = true;
                    }

                    eventChain.setRecords(processedRecords);
                }

                if (!handlerHandled) {
                    handler.apply(eventChain, null);
                }


            }
        }

        return eventChain;
    }

    @Override
    public String registerExtensionWithOperator(ExtensionInfo extensionInfo, BiFunction<Extension.Event, Record, Record> operator) {
        String id = UUID.randomUUID().toString();
        extensionInfoSet.add(extensionInfo);

        extensionHandlerMap.putIfAbsent(extensionInfo, new ArrayList<>());

        extensionHandlerMap.get(extensionInfo).add(operator);

        operatorMap.put(id, operator);
        operatorIdExtensionInfoMap.put(id, extensionInfo);

        log.info("ExtensionService: {} / Registered operator: {} => {}", serviceName, id, extensionInfo);

        return id;
    }

    @Override
    public void unRegisterOperator(String id) {
        BiFunction<Extension.Event, Record, Record> operator = operatorMap.get(id);

        if (operator == null) {
            throw new RuntimeException("Operator not found for id: " + id);
        }

        ExtensionInfo extensionInfo = operatorIdExtensionInfoMap.get(id);
        extensionHandlerMap.get(extensionInfo).remove(operator);

        if (extensionHandlerMap.get(extensionInfo).isEmpty()) {
            extensionHandlerMap.remove(extensionInfo);
            extensionInfoSet.remove(extensionInfo);
        }

        operatorMap.remove(id);
        operatorIdExtensionInfoMap.remove(id);

        log.info("Unregistered operator: {} => {}", id, extensionInfo);
    }

    @Override
    public void registerPendingItems() {
        registerExtensions();
    }

    @Override
    public void runAsync() {
        Thread thread = new Thread(() -> {
            try {
                run();
            } catch (IOException e) {
                throw new RuntimeException(e);
            }
        });

        thread.setName("ExtensionService: " + UUID.randomUUID());

        thread.start();
    }
}
