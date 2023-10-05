package io.apibrew.client.model.logic;

import java.util.Objects;
import io.apibrew.client.EntityInfo;
import io.apibrew.client.Entity;
import com.fasterxml.jackson.annotation.JsonValue;
import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class FunctionExecution extends Entity {
    
    private java.util.UUID id;
    
    private Function function;
    
    private Object input;
    
    private Object output;
    
    private Object error;
    
    private FunctionExecution.Status status;
    
    private java.util.Map<String, String> annotations;
    
    private int version;

    public static final String NAMESPACE = "logic";
    public static final String RESOURCE = "FunctionExecution";

    @JsonIgnore
    public static final EntityInfo<FunctionExecution> entityInfo = new EntityInfo<>("logic", "FunctionExecution", FunctionExecution.class, "logic-functionexecution");


    public static enum Status {
        PENDING("pending"),
        SUCCESS("success"),
        ERROR("error");

        private final String value;

        Status(String value) {
            this.value = value;
        }

        @JsonValue
        public String getValue() {
            return value;
        }
    }

    public FunctionExecution() {
    }

    public java.util.UUID getId() {
        return id;
    }

    public void setId(java.util.UUID id) {
        this.id = id;
    }

    public FunctionExecution withId(java.util.UUID id) {
        this.id = id;

        return this;
    }
    public Function getFunction() {
        return function;
    }

    public void setFunction(Function function) {
        this.function = function;
    }

    public FunctionExecution withFunction(Function function) {
        this.function = function;

        return this;
    }
    public Object getInput() {
        return input;
    }

    public void setInput(Object input) {
        this.input = input;
    }

    public FunctionExecution withInput(Object input) {
        this.input = input;

        return this;
    }
    public Object getOutput() {
        return output;
    }

    public void setOutput(Object output) {
        this.output = output;
    }

    public FunctionExecution withOutput(Object output) {
        this.output = output;

        return this;
    }
    public Object getError() {
        return error;
    }

    public void setError(Object error) {
        this.error = error;
    }

    public FunctionExecution withError(Object error) {
        this.error = error;

        return this;
    }
    public FunctionExecution.Status getStatus() {
        return status;
    }

    public void setStatus(FunctionExecution.Status status) {
        this.status = status;
    }

    public FunctionExecution withStatus(FunctionExecution.Status status) {
        this.status = status;

        return this;
    }
    public java.util.Map<String, String> getAnnotations() {
        return annotations;
    }

    public void setAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;
    }

    public FunctionExecution withAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;

        return this;
    }
    public int getVersion() {
        return version;
    }

    public void setVersion(int version) {
        this.version = version;
    }

    public FunctionExecution withVersion(int version) {
        this.version = version;

        return this;
    }

    @Override
    public boolean equals(Object o) {
        if (!(o instanceof FunctionExecution)) {
            return false;
        }

        FunctionExecution obj = (FunctionExecution) o;

        if (!Objects.equals(this.id, obj.id)) {
            return false;
        }
        if (!Objects.equals(this.function, obj.function)) {
            return false;
        }
        if (!Objects.equals(this.input, obj.input)) {
            return false;
        }
        if (!Objects.equals(this.output, obj.output)) {
            return false;
        }
        if (!Objects.equals(this.error, obj.error)) {
            return false;
        }
        if (!Objects.equals(this.status, obj.status)) {
            return false;
        }
        if (!Objects.equals(this.annotations, obj.annotations)) {
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


