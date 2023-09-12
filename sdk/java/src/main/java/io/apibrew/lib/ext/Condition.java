package io.apibrew.lib.ext;

import io.apibrew.lib.Entity;
import io.apibrew.lib.model.User;

public interface Condition<T extends Entity> {
    static <T extends Entity> Condition<T> before() {
    }

    static <T extends Entity> Condition<T> create() {
    }

    static <T extends Entity> Condition<T> beforeCreate() {
    }

    static <T extends Entity> Condition<T> user(String username) {
    }

    static <T extends Entity> Condition<T> group(String group) {
    }
}
