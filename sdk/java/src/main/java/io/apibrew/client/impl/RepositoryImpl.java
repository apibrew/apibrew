package io.apibrew.client.impl;

import io.apibrew.client.*;
import io.apibrew.client.model.Extension;

import java.util.List;
import java.util.Map;
import java.util.function.Consumer;

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
        return client.getRecord(entityInfo, GetRecordParams.builder().id(id).build());
    }

    @Override
    public T get(String id, List<String> resolveReferences) {
        return client.getRecord(entityInfo, GetRecordParams.builder().id(id).resolveReferences(resolveReferences).build());
    }

    @Override
    public T findBy(Map<String, String> filters, List<String> resolveReferences) {
        Container<T> result = list(ListRecordParams.builder()
                .filters(filters)
                .limit(1)
                .resolveReferences(resolveReferences)
                .build());

        if (result.getTotal() > 1) {
            throw new ApiException(Extension.Code.RECORD_VALIDATION_ERROR, "More than one record found");
        } else if (result.getTotal() == 1) {
            return result.getContent().get(0);
        } else {
            throw new ApiException(Extension.Code.RECORD_NOT_FOUND, "Record not found");
        }
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
        return client.listRecords(entityInfo, null);
    }

    @Override
    public Container<T> list(ListRecordParams params) {
        return client.listRecords(entityInfo, params);
    }

    @Override
    public T load(T record) {
        return client.loadRecord(entityInfo, record);
    }

    @Override
    public Container<T> list(Extension.BooleanExpression... query) {
        if (query.length == 0) {
            return client.listRecords(entityInfo, null);
        } else if (query.length == 1) {
            return client.listRecords(entityInfo, ListRecordParams.builder().query(query[0]).build());
        } else {
            return client.listRecords(entityInfo, ListRecordParams.builder().query(BooleanExpressionBuilder.and(query)).build());
        }
    }

    @Override
    public Watcher<T> watch(Consumer<Extension.Event> eventConsumer) {
        return new WatcherImpl<>(client, entityInfo, eventConsumer);
    }
}
