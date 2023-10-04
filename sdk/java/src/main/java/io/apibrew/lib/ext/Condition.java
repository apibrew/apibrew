package io.apibrew.common.ext;

import io.apibrew.common.Entity;
import io.apibrew.common.model.User;

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
