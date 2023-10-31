package io.apibrew.client.storage.model;

import java.util.Objects;
import io.apibrew.client.EntityInfo;
import io.apibrew.client.Entity;
import io.apibrew.client.Client;
import com.fasterxml.jackson.annotation.JsonValue;
import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class StorageObject extends Entity {
    
    private java.util.UUID id;
    
    private String name;
    
    private java.util.Map<String, String> annotations;
    
    private String contentType;
    
    private Long size;
    
    private boolean allowDownloadPublicly;
    
    private boolean allowUploadPublicly;
    
    private int version;
    
    private StorageObject.AuditData auditData;

    public static final String NAMESPACE = "storage";
    public static final String RESOURCE = "StorageObject";

    @JsonIgnore
    public static final EntityInfo<StorageObject> entityInfo = new EntityInfo<>("storage", "StorageObject", StorageObject.class, "storage-storageobject");

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

        @Override
        public boolean equals(Object o) {
            if (!(o instanceof AuditData)) {
                return false;
            }

            AuditData obj = (AuditData) o;

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

            return true;
        }

        @Override
        public int hashCode() {
           return Objects.hash(createdBy, updatedBy, createdOn, updatedOn);
        }
    }


    

    public StorageObject() {
    }

    public java.util.UUID getId() {
        return id;
    }

    public void setId(java.util.UUID id) {
        this.id = id;
    }

    public StorageObject withId(java.util.UUID id) {
        this.id = id;

        return this;
    }
    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public StorageObject withName(String name) {
        this.name = name;

        return this;
    }
    public java.util.Map<String, String> getAnnotations() {
        return annotations;
    }

    public void setAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;
    }

    public StorageObject withAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;

        return this;
    }
    public String getContentType() {
        return contentType;
    }

    public void setContentType(String contentType) {
        this.contentType = contentType;
    }

    public StorageObject withContentType(String contentType) {
        this.contentType = contentType;

        return this;
    }
    public Long getSize() {
        return size;
    }

    public void setSize(Long size) {
        this.size = size;
    }

    public StorageObject withSize(Long size) {
        this.size = size;

        return this;
    }
    public boolean getAllowDownloadPublicly() {
        return allowDownloadPublicly;
    }

    public void setAllowDownloadPublicly(boolean allowDownloadPublicly) {
        this.allowDownloadPublicly = allowDownloadPublicly;
    }

    public StorageObject withAllowDownloadPublicly(boolean allowDownloadPublicly) {
        this.allowDownloadPublicly = allowDownloadPublicly;

        return this;
    }
    public boolean getAllowUploadPublicly() {
        return allowUploadPublicly;
    }

    public void setAllowUploadPublicly(boolean allowUploadPublicly) {
        this.allowUploadPublicly = allowUploadPublicly;
    }

    public StorageObject withAllowUploadPublicly(boolean allowUploadPublicly) {
        this.allowUploadPublicly = allowUploadPublicly;

        return this;
    }
    public int getVersion() {
        return version;
    }

    public void setVersion(int version) {
        this.version = version;
    }

    public StorageObject withVersion(int version) {
        this.version = version;

        return this;
    }
    public StorageObject.AuditData getAuditData() {
        return auditData;
    }

    public void setAuditData(StorageObject.AuditData auditData) {
        this.auditData = auditData;
    }

    public StorageObject withAuditData(StorageObject.AuditData auditData) {
        this.auditData = auditData;

        return this;
    }

    @Override
    public boolean equals(Object o) {
        if (!(o instanceof StorageObject)) {
            return false;
        }

        StorageObject obj = (StorageObject) o;

        if (!Objects.equals(this.id, obj.id)) {
            return false;
        }
        if (!Objects.equals(this.name, obj.name)) {
            return false;
        }
        if (!Objects.equals(this.annotations, obj.annotations)) {
            return false;
        }
        if (!Objects.equals(this.contentType, obj.contentType)) {
            return false;
        }
        if (!Objects.equals(this.size, obj.size)) {
            return false;
        }
        if (!Objects.equals(this.allowDownloadPublicly, obj.allowDownloadPublicly)) {
            return false;
        }
        if (!Objects.equals(this.allowUploadPublicly, obj.allowUploadPublicly)) {
            return false;
        }
        if (!Objects.equals(this.version, obj.version)) {
            return false;
        }
        if (!Objects.equals(this.auditData, obj.auditData)) {
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


