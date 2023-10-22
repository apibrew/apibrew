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
public class ResourceAction extends Entity {
    
    private java.util.UUID id;
    
    private int version;
    
    private ResourceAction.AuditData auditData;
    
    private Resource resource;
    
    private String name;
    
    private String title;
    
    private String description;
    
    private boolean internal;
    
    private java.util.List<ResourceAction.SubType> types;
    
    private java.util.List<ResourceAction.Property> input;
    
    private ResourceAction.Property output;
    
    private java.util.Map<String, String> annotations;

    public static final String NAMESPACE = "system";
    public static final String RESOURCE = "ResourceAction";

    @JsonIgnore
    public static final EntityInfo<ResourceAction> entityInfo = new EntityInfo<>("system", "ResourceAction", ResourceAction.class, "system-resourceaction");

    public static class SubType {
        
        private String name;
        
        private String title;
        
        private String description;
        
        private java.util.List<ResourceAction.Property> properties;

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }

        public SubType withName(String name) {
            this.name = name;

            return this;
        }
        public String getTitle() {
            return title;
        }

        public void setTitle(String title) {
            this.title = title;
        }

        public SubType withTitle(String title) {
            this.title = title;

            return this;
        }
        public String getDescription() {
            return description;
        }

        public void setDescription(String description) {
            this.description = description;
        }

        public SubType withDescription(String description) {
            this.description = description;

            return this;
        }
        public java.util.List<ResourceAction.Property> getProperties() {
            return properties;
        }

        public void setProperties(java.util.List<ResourceAction.Property> properties) {
            this.properties = properties;
        }

