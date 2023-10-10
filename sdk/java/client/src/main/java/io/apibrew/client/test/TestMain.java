package io.apibrew.client.test;

import io.apibrew.client.Client;
import io.apibrew.client.model.logic.Function;

import java.util.UUID;

public class TestMain {

    public static void main(String[] args) {
        Client client = Client.newClientByServerName("local");

        TestService testService = new TestService(client);

        String resp = testService.test(new Function().withId(UUID.randomUUID()), "abccc");

        System.out.println(resp);
    }
}
