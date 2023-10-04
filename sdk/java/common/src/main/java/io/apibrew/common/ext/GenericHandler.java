package io.apibrew.common.ext;

import io.apibrew.client.Entity;
import io.apibrew.common.ExtensionInfo;
import io.apibrew.client.model.Extension;

import java.util.function.BiFunction;
import java.util.function.Function;

public interface GenericHandler {

    <T extends Entity> GenericHandler when(Class<T> entityClass, Condition<T> condition);

    default GenericHandler when(Condition<Entity> condition) {
        return when(Entity.class, condition);
    }

    GenericHandler configure(Function<ExtensionInfo, ExtensionInfo> configurer);

    <T extends Entity> GenericHandler operate(Class<T> entityClass, Operator<T> entityOperator);

    <T extends Entity> GenericHandler operate(Class<T> entityClass,BiFunction<Extension.Event, T, T> entityOperator);
}
