package io.apibrew.client.nano.model;

import java.util.Objects;
import io.apibrew.client.EntityInfo;
import io.apibrew.client.Entity;
import io.apibrew.client.Client;
import com.fasterxml.jackson.annotation.JsonValue;
import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class Code extends Entity {
    
    private java.util.UUID id;
    
    private String name;
    
    private Code.Language language;
    
    private String content;
    
    private Code.ContentFormat contentFormat;
    
    private java.util.Map<String, String> annotations;
    
    private int version;
    
    private Code.AuditData auditData;

    public static final String NAMESPACE = "nano";
    public static final String RESOURCE = "Code";

    @JsonIgnore
    public static final EntityInfo<Code> entityInfo = new EntityInfo<>("nano", "Code", Code.class, "nano-code");

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

    public static enum Language {
        PYTHON("PYTHON"),
        JAVASCRIPT("JAVASCRIPT");

        private final String value;

        Language(String value) {
            this.value = value;
        }

        @JsonValue
        public String getValue() {
            return value;
        }
    }
    public static enum ContentFormat {
        TEXT("TEXT"),
        TAR("TAR"),
        TAR_GZ("TAR_GZ");

        private final String value;

        ContentFormat(String value) {
            this.value = value;
        }

        @JsonValue
        public String getValue() {
            return value;
        }
    }

    

    public Code() {
    }

    public java.util.UUID getId() {
        return id;
    }

    public void setId(java.util.UUID id) {
        this.id = id;
    }

    public Code withId(java.util.UUID id) {
        this.id = id;

        return this;
    }
    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Code withName(String name) {
        this.name = name;

        return this;
    }
    public Code.Language getLanguage() {
        return language;
    }

    public void setLanguage(Code.Language language) {
        this.language = language;
    }

    public Code withLanguage(Code.Language language) {
        this.language = language;

        return this;
    }
    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
    }

    public Code withContent(String content) {
        this.content = content;

        return this;
    }
    public Code.ContentFormat getContentFormat() {
        return contentFormat;
    }

    public void setContentFormat(Code.ContentFormat contentFormat) {
        this.contentFormat = contentFormat;
    }

    public Code withContentFormat(Code.ContentFormat contentFormat) {
        this.contentFormat = contentFormat;

        return this;
    }
    public java.util.Map<String, String> getAnnotations() {
        return annotations;
    }

    public void setAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;
    }

    public Code withAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;

        return this;
    }
    public int getVersion() {
        return version;
    }

    public void setVersion(int version) {
        this.version = version;
    }

    public Code withVersion(int version) {
        this.version = version;

        return this;
    }
    public Code.AuditData getAuditData() {
        return auditData;
    }

    public void setAuditData(Code.AuditData auditData) {
        this.auditData = auditData;
    }

    public Code withAuditData(Code.AuditData auditData) {
        this.auditData = auditData;

        return this;
    }

    @Override
    public boolean equals(Object o) {
        if (!(o instanceof Code)) {
            return false;
        }

        Code obj = (Code) o;

        if (!Objects.equals(this.id, obj.id)) {
            return false;
        }
        if (!Objects.equals(this.name, obj.name)) {
            return false;
        }
        if (!Objects.equals(this.language, obj.language)) {
            return false;
        }
        if (!Objects.equals(this.content, obj.content)) {
            return false;
        }
        if (!Objects.equals(this.contentFormat, obj.contentFormat)) {
            return false;
        }
        if (!Objects.equals(this.annotations, obj.annotations)) {
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


