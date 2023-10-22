package io.apibrew.client.model;

import java.util.Objects;
import io.apibrew.client.EntityInfo;
import io.apibrew.client.Entity;
import io.apibrew.client.Client;
import com.fasterxml.jackson.annotation.JsonValue;
import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class DataSource extends Entity {
    
    private java.util.UUID id;
    
    private int version;
    
    private DataSource.AuditData auditData;
    
    private String name;
    
    private String description;
    
    private DataSource.Backend backend;
    
    private java.util.Map<String, String> options;

    public static final String NAMESPACE = "system";
    public static final String RESOURCE = "DataSource";

    @JsonIgnore
    public static final EntityInfo<DataSource> entityInfo = new EntityInfo<>("system", "DataSource", DataSource.class, "system-datasource");

    public static class AuditData {
        
        private String createdBy;
        
        private String updatedBy;
        @JsonFormat(shape = JsonFormat.Shape.STRING, timezone = "UTC")
        private java.time.Instant createdOn;
        @JsonFormat(shape = JsonFormat.Shape.STRING, timezone = "UTC")
        private java.time.Instant updatedOn;

        public String getCreatedBy() {
            return createdBy;
        }

        public void setCreatedBy(String createdBy) {
            this.createdBy = createdBy;
        }

        public AuditData withCreatedBy(String createdBy) {
            this.createdBy = createdBy;

            return this;
        }
        public String getUpdatedBy() {
            return updatedBy;
        }

        public void setUpdatedBy(String updatedBy) {
            this.updatedBy = updatedBy;
        }

        public AuditData withUpdatedBy(String updatedBy) {
            this.updatedBy = updatedBy;

            return this;
        }
        public java.time.Instant getCreatedOn() {
            return createdOn;
        }

        public void setCreatedOn(java.time.Instant createdOn) {
            this.createdOn = createdOn;
        }

        public AuditData withCreatedOn(java.time.Instant createdOn) {
            this.createdOn = createdOn;

            return this;
        }
        public java.time.Instant getUpdatedOn() {
            return updatedOn;
        }

        public void setUpdatedOn(java.time.Instant updatedOn) {
            this.updatedOn = updatedOn;
        }

        public AuditData withUpdatedOn(java.time.Instant updatedOn) {
            this.updatedOn = updatedOn;

            return this;
        }
    }

    public static enum Backend {
        POSTGRESQL("POSTGRESQL"),
        MYSQL("MYSQL"),
        MONGODB("MONGODB"),
        REDIS("REDIS");

        private final String value;

        Backend(String value) {
            this.value = value;
        }

        @JsonValue
        public String getValue() {
            return value;
        }
    }

    

    public DataSource() {
    }

    public java.util.UUID getId() {
        return id;
    }

    public void setId(java.util.UUID id) {
        this.id = id;
    }

    public DataSource withId(java.util.UUID id) {
        this.id = id;

        return this;
    }
    public int getVersion() {
        return version;
    }

    public void setVersion(int version) {
        this.version = version;
    }

    public DataSource withVersion(int version) {
        this.version = version;

        return this;
    }
    public DataSource.AuditData getAuditData() {
        return auditData;
    }

    public void setAuditData(DataSource.AuditData auditData) {
        this.auditData = auditData;
    }

    public DataSource withAuditData(DataSource.AuditData auditData) {
        this.auditData = auditData;

        return this;
    }
    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public DataSource withName(String name) {
        this.name = name;

        return this;
    }
    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public DataSource withDescription(String description) {
        this.description = description;

        return this;
    }
    public DataSource.Backend getBackend() {
        return backend;
    }

    public void setBackend(DataSource.Backend backend) {
        this.backend = backend;
    }

    public DataSource withBackend(DataSource.Backend backend) {
        this.backend = backend;

        return this;
    }
    public java.util.Map<String, String> getOptions() {
        return options;
    }

    public void setOptions(java.util.Map<String, String> options) {
        this.options = options;
    }

    public DataSource withOptions(java.util.Map<String, String> options) {
        this.options = options;

        return this;
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
        if (!Objects.equals(this.auditData, obj.auditData)) {
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


