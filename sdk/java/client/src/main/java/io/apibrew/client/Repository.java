package io.apibrew.client;

import io.apibrew.client.model.Extension;

import java.util.List;
import java.util.function.Consumer;

public interface Repository<T extends Entity> {

    T create(T record);
    T get(String id);
    T update(T record);
    T delete(String id);
    T apply(T record);
     Container<T> list();

     T load(T record);

    Container<T> list(Extension.BooleanExpression query);

    Watcher<T> watch(Consumer<Extension.Event> eventConsumer);
}
