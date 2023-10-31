package io.apibrew.client;

import io.apibrew.client.model.Extension;

import java.util.List;
import java.util.function.Consumer;

public interface Repository<T extends Entity> {

    T create(T record);
    T get(String id);
    T get(String id, List<String> resolveReferences);
    T update(T record);
    T delete(String id);
    T apply(T record);
     Container<T> list();
     Container<T> list(ListRecordParams params);

    Container<T> list(Extension.BooleanExpression ...query);

     T load(T record);

    Watcher<T> watch(Consumer<Extension.Event> eventConsumer);
}
