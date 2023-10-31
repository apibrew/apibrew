package io.apibrew.client.controller.model;

import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public interface ControllerInstance {

    public String getName();

    ServerConfig getServerConfig();

    public interface ServerConfig {

        String getHost();

        boolean getInsecure();

        Authentication getAuthentication();

        int getPort();

        int getHttpPort();
    }

    public interface Authentication {

        String getUsername();

        String getPassword();

        String getToken();
    }

}
