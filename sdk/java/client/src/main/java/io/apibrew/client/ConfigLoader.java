package io.apibrew.client;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.dataformat.yaml.YAMLFactory;
import lombok.SneakyThrows;

import java.io.File;

public class ConfigLoader {

    private final static ObjectMapper objectMapper = new ObjectMapper(new YAMLFactory());

    @SneakyThrows
    public static Config load() {
        File file = new File(System.getenv("HOME") + "/.apbr/config");

        if (file.exists()) {
            return objectMapper.readValue(file, Config.class);
        }

        return null;
    }

}
