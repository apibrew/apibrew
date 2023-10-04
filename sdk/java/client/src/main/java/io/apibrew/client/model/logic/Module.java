package io.apibrew.client.model.logic;

import java.util.Objects;
import io.apibrew.client.EntityInfo;
import io.apibrew.client.Entity;
import com.fasterxml.jackson.annotation.JsonValue;
import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class Module extends Entity {
    
    private java.util.UUID id;
    
    private String $package;
    
    private byte[] content;
    
    private FunctionExecutionEngine engine;
    
    private java.util.Map<String, String> annotations;
    
    private String createdBy;
    
    private String updatedBy;
    @JsonFormat(shape = JsonFormat.Shape.STRING, timezone = "UTC")
    private java.time.Instant createdOn;
    @JsonFormat(shape = JsonFormat.Shape.STRING, timezone = "UTC")
    private java.time.Instant updatedOn;
    
    private int version;

    @JsonIgnore
    public static final EntityInfo<Module> entityInfo = new EntityInfo<>("logic", "Module", Module.class, "logic-module");



    public Module() {
    }

    public java.util.UUID getId() {
        return id;
    }

    public void setId(java.util.UUID id) {
        this.id = id;
    }

    public Module withId(java.util.UUID id) {
        this.id = id;

        return this;
    }
    public String getPackage() {
        return $package;
    }

    public void setPackage(String $package) {
        this.$package = $package;
    }

    public Module withPackage(String $package) {
        this.$package = $package;

        return this;
    }
    public byte[] getContent() {
        return content;
    }

    public void setContent(byte[] content) {
        this.content = content;
    }

    public Module withContent(byte[] content) {
        this.content = content;

        return this;
    }
    public FunctionExecutionEngine getEngine() {
        return engine;
    }

    public void setEngine(FunctionExecutionEngine engine) {
        this.engine = engine;
    }

    public Module withEngine(FunctionExecutionEngine engine) {
        this.engine = engine;

        return this;
    }
    public java.util.Map<String, String> getAnnotations() {
        return annotations;
    }

    public void setAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;
    }

    public Module withAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;

        return this;
    }
    public String getCreatedBy() {
        return createdBy;
    }

    public void setCreatedBy(String createdBy) {
        this.createdBy = createdBy;
    }

    public Module withCreatedBy(String createdBy) {
        this.createdBy = createdBy;

        return this;
    }
    public String getUpdatedBy() {
        return updatedBy;
    }

    public void setUpdatedBy(String updatedBy) {
        this.updatedBy = updatedBy;
    }

    public Module withUpdatedBy(String updatedBy) {
        this.updatedBy = updatedBy;

        return this;
    }
    public java.time.Instant getCreatedOn() {
        return createdOn;
    }

    public void setCreatedOn(java.time.Instant createdOn) {
        this.createdOn = createdOn;
    }

    public Module withCreatedOn(java.time.Instant createdOn) {
        this.createdOn = createdOn;

        return this;
    }
    public java.time.Instant getUpdatedOn() {
        return updatedOn;
    }

    public void setUpdatedOn(java.time.Instant updatedOn) {
        this.updatedOn = updatedOn;
    }

    public Module withUpdatedOn(java.time.Instant updatedOn) {
        this.updatedOn = updatedOn;

        return this;
    }
    public int getVersion() {
        return version;
    }

    public void setVersion(int version) {
        this.version = version;
    }

    public Module withVersion(int version) {
        this.version = version;

        return this;
    }

    @Override
    public boolean equals(Object o) {
        if (!(o instanceof Module)) {
            return false;
        }

        Module obj = (Module) o;

        if (!Objects.equals(this.id, obj.id)) {
            return false;
        }
        if (!Objects.equals(this.$package, obj.$package)) {
            return false;
        }
        if (!Objects.equals(this.content, obj.content)) {
            return false;
        }
        if (!Objects.equals(this.engine, obj.engine)) {
            return false;
        }
        if (!Objects.equals(this.annotations, obj.annotations)) {
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
        if (!Objects.equals(this.version, obj.version)) {
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


