package io.apibrew.common;

public final class EntityInfo<T> {
    private final String namespace;
    private final String resource;
    private final Class<T> entityClass;

    public EntityInfo(String namespace, String resource, Class<T> entityClass) {
        this.namespace = namespace;
        this.resource = resource;
        this.entityClass = entityClass;
    }

    public String getNamespace() {
        return namespace;
    }

    public String getResource() {
        return resource;
    }

    public Class<T> getEntityClass() {
        return entityClass;
    }
}
