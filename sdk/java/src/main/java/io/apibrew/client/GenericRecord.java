package io.apibrew.client;

import com.fasterxml.jackson.annotation.JsonAnyGetter;
import com.fasterxml.jackson.annotation.JsonAnySetter;
import lombok.Data;

import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

@Data
public class GenericRecord extends Entity {

    private final Map<String, Object> properties = new HashMap<>();

    @JsonAnyGetter
    public Map<String, Object> getProperties() {
        return properties;
    }

    @JsonAnySetter
    public void setProperties(String name, Object value) {
        properties.put(name, value);
    }

    @Override
    public UUID getId() {
        if (properties.get("id") == null) {
            return null;
        }

        return UUID.fromString((String) properties.get("id"));
    }

    @Override
    public void setId(UUID id) {
        properties.put("id", id.toString());
    }

    @Override
    public Map<String, Object> getPropertyMap() {
        return getProperties();
    }

    @Override
    public Object getProperty(String property) {
        return getProperties().get(property);
    }

    @Override
    public void setProperty(String property, Object value) {
        getProperties().put(property, value);
    }



}
