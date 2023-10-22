package io.apibrew.client;

import io.apibrew.client.model.Extension;

import java.util.function.Consumer;

public interface Watcher<T extends Entity> extends AutoCloseable {

    void start();

    void run();
}
