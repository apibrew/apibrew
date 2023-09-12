package io.apibrew.lib.model;

import java.util.Objects;
import io.apibrew.lib.EntityInfo;
import io.apibrew.lib.Entity;

public class DataSource extends Entity {
    private java.util.UUID id;
    private int version;
    private String createdBy;
    private String updatedBy;
    private java.time.LocalDateTime createdOn;
    private java.time.LocalDateTime updatedOn;
    private String name;
    private String description;
    private DataSource.Backend backend;
    private java.util.Map<String, String> options;

    public static final EntityInfo<DataSource> entityInfo = new EntityInfo<>("system", "DataSource", DataSource.class);


    public static enum Backend {
        POSTGRESQL("POSTGRESQL"),
        MYSQL("MYSQL"),
        MONGODB("MONGODB"),
        REDIS("REDIS");

        private final String value;

        Backend(String value) {
            this.value = value;
        }

        public String getValue() {
            return value;
        }
    }

    public DataSource() {
    }

    public EntityInfo<DataSource> getEntityInfo() {
        return entityInfo;
    }

    public java.util.UUID getId() {
        return id;
    }

    public void setId(java.util.UUID id) {
        this.id = id;
    }
    public int getVersion() {
        return version;
    }

    public void setVersion(int version) {
        this.version = version;
    }
    public String getCreatedBy() {
        return createdBy;
    }

    public void setCreatedBy(String createdBy) {
        this.createdBy = createdBy;
    }
    public String getUpdatedBy() {
        return updatedBy;
    }

    public void setUpdatedBy(String updatedBy) {
        this.updatedBy = updatedBy;
    }
    public java.time.LocalDateTime getCreatedOn() {
        return createdOn;
    }

    public void setCreatedOn(java.time.LocalDateTime createdOn) {
        this.createdOn = createdOn;
    }
    public java.time.LocalDateTime getUpdatedOn() {
        return updatedOn;
    }

    public void setUpdatedOn(java.time.LocalDateTime updatedOn) {
        this.updatedOn = updatedOn;
    }
    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }
    public DataSource.Backend getBackend() {
        return backend;
    }

    public void setBackend(DataSource.Backend backend) {
        this.backend = backend;
    }
    public java.util.Map<String, String> getOptions() {
        return options;
    }

    public void setOptions(java.util.Map<String, String> options) {
        this.options = options;
    }

    @Override
    public boolean equals(Object o) {
        if (!(o instanceof DataSource)) {
            return false;
        }

        DataSource obj = (DataSource) o;

        if (!Objects.equals(this.id, obj.id)) {
            return false;
        }
        if (!Objects.equals(this.version, obj.version)) {
            return false;
        }
        if (!Objects.equals(this.createdBy, obj.createdBy)) {
            return false;
        }
        if (!Objects.equals(this.updatedBy, obj.updatedBy)) {
            return false;
        }
        if (!Objects.equals(this.createdOn, obj.createdOn)) {
            return false;
        }
        if (!Objects.equals(this.updatedOn, obj.updatedOn)) {
            return false;
        }
        if (!Objects.equals(this.name, obj.name)) {
            return false;
        }
        if (!Objects.equals(this.description, obj.description)) {
            return false;
        }
        if (!Objects.equals(this.backend, obj.backend)) {
            return false;
        }
        if (!Objects.equals(this.options, obj.options)) {
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


