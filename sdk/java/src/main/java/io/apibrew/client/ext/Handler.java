package io.apibrew.client.ext;

import io.apibrew.client.Entity;
import io.apibrew.client.ExtensionInfo;
import io.apibrew.client.model.Extension;

import java.util.function.BiFunction;
import java.util.function.Function;

public interface Handler<T extends Entity> {

    Handler<T> when(Condition<T> condition);

    Handler<T> configure(Function<ExtensionInfo, ExtensionInfo> configurer);

    String operate(Operator<T> entityOperator);
    String operate(BiFunction<Extension.Event, T, T> entityOperator);

    void unRegister(String id);
}
