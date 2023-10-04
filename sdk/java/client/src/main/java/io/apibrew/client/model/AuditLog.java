package io.apibrew.client.model;

import java.util.Objects;
import io.apibrew.client.EntityInfo;
import io.apibrew.client.Entity;
import com.fasterxml.jackson.annotation.JsonValue;
import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class AuditLog extends Entity {
    
    private java.util.UUID id;
    
    private int version;
    
    private String namespace;
    
    private String resource;
    
    private String recordId;
    @JsonFormat(shape = JsonFormat.Shape.STRING, timezone = "UTC")
    private java.time.Instant time;
    
    private String username;
    
    private AuditLog.Operation operation;
    
    private Object properties;
    
    private java.util.Map<String, String> annotations;

    @JsonIgnore
    public static final EntityInfo<AuditLog> entityInfo = new EntityInfo<>("system", "AuditLog", AuditLog.class, "system-auditlog");


    public static enum Operation {
        CREATE("CREATE"),
        UPDATE("UPDATE"),
        DELETE("DELETE");

        private final String value;

        Operation(String value) {
            this.value = value;
        }

        @JsonValue
        public String getValue() {
            return value;
        }
    }

    public AuditLog() {
    }

    public java.util.UUID getId() {
        return id;
    }

    public void setId(java.util.UUID id) {
        this.id = id;
    }

    public AuditLog withId(java.util.UUID id) {
        this.id = id;

        return this;
    }
    public int getVersion() {
        return version;
    }

    public void setVersion(int version) {
        this.version = version;
    }

    public AuditLog withVersion(int version) {
        this.version = version;

        return this;
    }
    public String getNamespace() {
        return namespace;
    }

    public void setNamespace(String namespace) {
        this.namespace = namespace;
    }

    public AuditLog withNamespace(String namespace) {
        this.namespace = namespace;

        return this;
    }
    public String getResource() {
        return resource;
    }

    public void setResource(String resource) {
        this.resource = resource;
    }

    public AuditLog withResource(String resource) {
        this.resource = resource;

        return this;
    }
    public String getRecordId() {
        return recordId;
    }

    public void setRecordId(String recordId) {
        this.recordId = recordId;
    }

    public AuditLog withRecordId(String recordId) {
        this.recordId = recordId;

        return this;
    }
    public java.time.Instant getTime() {
        return time;
    }

    public void setTime(java.time.Instant time) {
        this.time = time;
    }

    public AuditLog withTime(java.time.Instant time) {
        this.time = time;

        return this;
    }
    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public AuditLog withUsername(String username) {
        this.username = username;

        return this;
    }
    public AuditLog.Operation getOperation() {
        return operation;
    }

    public void setOperation(AuditLog.Operation operation) {
        this.operation = operation;
    }

    public AuditLog withOperation(AuditLog.Operation operation) {
        this.operation = operation;

        return this;
    }
    public Object getProperties() {
        return properties;
    }

    public void setProperties(Object properties) {
        this.properties = properties;
    }

    public AuditLog withProperties(Object properties) {
        this.properties = properties;

        return this;
    }
    public java.util.Map<String, String> getAnnotations() {
        return annotations;
    }

    public void setAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;
    }

    public AuditLog withAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;

        return this;
    }

    @Override
    public boolean equals(Object o) {
        if (!(o instanceof AuditLog)) {
            return false;
        }

        AuditLog obj = (AuditLog) o;

        if (!Objects.equals(this.id, obj.id)) {
            return false;
        }
        if (!Objects.equals(this.version, obj.version)) {
            return false;
        }
        if (!Objects.equals(this.namespace, obj.namespace)) {
            return false;
        }
        if (!Objects.equals(this.resource, obj.resource)) {
            return false;
        }
        if (!Objects.equals(this.recordId, obj.recordId)) {
            return false;
        }
        if (!Objects.equals(this.time, obj.time)) {
            return false;
        }
        if (!Objects.equals(this.username, obj.username)) {
            return false;
        }
        if (!Objects.equals(this.operation, obj.operation)) {
            return false;
        }
        if (!Objects.equals(this.properties, obj.properties)) {
            return false;
        }
        if (!Objects.equals(this.annotations, obj.annotations)) {
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


