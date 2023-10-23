package io.apibrew.common.impl;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import io.apibrew.client.Client;
import io.apibrew.client.Entity;
import io.apibrew.client.EntityInfo;
import io.apibrew.common.ExtensionInfo;
import io.apibrew.common.ext.Condition;
import io.apibrew.common.ext.ExtensionService;
import io.apibrew.common.ext.Handler;
import io.apibrew.common.ext.Operator;
import io.apibrew.client.model.Record;

import java.util.ArrayList;
import java.util.List;
import java.util.function.BiFunction;
import java.util.function.BiPredicate;
import java.util.function.Function;

import static io.apibrew.client.model.Extension.*;

public class HandlerImpl<T extends Entity> implements Handler<T> {
    private final Client client;
    private final ExtensionService extensionService;

    private final ExtensionInfo extensionInfo;
    private final List<BiPredicate<Event, T>> predicates;

    private final EntityInfo<T> entityInfo;

    private final ObjectMapper objectMapper = new ObjectMapper();

    public HandlerImpl(Client client, ExtensionService extensionService, EntityInfo<T> entityInfo) {
        this(
                client,
                extensionService,
                entityInfo,
                new ExtensionInfo().withNamespace(entityInfo.getNamespace()).withResource(entityInfo.getResource()),
                new ArrayList<>()
        );
    }

    public HandlerImpl(Client client, ExtensionService extensionService, Class<T> entityClass) {
        this(client, extensionService, EntityInfo.fromEntityClass(entityClass));
    }

    public HandlerImpl(Client client, ExtensionService extensionService, EntityInfo<T> entityInfo, ExtensionInfo extensionInfo, List<BiPredicate<Event, T>> predicates) {
        this.client = client;
        this.extensionService = extensionService;
        this.entityInfo = entityInfo;
        this.extensionInfo = extensionInfo.withSealResource(true);
        this.predicates = predicates;

        objectMapper.registerModule(new JavaTimeModule());
        objectMapper.disable(SerializationFeature.WRITE_DATES_AS_TIMESTAMPS);
    }

    public Handler<T> withExtensionInfo(ExtensionInfo extensionInfo) {
        return new HandlerImpl<>(client, extensionService, entityInfo, extensionInfo, predicates);
    }

    public Handler<T> withPredicates(List<BiPredicate<Event, T>> predicates) {
        return new HandlerImpl<>(client, extensionService, entityInfo, extensionInfo, predicates);
    }

    public Handler<T> withPredicate(BiPredicate<Event, T> predicate) {
        ArrayList<BiPredicate<Event, T>> predicatesCopy = new ArrayList<>(predicates);
        predicatesCopy.add(predicate);

        return withPredicates(predicatesCopy);
    }

    @Override
    public Handler<T> when(Condition<T> condition) {
        return ((HandlerImpl<T>) withExtensionInfo(condition.configureExtensionInfo(extensionInfo)))
                .withPredicate(condition::eventMatches);
    }

    @Override
    public Handler<T> configure(Function<ExtensionInfo, ExtensionInfo> configurer) {
        return withExtensionInfo(configurer.apply(extensionInfo));
    }

    private T recordToEntity(Record entity) {
        if (entity == null) {
            return null;
        }

        return objectMapper.convertValue(entity.getProperties(), entityInfo.getEntityClass());
    }

    private Record recordFromEntity(T entity) {
        if (entity == null) {
            return null;
        }

        Record record = new Record();
        record.setProperties(objectMapper.convertValue(entity, Object.class));
        record.setId(entity.getId());

        return record;
    }

    private boolean checkPredicates(Event event, T entity) {
        for (BiPredicate<Event, T> predicate : predicates) {
            if (!predicate.test(event, entity)) {
                return false;
            }
        }

        return true;
    }

    @Override
    public String operate(Operator<T> entityOperator) {
        return entityOperator.operate(this);
    }

    @Override
    public String operate(BiFunction<Event, T, T> entityOperator) {
        return extensionService.registerExtensionWithOperator(extensionInfo, (event, record) -> {
            T castedEntity = recordToEntity(record);
            if (!checkPredicates(event, castedEntity)) {
                return record;
            }

            T result = entityOperator.apply(event, castedEntity);

            return recordFromEntity(result);
        });
    }

    @Override
    public void unRegister(String id) {
        extensionService.unRegisterOperator(id);
    }
}
