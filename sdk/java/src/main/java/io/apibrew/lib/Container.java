package io.apibrew.lib;

import java.util.List;

public class Container<T extends Entity> {
    private final List<T> records;

    private final int total;

    public Container(List<T> records, int total) {
        this.records = records;
        this.total = total;
    }
}
