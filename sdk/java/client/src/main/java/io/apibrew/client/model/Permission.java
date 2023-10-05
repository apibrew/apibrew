package io.apibrew.client.model;

import java.util.Objects;
import io.apibrew.client.EntityInfo;
import io.apibrew.client.Entity;
import com.fasterxml.jackson.annotation.JsonValue;
import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class Permission extends Entity {
    
    private java.util.UUID id;
    
    private int version;
    
    private String createdBy;
    
    private String updatedBy;
    @JsonFormat(shape = JsonFormat.Shape.STRING, timezone = "UTC")
    private java.time.Instant createdOn;
    @JsonFormat(shape = JsonFormat.Shape.STRING, timezone = "UTC")
    private java.time.Instant updatedOn;
    
    private String namespace;
    
    private String resource;
    
    private String property;
    
    private String propertyValue;
    
    private Permission.PropertyMode propertyMode;
    
    private Permission.Operation operation;
    
    private java.util.List<String> recordIds;
    @JsonFormat(shape = JsonFormat.Shape.STRING, timezone = "UTC")
    private java.time.Instant before;
    @JsonFormat(shape = JsonFormat.Shape.STRING, timezone = "UTC")
    private java.time.Instant after;
    
    private User user;
    
    private Role role;
    
    private Permission.Permit permit;
    
    private Object localFlags;

    public static final String NAMESPACE = "system";
    public static final String RESOURCE = "Permission";

    @JsonIgnore
    public static final EntityInfo<Permission> entityInfo = new EntityInfo<>("system", "Permission", Permission.class, "system-permission");


    public static enum PropertyMode {
        PROPERTY_MATCH_ONLY("PROPERTY_MATCH_ONLY"),
        PROPERTY_MATCH_ANY("PROPERTY_MATCH_ANY");

        private final String value;

        PropertyMode(String value) {
            this.value = value;
        }

        @JsonValue
        public String getValue() {
            return value;
        }
    }
    public static enum Operation {
        READ("READ"),
        CREATE("CREATE"),
        UPDATE("UPDATE"),
        DELETE("DELETE"),
        FULL("FULL");

        private final String value;

        Operation(String value) {
            this.value = value;
        }

        @JsonValue
        public String getValue() {
            return value;
        }
    }
    public static enum Permit {
        ALLOW("ALLOW"),
        REJECT("REJECT");

        private final String value;

        Permit(String value) {
            this.value = value;
        }

        @JsonValue
        public String getValue() {
            return value;
        }
    }

    public Permission() {
    }

    public java.util.UUID getId() {
        return id;
    }

    public void setId(java.util.UUID id) {
        this.id = id;
    }

    public Permission withId(java.util.UUID id) {
        this.id = id;

        return this;
    }
    public int getVersion() {
        return version;
    }

    public void setVersion(int version) {
        this.version = version;
    }

    public Permission withVersion(int version) {
        this.version = version;

        return this;
    }
    public String getCreatedBy() {
        return createdBy;
    }

    public void setCreatedBy(String createdBy) {
        this.createdBy = createdBy;
    }

    public Permission withCreatedBy(String createdBy) {
        this.createdBy = createdBy;

        return this;
    }
    public String getUpdatedBy() {
        return updatedBy;
    }

    public void setUpdatedBy(String updatedBy) {
        this.updatedBy = updatedBy;
    }

    public Permission withUpdatedBy(String updatedBy) {
        this.updatedBy = updatedBy;

        return this;
    }
    public java.time.Instant getCreatedOn() {
        return createdOn;
    }

    public void setCreatedOn(java.time.Instant createdOn) {
        this.createdOn = createdOn;
    }

    public Permission withCreatedOn(java.time.Instant createdOn) {
        this.createdOn = createdOn;

        return this;
    }
    public java.time.Instant getUpdatedOn() {
        return updatedOn;
    }

    public void setUpdatedOn(java.time.Instant updatedOn) {
        this.updatedOn = updatedOn;
    }

    public Permission withUpdatedOn(java.time.Instant updatedOn) {
        this.updatedOn = updatedOn;

        return this;
    }
    public String getNamespace() {
        return namespace;
    }

    public void setNamespace(String namespace) {
        this.namespace = namespace;
    }

    public Permission withNamespace(String namespace) {
        this.namespace = namespace;

        return this;
    }
    public String getResource() {
        return resource;
    }

    public void setResource(String resource) {
        this.resource = resource;
    }

    public Permission withResource(String resource) {
        this.resource = resource;

        return this;
    }
    public String getProperty() {
        return property;
    }

    public void setProperty(String property) {
        this.property = property;
    }

    public Permission withProperty(String property) {
        this.property = property;

        return this;
    }
    public String getPropertyValue() {
        return propertyValue;
    }

    public void setPropertyValue(String propertyValue) {
        this.propertyValue = propertyValue;
    }

    public Permission withPropertyValue(String propertyValue) {
        this.propertyValue = propertyValue;

        return this;
    }
    public Permission.PropertyMode getPropertyMode() {
        return propertyMode;
    }

    public void setPropertyMode(Permission.PropertyMode propertyMode) {
        this.propertyMode = propertyMode;
    }

    public Permission withPropertyMode(Permission.PropertyMode propertyMode) {
        this.propertyMode = propertyMode;

        return this;
    }
    public Permission.Operation getOperation() {
        return operation;
    }

    public void setOperation(Permission.Operation operation) {
        this.operation = operation;
    }

    public Permission withOperation(Permission.Operation operation) {
        this.operation = operation;

        return this;
    }
    public java.util.List<String> getRecordIds() {
        return recordIds;
    }

    public void setRecordIds(java.util.List<String> recordIds) {
        this.recordIds = recordIds;
    }

    public Permission withRecordIds(java.util.List<String> recordIds) {
        this.recordIds = recordIds;

        return this;
    }
    public java.time.Instant getBefore() {
        return before;
    }

    public void setBefore(java.time.Instant before) {
        this.before = before;
    }

    public Permission withBefore(java.time.Instant before) {
        this.before = before;

        return this;
    }
    public java.time.Instant getAfter() {
        return after;
    }

    public void setAfter(java.time.Instant after) {
        this.after = after;
    }

    public Permission withAfter(java.time.Instant after) {
        this.after = after;

        return this;
    }
    public User getUser() {
        return user;
    }

    public void setUser(User user) {
        this.user = user;
    }

    public Permission withUser(User user) {
        this.user = user;

        return this;
    }
    public Role getRole() {
        return role;
    }

    public void setRole(Role role) {
        this.role = role;
    }

    public Permission withRole(Role role) {
        this.role = role;

        return this;
    }
    public Permission.Permit getPermit() {
        return permit;
    }

    public void setPermit(Permission.Permit permit) {
        this.permit = permit;
    }

    public Permission withPermit(Permission.Permit permit) {
        this.permit = permit;

        return this;
    }
    public Object getLocalFlags() {
        return localFlags;
    }

    public void setLocalFlags(Object localFlags) {
        this.localFlags = localFlags;
    }

    public Permission withLocalFlags(Object localFlags) {
        this.localFlags = localFlags;

        return this;
    }

    @Override
    public boolean equals(Object o) {
        if (!(o instanceof Permission)) {
            return false;
        }

        Permission obj = (Permission) o;

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
        if (!Objects.equals(this.namespace, obj.namespace)) {
            return false;
        }
        if (!Objects.equals(this.resource, obj.resource)) {
            return false;
        }
        if (!Objects.equals(this.property, obj.property)) {
            return false;
        }
        if (!Objects.equals(this.propertyValue, obj.propertyValue)) {
            return false;
        }
        if (!Objects.equals(this.propertyMode, obj.propertyMode)) {
            return false;
        }
        if (!Objects.equals(this.operation, obj.operation)) {
            return false;
        }
        if (!Objects.equals(this.recordIds, obj.recordIds)) {
            return false;
        }
        if (!Objects.equals(this.before, obj.before)) {
            return false;
        }
        if (!Objects.equals(this.after, obj.after)) {
            return false;
        }
        if (!Objects.equals(this.user, obj.user)) {
            return false;
        }
        if (!Objects.equals(this.role, obj.role)) {
            return false;
        }
        if (!Objects.equals(this.permit, obj.permit)) {
            return false;
        }
        if (!Objects.equals(this.localFlags, obj.localFlags)) {
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


