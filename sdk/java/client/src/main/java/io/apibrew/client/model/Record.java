package io.apibrew.client.model;

import java.util.Objects;
import io.apibrew.client.EntityInfo;
import io.apibrew.client.Entity;
import com.fasterxml.jackson.annotation.JsonValue;
import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class Record extends Entity {
    
    private java.util.UUID id;
    
    private Object properties;
    
    private java.util.List<Object> packedProperties;

    public static final String NAMESPACE = "system";
    public static final String RESOURCE = "Record";

    @JsonIgnore
    public static final EntityInfo<Record> entityInfo = new EntityInfo<>("system", "Record", Record.class, "system-record");



    public Record() {
    }

    public java.util.UUID getId() {
        return id;
    }

    public void setId(java.util.UUID id) {
        this.id = id;
    }

    public Record withId(java.util.UUID id) {
        this.id = id;

        return this;
    }
    public Object getProperties() {
        return properties;
    }

    public void setProperties(Object properties) {
        this.properties = properties;
    }

    public Record withProperties(Object properties) {
        this.properties = properties;

        return this;
    }
    public java.util.List<Object> getPackedProperties() {
        return packedProperties;
    }

    public void setPackedProperties(java.util.List<Object> packedProperties) {
        this.packedProperties = packedProperties;
    }

    public Record withPackedProperties(java.util.List<Object> packedProperties) {
        this.packedProperties = packedProperties;

        return this;
    }

    @Override
    public boolean equals(Object o) {
        if (!(o instanceof Record)) {
            return false;
        }

        Record obj = (Record) o;

        if (!Objects.equals(this.id, obj.id)) {
            return false;
        }
        if (!Objects.equals(this.properties, obj.properties)) {
            return false;
        }
        if (!Objects.equals(this.packedProperties, obj.packedProperties)) {
            return false;
        }

        return true;
    }

    @Override
    public int hashCode() {
        if (id == null) {
            return super.hashCode();
        }

        return id.hashCode();
    }
}


