package io.apibrew.common;

import java.util.UUID;

public abstract class Entity {
    protected UUID id;

    public UUID getId() {
        return id;
    }

    public void setId(UUID id) {
        this.id = id;
    }
}
