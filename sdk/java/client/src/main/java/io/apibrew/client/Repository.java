package io.apibrew.client;

import io.apibrew.client.model.Extension;

public interface Repository<T extends Entity> {

    T create(T record);
    T get(String id);
    T update(T record);
    T delete(String id);
    T apply(T record);
     Container<T> list();

    Container<T> list(Extension.BooleanExpression query);
}
