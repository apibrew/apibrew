package io.apibrew.lib.model;

import java.util.Objects;
import io.apibrew.lib.EntityInfo;
import io.apibrew.lib.Entity;

public class Resource extends Entity {
    private java.util.UUID id;
    private int version;
    private String createdBy;
    private String updatedBy;
    private java.time.LocalDateTime createdOn;
    private java.time.LocalDateTime updatedOn;
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

    public static final EntityInfo<Resource> entityInfo = new EntityInfo<>("system", "Resource", Resource.class);

    public static class Property {
        private String name;
        private Resource.Type type;
        private String typeRef;
        private String mapping;
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
        public Resource.Type getType() {
            return type;
        }

        public void setType(Resource.Type type) {
            this.type = type;
        }
        public String getTypeRef() {
            return typeRef;
        }

        public void setTypeRef(String typeRef) {
            this.typeRef = typeRef;
        }
        public String getMapping() {
            return mapping;
        }

        public void setMapping(String mapping) {
            this.mapping = mapping;
        }
        public boolean getPrimary() {
            return primary;
        }

        public void setPrimary(boolean primary) {
            this.primary = primary;
        }
        public boolean getRequired() {
            return required;
        }

        public void setRequired(boolean required) {
            this.required = required;
        }
        public boolean getUnique() {
            return unique;
        }

        public void setUnique(boolean unique) {
            this.unique = unique;
        }
        public boolean getImmutable() {
            return immutable;
        }

        public void setImmutable(boolean immutable) {
            this.immutable = immutable;
        }
        public int getLength() {
            return length;
        }

        public void setLength(int length) {
            this.length = length;
        }
        public Resource.Property getItem() {
            return item;
        }

        public void setItem(Resource.Property item) {
            this.item = item;
        }
        public Resource.Reference getReference() {
            return reference;
        }

        public void setReference(Resource.Reference reference) {
            this.reference = reference;
        }
        public Object getDefaultValue() {
            return defaultValue;
        }

        public void setDefaultValue(Object defaultValue) {
            this.defaultValue = defaultValue;
        }
        public java.util.List<String> getEnumValues() {
            return enumValues;
        }

        public void setEnumValues(java.util.List<String> enumValues) {
            this.enumValues = enumValues;
        }
        public Object getExampleValue() {
            return exampleValue;
        }

        public void setExampleValue(Object exampleValue) {
            this.exampleValue = exampleValue;
        }
        public String getTitle() {
            return title;
        }

        public void setTitle(String title) {
            this.title = title;
        }
        public String getDescription() {
            return description;
        }

        public void setDescription(String description) {
            this.description = description;
        }
        public java.util.Map<String, String> getAnnotations() {
            return annotations;
        }

        public void setAnnotations(java.util.Map<String, String> annotations) {
            this.annotations = annotations;
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
        public String getTitle() {
            return title;
        }

        public void setTitle(String title) {
            this.title = title;
        }
        public String getDescription() {
            return description;
        }

        public void setDescription(String description) {
            this.description = description;
        }
        public java.util.List<Resource.Property> getProperties() {
            return properties;
        }

        public void setProperties(java.util.List<Resource.Property> properties) {
            this.properties = properties;
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
        public Resource.Order getOrder() {
            return order;
        }

        public void setOrder(Resource.Order order) {
            this.order = order;
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
        public Resource.IndexType getIndexType() {
            return indexType;
        }

        public void setIndexType(Resource.IndexType indexType) {
            this.indexType = indexType;
        }
        public Boolean getUnique() {
            return unique;
        }

        public void setUnique(Boolean unique) {
            this.unique = unique;
        }
        public java.util.Map<String, String> getAnnotations() {
            return annotations;
        }

        public void setAnnotations(java.util.Map<String, String> annotations) {
            this.annotations = annotations;
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
        public Boolean getCascade() {
            return cascade;
        }

        public void setCascade(Boolean cascade) {
            this.cascade = cascade;
        }
        public String getBackReference() {
            return backReference;
        }

        public void setBackReference(String backReference) {
            this.backReference = backReference;
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

        public String getValue() {
            return value;
        }
    }

    public Resource() {
    }

    public EntityInfo<Resource> getEntityInfo() {
        return entityInfo;
    }

    public java.util.UUID getId() {
        return id;
    }

    public void setId(java.util.UUID id) {
        this.id = id;
    }
    public int getVersion() {
        return version;
    }

    public void setVersion(int version) {
        this.version = version;
    }
    public String getCreatedBy() {
        return createdBy;
    }

    public void setCreatedBy(String createdBy) {
        this.createdBy = createdBy;
    }
    public String getUpdatedBy() {
        return updatedBy;
    }

    public void setUpdatedBy(String updatedBy) {
        this.updatedBy = updatedBy;
    }
    public java.time.LocalDateTime getCreatedOn() {
        return createdOn;
    }

    public void setCreatedOn(java.time.LocalDateTime createdOn) {
        this.createdOn = createdOn;
    }
    public java.time.LocalDateTime getUpdatedOn() {
        return updatedOn;
    }

    public void setUpdatedOn(java.time.LocalDateTime updatedOn) {
        this.updatedOn = updatedOn;
    }
    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
    public Namespace getNamespace() {
        return namespace;
    }

    public void setNamespace(Namespace namespace) {
        this.namespace = namespace;
    }
    public boolean getVirtual() {
        return virtual;
    }

    public void setVirtual(boolean virtual) {
        this.virtual = virtual;
    }
    public java.util.List<Resource.Property> getProperties() {
        return properties;
    }

    public void setProperties(java.util.List<Resource.Property> properties) {
        this.properties = properties;
    }
    public java.util.List<Resource.Index> getIndexes() {
        return indexes;
    }

    public void setIndexes(java.util.List<Resource.Index> indexes) {
        this.indexes = indexes;
    }
    public java.util.List<Resource.SubType> getTypes() {
        return types;
    }

    public void setTypes(java.util.List<Resource.SubType> types) {
        this.types = types;
    }
    public boolean getImmutable() {
        return immutable;
    }

    public void setImmutable(boolean immutable) {
        this.immutable = immutable;
    }
    public boolean getAbstract() {
        return $abstract;
    }

    public void setAbstract(boolean $abstract) {
        this.$abstract = $abstract;
    }
    public boolean getCheckReferences() {
        return checkReferences;
    }

    public void setCheckReferences(boolean checkReferences) {
        this.checkReferences = checkReferences;
    }
    public DataSource getDataSource() {
        return dataSource;
    }

    public void setDataSource(DataSource dataSource) {
        this.dataSource = dataSource;
    }
    public String getEntity() {
        return entity;
    }

    public void setEntity(String entity) {
        this.entity = entity;
    }
    public String getCatalog() {
        return catalog;
    }

    public void setCatalog(String catalog) {
        this.catalog = catalog;
    }
    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }
    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }
    public java.util.Map<String, String> getAnnotations() {
        return annotations;
    }

    public void setAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;
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


