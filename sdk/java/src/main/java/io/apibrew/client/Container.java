package io.apibrew.client;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

public class Container<T extends Entity> {
    @JsonProperty
    private List<T> content;

    @JsonProperty
    private int total;

    public int getTotal() {
        return total;
    }

    public List<T> getContent() {
        return content;
    }

    public void setContent(List<T> content) {
        this.content = content;
    }

    public void setTotal(int total) {
        this.total = total;
    }
}
