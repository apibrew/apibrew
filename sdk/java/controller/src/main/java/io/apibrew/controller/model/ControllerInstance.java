package io.apibrew.controller.model;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonInclude;
import io.apibrew.client.Entity;
import io.apibrew.client.EntityInfo;
import io.apibrew.client.model.AuditLog;
import lombok.Getter;

import java.util.Objects;

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
