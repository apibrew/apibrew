package io.apibrew.client.test;

import io.apibrew.client.Client;
import io.apibrew.client.nano.NanoService;
import io.apibrew.client.nano.impl.NanoServiceImpl;

public class TestMain {

    public static void main(String[] args) {
        Client client = Client.newClient();

        // Firstly we need to prepare nano service
        NanoService nanoService = new NanoServiceImpl(client);

        // Then lets deploy our firs code to nano service

        // There are multiple ways to deploy code to nano service

        // 1. Deploy code from file


        Student student1 = new Student();

        student1.setName("John");

        student1.toString();


    }
}
