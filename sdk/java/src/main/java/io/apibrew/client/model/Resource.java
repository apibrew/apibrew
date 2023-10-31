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
public class Resource extends Entity {
    
    private java.util.UUID id;
    
    private int version;
    
    private Resource.AuditData auditData;
    
    private String name;
    
    private Namespace namespace;
    
    private boolean virtual;
    
    private java.util.List<Resource.Property> properties;
    
    private java.util.List<Resource.Index> indexes;
    
    private java.util.List<Resource.SubType> types;
    
    private boolean immutable;
    
    private boolean $abstract;
    
    private boolean checkReferences;
    
    private DataSource dataSource;
    
    private String entity;
    
    private String catalog;
    
    private String title;
    
    private String description;
    
    private java.util.Map<String, String> annotations;

    public static final String NAMESPACE = "system";
    public static final String RESOURCE = "Resource";

    @JsonIgnore
    public static final EntityInfo<Resource> entityInfo = new EntityInfo<>("system", "Resource", Resource.class, "resources");

    public static class Property {
        
        private String name;
        
        private Resource.Type type;
        
        private String typeRef;
        
        private boolean primary;
        
        private boolean required;
        
        private boolean unique;
        
        private boolean immutable;
        
        private int length;
        
        private Resource.Property item;
        
        private Resource.Reference reference;
        
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
        public Resource.Type getType() {
            return type;
        }

        public void setType(Resource.Type type) {
            this.type = type;
        }