        public SubType withProperties(java.util.List<ResourceAction.Property> properties) {
            this.properties = properties;

            return this;
        }
    }
    public static class Reference {
        
        private Resource resource;
        
        private Boolean cascade;
        
        private String backReference;

        public Resource getResource() {
            return resource;
        }

        public void setResource(Resource resource) {
            this.resource = resource;
        }

        public Reference withResource(Resource resource) {
            this.resource = resource;

            return this;
        }
        public Boolean getCascade() {
            return cascade;
        }

        public void setCascade(Boolean cascade) {
            this.cascade = cascade;
        }

        public Reference withCascade(Boolean cascade) {
            this.cascade = cascade;

            return this;
        }
        public String getBackReference() {
            return backReference;
        }

        public void setBackReference(String backReference) {
            this.backReference = backReference;
        }

        public Reference withBackReference(String backReference) {
            this.backReference = backReference;

            return this;
        }
    }
    public static class Property {
        
        private String name;
        
        private ResourceAction.Type type;
        
        private String typeRef;
        
        private boolean primary;
        
        private boolean required;
        
        private boolean unique;
        
        private boolean immutable;
        
        private int length;
        
        private ResourceAction.Property item;
        
        private ResourceAction.Reference reference;
        
        private Object defaultValue;
        
        private java.util.List<String> enumValues;
        
        private Object exampleValue;
        
        private String title;
        
        private String description;
        
        private java.util.Map<String, String> annotations;

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }

        public Property withName(String name) {
            this.name = name;

            return this;
        }
        public ResourceAction.Type getType() {
            return type;
        }

        public void setType(ResourceAction.Type type) {
            this.type = type;
        }

        public Property withType(ResourceAction.Type type) {
            this.type = type;

            return this;
        }
        public String getTypeRef() {
            return typeRef;
        }

        public void setTypeRef(String typeRef) {
            this.typeRef = typeRef;
        }

        public Property withTypeRef(String typeRef) {
            this.typeRef = typeRef;

            return this;
        }
        public boolean getPrimary() {
            return primary;
        }

        public void setPrimary(boolean primary) {
            this.primary = primary;
        }

        public Property withPrimary(boolean primary) {
            this.primary = primary;

            return this;
        }
        public boolean getRequired() {
            return required;
        }

        public void setRequired(boolean required) {
            this.required = required;
        }

        public Property withRequired(boolean required) {
            this.required = required;

            return this;
        }
        public boolean getUnique() {
            return unique;
        }

        public void setUnique(boolean unique) {
            this.unique = unique;
        }

        public Property withUnique(boolean unique) {
            this.unique = unique;

            return this;
        }
        public boolean getImmutable() {
            return immutable;
        }

        public void setImmutable(boolean immutable) {
            this.immutable = immutable;
        }

        public Property withImmutable(boolean immutable) {
            this.immutable = immutable;

            return this;
        }
        public int getLength() {
            return length;
        }

        public void setLength(int length) {
            this.length = length;
        }

        public Property withLength(int length) {
            this.length = length;

            return this;
        }
        public ResourceAction.Property getItem() {
            return item;
        }

        public void setItem(ResourceAction.Property item) {
            this.item = item;
        }

        public Property withItem(ResourceAction.Property item) {
            this.item = item;

            return this;
        }
        public ResourceAction.Reference getReference() {
            return reference;
        }

        public void setReference(ResourceAction.Reference reference) {
            this.reference = reference;
        }

        public Property withReference(ResourceAction.Reference reference) {
            this.reference = reference;

            return this;
        }
        public Object getDefaultValue() {
            return defaultValue;
        }

        public void setDefaultValue(Object defaultValue) {
            this.defaultValue = defaultValue;
        }

        public Property withDefaultValue(Object defaultValue) {
            this.defaultValue = defaultValue;

            return this;
        }
        public java.util.List<String> getEnumValues() {
            return enumValues;
        }

        public void setEnumValues(java.util.List<String> enumValues) {
            this.enumValues = enumValues;
        }

        public Property withEnumValues(java.util.List<String> enumValues) {
            this.enumValues = enumValues;

            return this;
        }
        public Object getExampleValue() {
            return exampleValue;
        }

        public void setExampleValue(Object exampleValue) {
            this.exampleValue = exampleValue;
        }

        public Property withExampleValue(Object exampleValue) {
            this.exampleValue = exampleValue;

            return this;
        }
        public String getTitle() {
            return title;
        }

        public void setTitle(String title) {
            this.title = title;
        }

        public Property withTitle(String title) {
            this.title = title;

            return this;
        }
        public String getDescription() {
            return description;
        }

        public void setDescription(String description) {
            this.description = description;
        }

        public Property withDescription(String description) {
            this.description = description;

            return this;
        }
        public java.util.Map<String, String> getAnnotations() {
            return annotations;
        }

        public void setAnnotations(java.util.Map<String, String> annotations) {
            this.annotations = annotations;
        }

        public Property withAnnotations(java.util.Map<String, String> annotations) {
            this.annotations = annotations;

            return this;
        }
    }
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

    public static enum Type {
        BOOL("BOOL"),
        STRING("STRING"),
        FLOAT32("FLOAT32"),
        FLOAT64("FLOAT64"),
        INT32("INT32"),
        INT64("INT64"),
        BYTES("BYTES"),
        UUID("UUID"),
        DATE("DATE"),
        TIME("TIME"),
        TIMESTAMP("TIMESTAMP"),
        OBJECT("OBJECT"),
        MAP("MAP"),
        LIST("LIST"),
        REFERENCE("REFERENCE"),
        ENUM("ENUM"),
        STRUCT("STRUCT");

        private final String value;

        Type(String value) {
            this.value = value;
        }

        @JsonValue
        public String getValue() {
            return value;
        }
    }

    

    public ResourceAction() {
    }

    public java.util.UUID getId() {
        return id;
    }

    public void setId(java.util.UUID id) {
        this.id = id;
    }

    public ResourceAction withId(java.util.UUID id) {
        this.id = id;

        return this;
    }
    public int getVersion() {
        return version;
    }

    public void setVersion(int version) {
        this.version = version;
    }

    public ResourceAction withVersion(int version) {
        this.version = version;

        return this;
    }
    public ResourceAction.AuditData getAuditData() {
        return auditData;
    }

    public void setAuditData(ResourceAction.AuditData auditData) {
        this.auditData = auditData;
    }

    public ResourceAction withAuditData(ResourceAction.AuditData auditData) {
        this.auditData = auditData;

        return this;
    }
    public Resource getResource() {
        return resource;
    }

    public void setResource(Resource resource) {
        this.resource = resource;
    }

    public ResourceAction withResource(Resource resource) {
        this.resource = resource;

        return this;
    }
    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public ResourceAction withName(String name) {
        this.name = name;

        return this;
    }
    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public ResourceAction withTitle(String title) {
        this.title = title;

        return this;
    }
    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public ResourceAction withDescription(String description) {
        this.description = description;

        return this;
    }
    public boolean getInternal() {
        return internal;
    }

    public void setInternal(boolean internal) {
        this.internal = internal;
    }

    public ResourceAction withInternal(boolean internal) {
        this.internal = internal;

        return this;
    }
    public java.util.List<ResourceAction.SubType> getTypes() {
        return types;
    }

    public void setTypes(java.util.List<ResourceAction.SubType> types) {
        this.types = types;
    }

    public ResourceAction withTypes(java.util.List<ResourceAction.SubType> types) {
        this.types = types;

        return this;
    }
    public java.util.List<ResourceAction.Property> getInput() {
        return input;
    }

    public void setInput(java.util.List<ResourceAction.Property> input) {
        this.input = input;
    }

    public ResourceAction withInput(java.util.List<ResourceAction.Property> input) {
        this.input = input;

        return this;
    }
    public ResourceAction.Property getOutput() {
        return output;
    }

    public void setOutput(ResourceAction.Property output) {
        this.output = output;
    }

    public ResourceAction withOutput(ResourceAction.Property output) {
        this.output = output;

        return this;
    }
    public java.util.Map<String, String> getAnnotations() {
        return annotations;
    }

    public void setAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;
    }

    public ResourceAction withAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;

        return this;
    }

    @Override
    public boolean equals(Object o) {
        if (!(o instanceof ResourceAction)) {
            return false;
        }

        ResourceAction obj = (ResourceAction) o;

        if (!Objects.equals(this.id, obj.id)) {
            return false;
        }
        if (!Objects.equals(this.version, obj.version)) {
            return false;
        }
        if (!Objects.equals(this.auditData, obj.auditData)) {
            return false;
        }
        if (!Objects.equals(this.resource, obj.resource)) {
            return false;
        }
        if (!Objects.equals(this.name, obj.name)) {
            return false;
        }
        if (!Objects.equals(this.title, obj.title)) {
            return false;
        }
        if (!Objects.equals(this.description, obj.description)) {
            return false;
        }
        if (!Objects.equals(this.internal, obj.internal)) {
            return false;
        }
        if (!Objects.equals(this.types, obj.types)) {
            return false;
        }
        if (!Objects.equals(this.input, obj.input)) {
            return false;
        }
        if (!Objects.equals(this.output, obj.output)) {
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


