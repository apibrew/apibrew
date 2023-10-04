package io.apibrew.client.model.logic;

import java.util.Objects;
import io.apibrew.client.EntityInfo;
import io.apibrew.client.Entity;
import com.fasterxml.jackson.annotation.JsonValue;
import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class Function extends Entity {
    
    private java.util.UUID id;
    
    private String $package;
    
    private String name;
    
    private String script;
    
    private Module module;
    
    private FunctionExecutionEngine engine;
    
    private Function.Options options;
    
    private java.util.List<Function.Argument> args;
    
    private java.util.Map<String, String> annotations;
    
    private String createdBy;
    
    private String updatedBy;
    @JsonFormat(shape = JsonFormat.Shape.STRING, timezone = "UTC")
    private java.time.Instant createdOn;
    @JsonFormat(shape = JsonFormat.Shape.STRING, timezone = "UTC")
    private java.time.Instant updatedOn;
    
    private int version;

    @JsonIgnore
    public static final EntityInfo<Function> entityInfo = new EntityInfo<>("logic", "Function", Function.class, "logic-function");

    public static class Options {
        
        private Boolean namedParams;

        public Boolean getNamedParams() {
            return namedParams;
        }

        public void setNamedParams(Boolean namedParams) {
            this.namedParams = namedParams;
        }

        public Options withNamedParams(Boolean namedParams) {
            this.namedParams = namedParams;

            return this;
        }
    }
    public static class Argument {
        
        private String name;
        
        private String label;

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }

        public Argument withName(String name) {
            this.name = name;

            return this;
        }
        public String getLabel() {
            return label;
        }

        public void setLabel(String label) {
            this.label = label;
        }

        public Argument withLabel(String label) {
            this.label = label;

            return this;
        }
    }


    public Function() {
    }

    public java.util.UUID getId() {
        return id;
    }

    public void setId(java.util.UUID id) {
        this.id = id;
    }

    public Function withId(java.util.UUID id) {
        this.id = id;

        return this;
    }
    public String getPackage() {
        return $package;
    }

    public void setPackage(String $package) {
        this.$package = $package;
    }

    public Function withPackage(String $package) {
        this.$package = $package;

        return this;
    }
    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Function withName(String name) {
        this.name = name;

        return this;
    }
    public String getScript() {
        return script;
    }

    public void setScript(String script) {
        this.script = script;
    }

    public Function withScript(String script) {
        this.script = script;

        return this;
    }
    public Module getModule() {
        return module;
    }

    public void setModule(Module module) {
        this.module = module;
    }

    public Function withModule(Module module) {
        this.module = module;

        return this;
    }
    public FunctionExecutionEngine getEngine() {
        return engine;
    }

    public void setEngine(FunctionExecutionEngine engine) {
        this.engine = engine;
    }

    public Function withEngine(FunctionExecutionEngine engine) {
        this.engine = engine;

        return this;
    }
    public Function.Options getOptions() {
        return options;
    }

    public void setOptions(Function.Options options) {
        this.options = options;
    }

    public Function withOptions(Function.Options options) {
        this.options = options;

        return this;
    }
    public java.util.List<Function.Argument> getArgs() {
        return args;
    }

    public void setArgs(java.util.List<Function.Argument> args) {
        this.args = args;
    }

    public Function withArgs(java.util.List<Function.Argument> args) {
        this.args = args;

        return this;
    }
    public java.util.Map<String, String> getAnnotations() {
        return annotations;
    }

    public void setAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;
    }

    public Function withAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;

        return this;
    }
    public String getCreatedBy() {
        return createdBy;
    }

    public void setCreatedBy(String createdBy) {
        this.createdBy = createdBy;
    }

    public Function withCreatedBy(String createdBy) {
        this.createdBy = createdBy;

        return this;
    }
    public String getUpdatedBy() {
        return updatedBy;
    }

    public void setUpdatedBy(String updatedBy) {
        this.updatedBy = updatedBy;
    }

    public Function withUpdatedBy(String updatedBy) {
        this.updatedBy = updatedBy;

        return this;
    }
    public java.time.Instant getCreatedOn() {
        return createdOn;
    }

    public void setCreatedOn(java.time.Instant createdOn) {
        this.createdOn = createdOn;
    }

    public Function withCreatedOn(java.time.Instant createdOn) {
        this.createdOn = createdOn;

        return this;
    }
    public java.time.Instant getUpdatedOn() {
        return updatedOn;
    }

    public void setUpdatedOn(java.time.Instant updatedOn) {
        this.updatedOn = updatedOn;
    }

    public Function withUpdatedOn(java.time.Instant updatedOn) {
        this.updatedOn = updatedOn;

        return this;
    }
    public int getVersion() {
        return version;
    }

    public void setVersion(int version) {
        this.version = version;
    }

    public Function withVersion(int version) {
        this.version = version;

        return this;
    }

    @Override
    public boolean equals(Object o) {
        if (!(o instanceof Function)) {
            return false;
        }

        Function obj = (Function) o;

        if (!Objects.equals(this.id, obj.id)) {
            return false;
        }
        if (!Objects.equals(this.$package, obj.$package)) {
            return false;
        }
        if (!Objects.equals(this.name, obj.name)) {
            return false;
        }
        if (!Objects.equals(this.script, obj.script)) {
            return false;
        }
        if (!Objects.equals(this.module, obj.module)) {
            return false;
        }
        if (!Objects.equals(this.engine, obj.engine)) {
            return false;
        }
        if (!Objects.equals(this.options, obj.options)) {
            return false;
        }
        if (!Objects.equals(this.args, obj.args)) {
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


