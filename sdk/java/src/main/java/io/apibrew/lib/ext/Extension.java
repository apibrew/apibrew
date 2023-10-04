package io.apibrew.common.ext;

import io.apibrew.common.Entity;
import io.apibrew.common.model.User;

public interface Extension {
    <T extends Entity> Handler<T> handler(Class<T> userClass);
}
