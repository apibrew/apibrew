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
public class Role extends Entity {
    
    private java.util.UUID id;
    
    private int version;
    
    private Role.AuditData auditData;
    
    private String name;
    
    private java.util.List<Permission> permissions;
    
    private Object details;

    public static final String NAMESPACE = "system";
    public static final String RESOURCE = "Role";

    @JsonIgnore
    public static final EntityInfo<Role> entityInfo = new EntityInfo<>("system", "Role", Role.class, "system-role");

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


    

    public Role() {
    }

    public java.util.UUID getId() {
        return id;
    }

    public void setId(java.util.UUID id) {
        this.id = id;
    }

    public Role withId(java.util.UUID id) {
        this.id = id;

        return this;
    }
    public int getVersion() {
        return version;
    }

    public void setVersion(int version) {
        this.version = version;
    }

    public Role withVersion(int version) {
        this.version = version;

        return this;
    }
    public Role.AuditData getAuditData() {
        return auditData;
    }

    public void setAuditData(Role.AuditData auditData) {
        this.auditData = auditData;
    }

    public Role withAuditData(Role.AuditData auditData) {
        this.auditData = auditData;

        return this;
    }
    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Role withName(String name) {
        this.name = name;

        return this;
    }
    public java.util.List<Permission> getPermissions() {
        return permissions;
    }

    public void setPermissions(java.util.List<Permission> permissions) {
        this.permissions = permissions;
    }

    public Role withPermissions(java.util.List<Permission> permissions) {
        this.permissions = permissions;

        return this;
    }
    public Object getDetails() {
        return details;
    }

    public void setDetails(Object details) {
        this.details = details;
    }

    public Role withDetails(Object details) {
        this.details = details;

        return this;
    }

    @Override
    public boolean equals(Object o) {
        if (!(o instanceof Role)) {
            return false;
        }

        Role obj = (Role) o;

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
        if (!Objects.equals(this.permissions, obj.permissions)) {
            return false;
        }
        if (!Objects.equals(this.details, obj.details)) {
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


