package io.apibrew.lib.model;

import java.util.Objects;
import io.apibrew.lib.EntityInfo;
import io.apibrew.lib.Entity;
import com.fasterxml.jackson.annotation.JsonValue;
import com.fasterxml.jackson.annotation.JsonFormat;

public class Record extends Entity {
    private java.util.UUID id;
    private Object properties;
    private java.util.List<Object> packedProperties;

    public static final EntityInfo<Record> entityInfo = new EntityInfo<>("system", "Record", Record.class, "system-record");



    public Record() {
    }

    public EntityInfo<Record> getEntityInfo() {
        return entityInfo;
    }

    public java.util.UUID getId() {
        return id;
    }

    public void setId(java.util.UUID id) {
        this.id = id;
    }
    public Object getProperties() {
        return properties;
    }

    public void setProperties(Object properties) {
        this.properties = properties;
    }
    public java.util.List<Object> getPackedProperties() {
        return packedProperties;
    }

    public void setPackedProperties(java.util.List<Object> packedProperties) {
        this.packedProperties = packedProperties;
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


