package io.apibrew.lib.ext;

import io.apibrew.lib.Entity;
import io.apibrew.lib.model.User;

public interface Extension {
    <T extends Entity> Handler<T> handler(Class<T> userClass);
}
