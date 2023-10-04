package io.apibrew.common;

import io.apibrew.common.model.Record;

public interface Repository<T extends Entity> {

    T Create(T record);
    T Get(String id);
    T Update(Record record);
    T Delete(String id);
    T Apply(Record record);
     Container<T> List(String resource, String query, int page, int pageSize);

}
