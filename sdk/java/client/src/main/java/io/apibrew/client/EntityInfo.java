package io.apibrew.client;

public final class EntityInfo<T> {
    private final String namespace;
    private final String resource;
    private final Class<T> entityClass;
    private final String restPath;

    public EntityInfo(String namespace, String resource, Class<T> entityClass, String restPath) {
        this.namespace = namespace;
        this.resource = resource;
        this.entityClass = entityClass;
        this.restPath = restPath;
    }

    @SuppressWarnings("unchecked")
    public static <T extends Entity> EntityInfo<T> fromEntityClass(Class<T> entityClass) {
        try {
            return (EntityInfo<T>) entityClass.getField("entityInfo").get(entityClass);
        } catch (IllegalAccessException | NoSuchFieldException e) {
            throw new UnsupportedOperationException(e);
        }
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

    public String getRestPath() {
        return restPath;
    }

    @Override
    public String toString() {
        return getNamespace() + "/" + getResource();
    }
}
