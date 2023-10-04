package io.apibrew.common.ext;

import io.apibrew.common.Entity;

public interface Handler<T extends Entity> {

    @SuppressWarnings("unchecked")
    Handler<T> when(Condition<T> condition);
    @SuppressWarnings("unchecked")
    Handler<T> then(Action<T> action);

}
