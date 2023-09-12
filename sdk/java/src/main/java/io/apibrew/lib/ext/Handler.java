package io.apibrew.lib.ext;

import io.apibrew.lib.Entity;

public interface Handler<T extends Entity> {

    @SuppressWarnings("unchecked")
    Handler<T> when(Condition<T> condition);
    @SuppressWarnings("unchecked")
    Handler<T> then(Action<T> action);

}
