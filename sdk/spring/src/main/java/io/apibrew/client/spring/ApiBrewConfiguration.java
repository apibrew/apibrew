package io.apibrew.client.spring;

import io.apibrew.client.Client;
import io.apibrew.client.Entity;
import io.apibrew.client.Repository;
import io.apibrew.client.ext.ExtensionService;
import io.apibrew.client.ext.Handler;
import io.apibrew.client.ext.impl.PollerExtensionService;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.InjectionPoint;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.config.DependencyDescriptor;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.*;
import org.springframework.core.ResolvableType;

@Configuration
@RequiredArgsConstructor
@EnableConfigurationProperties
@Import({ConfigValues.class})
@ComponentScan(basePackages = "io.apibrew.client.spring")
public class ApiBrewConfiguration {

    @Autowired
    ApplicationContext applicationContext;

    @Bean
    public Client client(ConfigValues configValues) {
        if (configValues.server != null) {
            return Client.newClientByServerConfig(configValues.server);
        }

        if (configValues.profile != null) {
            return Client.newClientByServerName(configValues.profile);
        }

        return Client.newClient();
    }

    @Bean
    ExtensionService extensionService(Client client, ConfigValues configValues) {
        PollerExtensionService ext = new PollerExtensionService(configValues.serviceName, client, configValues.channelKey);

        ext.runAsync();

        return ext;
    }

    @Bean
    @Scope("prototype")
    public <E extends Entity> Repository<E> repository(final Client client, final InjectionPoint ip) {
        ResolvableType resolved = ((DependencyDescriptor) ip).getResolvableType();

        Class<E> parameterClass = (Class<E>) resolved.getGeneric(0).resolve();

        return client.repo(parameterClass);
    }

//    @Bean
//    public <E extends Entity> Handler<E> handler(final ExtensionService extensionService, final InjectionPoint ip) {
//        ResolvableType resolved = ((DependencyDescriptor)ip).getResolvableType();
//
//        Class<E> parameterClass = (Class<E>) resolved.getGeneric(0).resolve();
//
//        return extensionService.handler(parameterClass);
//    }
}
