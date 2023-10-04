package io.apibrew.client.impl;

import io.apibrew.client.*;
import io.apibrew.client.model.Extension;

public class RepositoryImpl<T extends Entity> implements Repository<T> {
    private final Client client;
    private final EntityInfo<T> entityInfo;

    public RepositoryImpl(ClientImpl client, Class<T> entityClass) {
        this.client = client;
        this.entityInfo = EntityInfo.fromEntityClass(entityClass);
    }

    public RepositoryImpl(ClientImpl client, EntityInfo<T> entityInfo) {
        this.client = client;
        this.entityInfo = entityInfo;
    }

    @Override
    public T create(T record) {
        return client.createRecord(entityInfo.getEntityClass(), entityInfo.getNamespace(), entityInfo.getResource(), record);
    }

    @Override
    public T get(String id) {
        return client.getRecord(entityInfo.getEntityClass(), entityInfo.getNamespace(), entityInfo.getResource(), id);
    }

    @Override
    public T update(T record) {
        return client.updateRecord(entityInfo.getEntityClass(), entityInfo.getNamespace(), entityInfo.getResource(), record);
    }

    @Override
    public T delete(String id) {
        return client.deleteRecord(entityInfo.getEntityClass(), entityInfo.getNamespace(), entityInfo.getResource(), id);
    }

    @Override
    public T apply(T record) {
        return client.applyRecord(entityInfo.getEntityClass(), entityInfo.getNamespace(), entityInfo.getResource(), record);
    }

    @Override
    public Container<T> list() {
        return client.listRecords(entityInfo.getEntityClass(), entityInfo.getNamespace(), entityInfo.getResource());
    }

    @Override
    public Container<T> list(Extension.BooleanExpression query) {
        return client.listRecords(entityInfo.getEntityClass(), entityInfo.getNamespace(), entityInfo.getResource(), query);
    }
}
