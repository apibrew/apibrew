package io.apibrew.common.ext;

import io.apibrew.common.Entity;

import java.util.function.Consumer;

public interface Action<T extends Entity> {

    static <T extends Entity> Action<T> execute(Consumer<T> consumer) {
    }

    static <T extends Entity> Action<T> reject(String message) {
    }

    static <T extends Entity> Action<T> reject() {
    }
}
