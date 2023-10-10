package io.apibrew.client.test;

import io.apibrew.client.Client;
import io.apibrew.client.Service;
import io.apibrew.client.model.logic.Function;

public class TestService implements Service {

    private final Client client;

    public TestService(Client client) {
        this.client = client;
    }
    public String test(Function function, String input) {
        return client.executeRecordAction(String.class, Function.NAMESPACE, Function.RESOURCE, function.getId().toString(), "test", input);
    }
}
