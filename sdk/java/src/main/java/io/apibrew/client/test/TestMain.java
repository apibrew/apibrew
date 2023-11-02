package io.apibrew.client.test;

import io.apibrew.client.Client;
import io.apibrew.client.nano.NanoService;
import io.apibrew.client.nano.impl.NanoServiceImpl;
import io.apibrew.client.storage.StorageService;
import io.apibrew.client.storage.impl.StorageServiceImpl;
import io.apibrew.client.storage.model.StorageObject;

import java.io.File;
import java.io.FileInputStream;

public class TestMain {

    public static void main(String[] args) {
        Client client = Client.newClient();

        // Firstly we need to prepare nano service
        NanoService nanoService = new NanoServiceImpl(client);

        // Then lets deploy our firs code to nano service

        // There are multiple ways to deploy code to nano service

        // 1. Deploy code from file


    }
}
