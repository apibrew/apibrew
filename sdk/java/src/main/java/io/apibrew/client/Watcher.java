package io.apibrew.client;

public interface Watcher<T extends Entity> extends AutoCloseable {

    void start();

    void run();
}
