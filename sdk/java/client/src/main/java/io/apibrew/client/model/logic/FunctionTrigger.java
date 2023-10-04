package io.apibrew.client.model.logic;

import java.util.Objects;
import io.apibrew.client.EntityInfo;
import io.apibrew.client.Entity;
import com.fasterxml.jackson.annotation.JsonValue;
import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class FunctionTrigger extends Entity {
    
    private java.util.UUID id;
    
    private String name;
    
    private String resource;
    
    private String namespace;
    
    private FunctionTrigger.Action action;
    
    private FunctionTrigger.Order order;
    
    private boolean async;
    
    private Function function;
    
    private java.util.Map<String, String> annotations;
    
    private String createdBy;
    
    private String updatedBy;
    @JsonFormat(shape = JsonFormat.Shape.STRING, timezone = "UTC")
    private java.time.Instant createdOn;
    @JsonFormat(shape = JsonFormat.Shape.STRING, timezone = "UTC")
    private java.time.Instant updatedOn;
    
    private int version;

    @JsonIgnore
    public static final EntityInfo<FunctionTrigger> entityInfo = new EntityInfo<>("logic", "FunctionTrigger", FunctionTrigger.class, "logic-functiontrigger");


    public static enum Action {
        CREATE("create"),
        UPDATE("update"),
        DELETE("delete"),
        LIST("list"),
        GET("get");

        private final String value;

        Action(String value) {
            this.value = value;
        }

        @JsonValue
        public String getValue() {
            return value;
        }
    }
    public static enum Order {
        BEFORE("before"),
        AFTER("after"),
        INSTEAD("instead");

        private final String value;

        Order(String value) {
            this.value = value;
        }

        @JsonValue
        public String getValue() {
            return value;
        }
    }

    public FunctionTrigger() {
    }

    public java.util.UUID getId() {
        return id;
    }

    public void setId(java.util.UUID id) {
        this.id = id;
    }

    public FunctionTrigger withId(java.util.UUID id) {
        this.id = id;

        return this;
    }
    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public FunctionTrigger withName(String name) {
        this.name = name;

        return this;
    }
    public String getResource() {
        return resource;
    }

    public void setResource(String resource) {
        this.resource = resource;
    }

    public FunctionTrigger withResource(String resource) {
        this.resource = resource;

        return this;
    }
    public String getNamespace() {
        return namespace;
    }

    public void setNamespace(String namespace) {
        this.namespace = namespace;
    }

    public FunctionTrigger withNamespace(String namespace) {
        this.namespace = namespace;

        return this;
    }
    public FunctionTrigger.Action getAction() {
        return action;
    }

    public void setAction(FunctionTrigger.Action action) {
        this.action = action;
    }

    public FunctionTrigger withAction(FunctionTrigger.Action action) {
        this.action = action;

        return this;
    }
    public FunctionTrigger.Order getOrder() {
        return order;
    }

    public void setOrder(FunctionTrigger.Order order) {
        this.order = order;
    }

    public FunctionTrigger withOrder(FunctionTrigger.Order order) {
        this.order = order;

        return this;
    }
    public boolean getAsync() {
        return async;
    }

    public void setAsync(boolean async) {
        this.async = async;
    }

    public FunctionTrigger withAsync(boolean async) {
        this.async = async;

        return this;
    }
    public Function getFunction() {
        return function;
    }

    public void setFunction(Function function) {
        this.function = function;
    }

    public FunctionTrigger withFunction(Function function) {
        this.function = function;

        return this;
    }
    public java.util.Map<String, String> getAnnotations() {
        return annotations;
    }

    public void setAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;
    }

    public FunctionTrigger withAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;

        return this;
    }
    public String getCreatedBy() {
        return createdBy;
    }

    public void setCreatedBy(String createdBy) {
        this.createdBy = createdBy;
    }

    public FunctionTrigger withCreatedBy(String createdBy) {
        this.createdBy = createdBy;

        return this;
    }
    public String getUpdatedBy() {
        return updatedBy;
    }

    public void setUpdatedBy(String updatedBy) {
        this.updatedBy = updatedBy;
    }

    public FunctionTrigger withUpdatedBy(String updatedBy) {
        this.updatedBy = updatedBy;

        return this;
    }
    public java.time.Instant getCreatedOn() {
        return createdOn;
    }

    public void setCreatedOn(java.time.Instant createdOn) {
        this.createdOn = createdOn;
    }

    public FunctionTrigger withCreatedOn(java.time.Instant createdOn) {
        this.createdOn = createdOn;

        return this;
    }
    public java.time.Instant getUpdatedOn() {
        return updatedOn;
    }

    public void setUpdatedOn(java.time.Instant updatedOn) {
        this.updatedOn = updatedOn;
    }

    public FunctionTrigger withUpdatedOn(java.time.Instant updatedOn) {
        this.updatedOn = updatedOn;

        return this;
    }
    public int getVersion() {
        return version;
    }

    public void setVersion(int version) {
        this.version = version;
    }

    public FunctionTrigger withVersion(int version) {
        this.version = version;

        return this;
    }

    @Override
    public boolean equals(Object o) {
        if (!(o instanceof FunctionTrigger)) {
            return false;
        }

        FunctionTrigger obj = (FunctionTrigger) o;

        if (!Objects.equals(this.id, obj.id)) {
            return false;
        }
        if (!Objects.equals(this.name, obj.name)) {
            return false;
        }
        if (!Objects.equals(this.resource, obj.resource)) {
            return false;
        }
        if (!Objects.equals(this.namespace, obj.namespace)) {
            return false;
        }
        if (!Objects.equals(this.action, obj.action)) {
            return false;
        }
        if (!Objects.equals(this.order, obj.order)) {
            return false;
        }
        if (!Objects.equals(this.async, obj.async)) {
            return false;
        }
        if (!Objects.equals(this.function, obj.function)) {
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


