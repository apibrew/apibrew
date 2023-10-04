package io.apibrew.common.ext;

import io.apibrew.client.Entity;
import io.apibrew.common.ExtensionInfo;
import io.apibrew.client.model.Extension;
import io.apibrew.client.model.Record;

import java.io.IOException;
import java.util.function.BiFunction;

public interface ExtensionService {
    <T extends Entity> Handler<T> handler(Class<T> userClass);

    GenericHandler genericHandler();

    void run() throws IOException;

    void registerExtensionWithOperator(ExtensionInfo extensionInfo, BiFunction<Extension.Event, Record, Record> operator);
}
