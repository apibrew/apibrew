package io.apibrew.client.ext;

import io.apibrew.client.Entity;
import io.apibrew.client.EntityInfo;
import io.apibrew.client.ExtensionInfo;
import io.apibrew.client.model.Extension;
import io.apibrew.client.model.Record;

import java.io.IOException;
import java.util.function.BiFunction;

public interface ExtensionService extends AutoCloseable {
    <T extends Entity> Handler<T> handler(Class<T> userClass);
    <T extends Entity> Handler<T> handler(EntityInfo<T> entityInfo);

    GenericHandler genericHandler();

    void run() throws IOException;
    void runAsync() throws IOException;

    String registerExtensionWithOperator(ExtensionInfo extensionInfo, BiFunction<Extension.Event, Record, Record> operator);

    void unRegisterOperator(String id);

    void registerPendingItems();
}
