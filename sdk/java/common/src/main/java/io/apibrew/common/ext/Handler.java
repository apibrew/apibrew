package io.apibrew.common.ext;

import io.apibrew.client.Entity;
import io.apibrew.common.ext.Condition;
import io.apibrew.common.ext.Operator;
import io.apibrew.common.ExtensionInfo;
import io.apibrew.client.model.Extension;

import java.util.function.BiFunction;
import java.util.function.Function;

public interface Handler<T extends Entity> {

    Handler<T> when(Condition<T> condition);

    Handler<T> configure(Function<ExtensionInfo, ExtensionInfo> configurer);

    Handler<T> operate(Operator<T> entityOperator);
    Handler<T> operate(BiFunction<Extension.Event, T, T> entityOperator);
}
