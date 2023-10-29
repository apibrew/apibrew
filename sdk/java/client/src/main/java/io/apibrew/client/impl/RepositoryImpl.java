package io.apibrew.client.impl;

import io.apibrew.client.*;
import io.apibrew.client.model.Extension;
import kong.unirest.HttpResponse;
import kong.unirest.Unirest;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.List;
import java.util.concurrent.atomic.AtomicBoolean;
import java.util.function.Consumer;

import static io.apibrew.client.helper.EventHelper.shortInfo;

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
        return client.createRecord(entityInfo, record);
    }

    @Override
    public T get(String id) {
        return client.getRecord(entityInfo, id);
    }

    @Override
    public T update(T record) {
        return client.updateRecord(entityInfo, record);
    }

    @Override
    public T delete(String id) {
        return client.deleteRecord(entityInfo, id);
    }

    @Override
    public T apply(T record) {
        return client.applyRecord(entityInfo, record);
    }

    @Override
    public Container<T> list() {
        return client.listRecords(entityInfo);
    }

    @Override
    public T load(T record) {
        return client.loadRecord(entityInfo, record);
    }

    @Override
    public Container<T> list(Extension.BooleanExpression query) {
        return client.listRecords(entityInfo, query);
    }

    @Override
    public Watcher<T> watch(Consumer<Extension.Event> eventConsumer) {
        return new WatcherImpl<>(client, entityInfo, eventConsumer);
    }
}
