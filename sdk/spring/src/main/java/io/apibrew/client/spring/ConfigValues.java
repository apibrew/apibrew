package io.apibrew.client.spring;

import io.apibrew.client.Config;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Configuration;

@Configuration
public class ConfigValues {

    @Value("${apibrew.channelKey:default}")
    public String channelKey = "default";

    @Value("${apibrew.serviceName:default}")
    String serviceName = "default";

    Config.Server server;

    @Value("${apibrew.profile:default}")
    String profile;
}
