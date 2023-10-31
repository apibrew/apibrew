package io.apibrew.client;

import io.apibrew.client.model.Resource;

import static org.apache.logging.log4j.util.Strings.isBlank;

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

    public static EntityInfo<GenericRecord> fromResource(Resource resource) {
        return new EntityInfo<>(
                resource.getNamespace().getName(),
                resource.getName(),
                GenericRecord.class,
                getRestPath(resource)
        );
    }

    private static String getRestPath(Resource resource) {
        if (resource.getAnnotations() != null && resource.getAnnotations().get("OpenApiRestPath") != null) {
            return resource.getAnnotations().get("OpenApiRestPath");
        } else if (isBlank(resource.getNamespace().getName()) || resource.getNamespace().getName().equals("default")) {
            return slug(resource.getName());
        } else {
            return slug(resource.getNamespace().getName() + "/" + resource.getName());
        }
    }

    private static String slug(String name) {
        return name.toLowerCase().replaceAll("[^a-z0-9]+", "-");
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
