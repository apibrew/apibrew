package io.apibrew.client;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;

import java.util.List;

@Data
public class Config {
    @JsonProperty
    final String type = "server";
    private String defaultServer;

    private List<Server> servers;


    @Data
    public static class Server {
        private String name;
        private String host;
        private boolean insecure;
        private Authentication authentication;
    }

    @Data
    public static class Authentication {
        private String username;
        private String password;
        private String token;
    }
}