        public Property withType(Resource.Type type) {
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
        public Resource.Property getItem() {
            return item;
        }

        public void setItem(Resource.Property item) {
            this.item = item;
        }

        public Property withItem(Resource.Property item) {
            this.item = item;

            return this;
        }
        public Resource.Reference getReference() {
            return reference;
        }

        public void setReference(Resource.Reference reference) {
            this.reference = reference;
        }

        public Property withReference(Resource.Reference reference) {
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

        @Override
        public boolean equals(Object o) {
            if (!(o instanceof Property)) {
                return false;
            }

            Property obj = (Property) o;

            if (!Objects.equals(this.name, obj.name)) {
                return false;
            }
            if (!Objects.equals(this.type, obj.type)) {
                return false;
            }
            if (!Objects.equals(this.typeRef, obj.typeRef)) {
                return false;
            }
            if (!Objects.equals(this.primary, obj.primary)) {
                return false;
            }
            if (!Objects.equals(this.required, obj.required)) {
                return false;
            }
            if (!Objects.equals(this.unique, obj.unique)) {
                return false;
            }
            if (!Objects.equals(this.immutable, obj.immutable)) {
                return false;
            }
            if (!Objects.equals(this.length, obj.length)) {
                return false;
            }
            if (!Objects.equals(this.item, obj.item)) {
                return false;
            }
            if (!Objects.equals(this.reference, obj.reference)) {
                return false;
            }
            if (!Objects.equals(this.defaultValue, obj.defaultValue)) {
                return false;
            }
            if (!Objects.equals(this.enumValues, obj.enumValues)) {
                return false;
            }
            if (!Objects.equals(this.exampleValue, obj.exampleValue)) {
                return false;
            }
            if (!Objects.equals(this.title, obj.title)) {
                return false;
            }
            if (!Objects.equals(this.description, obj.description)) {
                return false;
            }
            if (!Objects.equals(this.annotations, obj.annotations)) {
                return false;
            }

            return true;
        }

        @Override
        public int hashCode() {
           return Objects.hash(name, type, typeRef, primary, required, unique, immutable, length, item, reference, defaultValue, enumValues, exampleValue, title, description, annotations);
        }
    }
    public static class SubType {
        
        private String name;
        
        private String title;
        
        private String description;
        
        private java.util.List<Resource.Property> properties;

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
        public java.util.List<Resource.Property> getProperties() {
            return properties;
        }

        public void setProperties(java.util.List<Resource.Property> properties) {
            this.properties = properties;
        }

        public SubType withProperties(java.util.List<Resource.Property> properties) {
            this.properties = properties;

            return this;
        }

        @Override
        public boolean equals(Object o) {
            if (!(o instanceof SubType)) {
                return false;
            }

            SubType obj = (SubType) o;

            if (!Objects.equals(this.name, obj.name)) {
                return false;
            }
            if (!Objects.equals(this.title, obj.title)) {
                return false;
            }
            if (!Objects.equals(this.description, obj.description)) {
                return false;
            }
            if (!Objects.equals(this.properties, obj.properties)) {
                return false;
            }

            return true;
        }

        @Override
        public int hashCode() {
           return Objects.hash(name, title, description, properties);
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
    public static class IndexProperty {
        
        private String name;
        
        private Resource.Order order;

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }

        public IndexProperty withName(String name) {
            this.name = name;

            return this;
        }
        public Resource.Order getOrder() {
            return order;
        }

        public void setOrder(Resource.Order order) {
            this.order = order;
        }

        public IndexProperty withOrder(Resource.Order order) {
            this.order = order;

            return this;
        }

        @Override
        public boolean equals(Object o) {
            if (!(o instanceof IndexProperty)) {
                return false;
            }

            IndexProperty obj = (IndexProperty) o;

            if (!Objects.equals(this.name, obj.name)) {
                return false;
            }
            if (!Objects.equals(this.order, obj.order)) {
                return false;
            }

            return true;
        }

        @Override
        public int hashCode() {
           return Objects.hash(name, order);
        }
    }
    public static class Index {
        
        private java.util.List<Resource.IndexProperty> properties;
        
        private Resource.IndexType indexType;
        
        private Boolean unique;
        
        private java.util.Map<String, String> annotations;

        public java.util.List<Resource.IndexProperty> getProperties() {
            return properties;
        }

        public void setProperties(java.util.List<Resource.IndexProperty> properties) {
            this.properties = properties;
        }

        public Index withProperties(java.util.List<Resource.IndexProperty> properties) {
            this.properties = properties;

            return this;
        }
        public Resource.IndexType getIndexType() {
            return indexType;
        }

        public void setIndexType(Resource.IndexType indexType) {
            this.indexType = indexType;
        }

        public Index withIndexType(Resource.IndexType indexType) {
            this.indexType = indexType;

            return this;
        }
        public Boolean getUnique() {
            return unique;
        }

        public void setUnique(Boolean unique) {
            this.unique = unique;
        }

        public Index withUnique(Boolean unique) {
            this.unique = unique;

            return this;
        }
        public java.util.Map<String, String> getAnnotations() {
            return annotations;
        }

        public void setAnnotations(java.util.Map<String, String> annotations) {
            this.annotations = annotations;
        }

        public Index withAnnotations(java.util.Map<String, String> annotations) {
            this.annotations = annotations;

            return this;
        }

        @Override
        public boolean equals(Object o) {
            if (!(o instanceof Index)) {
                return false;
            }

            Index obj = (Index) o;

            if (!Objects.equals(this.properties, obj.properties)) {
                return false;
            }
            if (!Objects.equals(this.indexType, obj.indexType)) {
                return false;
            }
            if (!Objects.equals(this.unique, obj.unique)) {
                return false;
            }
            if (!Objects.equals(this.annotations, obj.annotations)) {
                return false;
            }

            return true;
        }

        @Override
        public int hashCode() {
           return Objects.hash(properties, indexType, unique, annotations);
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

        @Override
        public boolean equals(Object o) {
            if (!(o instanceof Reference)) {
                return false;
            }

            Reference obj = (Reference) o;

            if (!Objects.equals(this.resource, obj.resource)) {
                return false;
            }
            if (!Objects.equals(this.cascade, obj.cascade)) {
                return false;
            }
            if (!Objects.equals(this.backReference, obj.backReference)) {
                return false;
            }

            return true;
        }

        @Override
        public int hashCode() {
           return Objects.hash(resource, cascade, backReference);
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
    public static enum Order {
        UNKNOWN("UNKNOWN"),
        ASC("ASC"),
        DESC("DESC");

        private final String value;

        Order(String value) {
            this.value = value;
        }

        @JsonValue
        public String getValue() {
            return value;
        }
    }
    public static enum IndexType {
        BTREE("BTREE"),
        HASH("HASH");

        private final String value;

        IndexType(String value) {
            this.value = value;
        }

        @JsonValue
        public String getValue() {
            return value;
        }
    }

    

    public Resource() {
    }

    public java.util.UUID getId() {
        return id;
    }

    public void setId(java.util.UUID id) {
        this.id = id;
    }

    public Resource withId(java.util.UUID id) {
        this.id = id;

        return this;
    }
    public int getVersion() {
        return version;
    }

    public void setVersion(int version) {
        this.version = version;
    }

    public Resource withVersion(int version) {
        this.version = version;

        return this;
    }
    public Resource.AuditData getAuditData() {
        return auditData;
    }

    public void setAuditData(Resource.AuditData auditData) {
        this.auditData = auditData;
    }

    public Resource withAuditData(Resource.AuditData auditData) {
        this.auditData = auditData;

        return this;
    }
    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Resource withName(String name) {
        this.name = name;

        return this;
    }
    public Namespace getNamespace() {
        return namespace;
    }

    public void setNamespace(Namespace namespace) {
        this.namespace = namespace;
    }

    public Resource withNamespace(Namespace namespace) {
        this.namespace = namespace;

        return this;
    }
    public boolean getVirtual() {
        return virtual;
    }

    public void setVirtual(boolean virtual) {
        this.virtual = virtual;
    }

    public Resource withVirtual(boolean virtual) {
        this.virtual = virtual;

        return this;
    }
    public java.util.List<Resource.Property> getProperties() {
        return properties;
    }

    public void setProperties(java.util.List<Resource.Property> properties) {
        this.properties = properties;
    }

    public Resource withProperties(java.util.List<Resource.Property> properties) {
        this.properties = properties;

        return this;
    }
    public java.util.List<Resource.Index> getIndexes() {
        return indexes;
    }

    public void setIndexes(java.util.List<Resource.Index> indexes) {
        this.indexes = indexes;
    }

    public Resource withIndexes(java.util.List<Resource.Index> indexes) {
        this.indexes = indexes;

        return this;
    }
    public java.util.List<Resource.SubType> getTypes() {
        return types;
    }

    public void setTypes(java.util.List<Resource.SubType> types) {
        this.types = types;
    }

    public Resource withTypes(java.util.List<Resource.SubType> types) {
        this.types = types;

        return this;
    }
    public boolean getImmutable() {
        return immutable;
    }

    public void setImmutable(boolean immutable) {
        this.immutable = immutable;
    }

    public Resource withImmutable(boolean immutable) {
        this.immutable = immutable;

        return this;
    }
    public boolean getAbstract() {
        return $abstract;
    }

    public void setAbstract(boolean $abstract) {
        this.$abstract = $abstract;
    }

    public Resource withAbstract(boolean $abstract) {
        this.$abstract = $abstract;

        return this;
    }
    public boolean getCheckReferences() {
        return checkReferences;
    }

    public void setCheckReferences(boolean checkReferences) {
        this.checkReferences = checkReferences;
    }

    public Resource withCheckReferences(boolean checkReferences) {
        this.checkReferences = checkReferences;

        return this;
    }
    public DataSource getDataSource() {
        return dataSource;
    }

    public void setDataSource(DataSource dataSource) {
        this.dataSource = dataSource;
    }

    public Resource withDataSource(DataSource dataSource) {
        this.dataSource = dataSource;

        return this;
    }
    public String getEntity() {
        return entity;
    }

    public void setEntity(String entity) {
        this.entity = entity;
    }

    public Resource withEntity(String entity) {
        this.entity = entity;

        return this;
    }
    public String getCatalog() {
        return catalog;
    }

    public void setCatalog(String catalog) {
        this.catalog = catalog;
    }

    public Resource withCatalog(String catalog) {
        this.catalog = catalog;

        return this;
    }
    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public Resource withTitle(String title) {
        this.title = title;

        return this;
    }
    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public Resource withDescription(String description) {
        this.description = description;

        return this;
    }
    public java.util.Map<String, String> getAnnotations() {
        return annotations;
    }

    public void setAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;
    }

    public Resource withAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;

        return this;
    }

    @Override
    public boolean equals(Object o) {
        if (!(o instanceof Resource)) {
            return false;
        }

        Resource obj = (Resource) o;

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
        if (!Objects.equals(this.namespace, obj.namespace)) {
            return false;
        }
        if (!Objects.equals(this.virtual, obj.virtual)) {
            return false;
        }
        if (!Objects.equals(this.properties, obj.properties)) {
            return false;
        }
        if (!Objects.equals(this.indexes, obj.indexes)) {
            return false;
        }
        if (!Objects.equals(this.types, obj.types)) {
            return false;
        }
        if (!Objects.equals(this.immutable, obj.immutable)) {
            return false;
        }
        if (!Objects.equals(this.$abstract, obj.$abstract)) {
            return false;
        }
        if (!Objects.equals(this.checkReferences, obj.checkReferences)) {
            return false;
        }
        if (!Objects.equals(this.dataSource, obj.dataSource)) {
            return false;
        }
        if (!Objects.equals(this.entity, obj.entity)) {
            return false;
        }
        if (!Objects.equals(this.catalog, obj.catalog)) {
            return false;
        }
        if (!Objects.equals(this.title, obj.title)) {
            return false;
        }
        if (!Objects.equals(this.description, obj.description)) {
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


