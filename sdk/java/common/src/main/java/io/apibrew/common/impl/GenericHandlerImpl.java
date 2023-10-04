package io.apibrew.common.impl;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import io.apibrew.client.Client;
import io.apibrew.client.Entity;
import io.apibrew.client.EntityInfo;
import io.apibrew.common.ExtensionInfo;
import io.apibrew.common.ext.*;
import io.apibrew.client.model.Record;

import java.util.ArrayList;
import java.util.List;
import java.util.function.BiFunction;
import java.util.function.BiPredicate;
import java.util.function.Function;

import static io.apibrew.client.model.Extension.*;

public class GenericHandlerImpl implements GenericHandler {
    private final Client client;
    private final ExtensionService extensionService;

    private final ExtensionInfo extensionInfo;
    private final List<BiPredicate<Event, Entity>> predicates;

    private final ObjectMapper objectMapper = new ObjectMapper();

    public GenericHandlerImpl(Client client, ExtensionService extensionService) {
        this(
                client,
                extensionService,
                new ExtensionInfo(),
                new ArrayList<>()
        );
    }

    public GenericHandlerImpl(Client client, ExtensionService extensionService, ExtensionInfo extensionInfo, List<BiPredicate<Event, Entity>> predicates) {
        this.client = client;
        this.extensionService = extensionService;
        this.extensionInfo = extensionInfo.withSealResource(false);
        this.predicates = predicates;

        objectMapper.registerModule(new JavaTimeModule());
        objectMapper.disable(SerializationFeature.WRITE_DATES_AS_TIMESTAMPS);
    }

    public GenericHandler withExtensionInfo(ExtensionInfo extensionInfo) {
        return new GenericHandlerImpl(client, extensionService, extensionInfo, predicates);
    }

    public GenericHandler withPredicates(List<BiPredicate<Event, Entity>> predicates) {
        return new GenericHandlerImpl(client, extensionService, extensionInfo, predicates);
    }

    public GenericHandler withPredicate(BiPredicate<Event, Entity> predicate) {
        ArrayList<BiPredicate<Event, Entity>> predicatesCopy = new ArrayList<>(predicates);
        predicatesCopy.add(predicate);

        return withPredicates(predicatesCopy);
    }

    @Override
    public <T extends Entity> GenericHandler when(Class<T> entityClass, Condition<T> condition) {
        return ((GenericHandlerImpl) withExtensionInfo(condition.configureExtensionInfo(extensionInfo)))
                .withPredicate((event, item) -> {
                    if (entityClass.isInstance(item)) {
                        return condition.eventMatches(event, entityClass.cast(item));
                    } else {
                        return false;
                    }
                });
    }

    @Override
    public GenericHandler configure(Function<ExtensionInfo, ExtensionInfo> configurer) {
        return withExtensionInfo(configurer.apply(extensionInfo));
    }

    private <T extends Entity> T recordToEntity(Class<T> entityClass, Record entity) {
        if (entity == null) {
            return null;
        }

        return objectMapper.convertValue(entity.getProperties(), entityClass);
    }

    private Record recordFromEntity(Entity entity) {
        if (entity == null) {
            return null;
        }

        Record record = new Record();
        record.setProperties(objectMapper.convertValue(entity, Object.class));
        record.setId(entity.getId());

        return record;
    }

    private boolean checkPredicates(Event event, Entity entity) {
        for (BiPredicate<Event, Entity> predicate : predicates) {
            if (!predicate.test(event, entity)) {
                return false;
            }
        }

        return true;
    }

    @Override
    public <T extends Entity> GenericHandler operate(Class<T> entityClass, Operator<T> entityOperator) {
        entityOperator.operate(this.toHandler(entityClass));

        return this;
    }

    private <T extends Entity> Handler<T> toHandler(Class<T> entityClass) {
        List<BiPredicate<Event, T>> castedPredicates = new ArrayList<>();

        for (BiPredicate<Event, Entity> predicate : predicates) {
            castedPredicates.add((event, item) -> {
                if (entityClass.isInstance(item)) {
                    return predicate.test(event, entityClass.cast(item));
                } else {
                    return false;
                }
            });
        }

        return new HandlerImpl<>(client, extensionService, EntityInfo.<T>fromEntityClass(entityClass), extensionInfo, castedPredicates);
    }

    @Override
    public <T extends Entity> GenericHandler operate(Class<T> entityClass, BiFunction<Event, T, T> entityOperator) {
        extensionService.registerExtensionWithOperator(extensionInfo, (event, record) -> {
            T castedEntity = recordToEntity(entityClass, record);
            if (!checkPredicates(event, castedEntity)) {
                return record;
            }

            T result = entityOperator.apply(event, castedEntity);

            return recordFromEntity(result);
        });

        return this;
    }
}
