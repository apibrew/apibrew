package io.apibrew.client.spring;

import io.apibrew.client.Config;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;

@Configuration
@ConfigurationProperties(prefix = "mail")
public class ConfigValues {
    public String channelKey = "default";
    String serviceName = "default";
    Config.Server server;
    String profile;
}
