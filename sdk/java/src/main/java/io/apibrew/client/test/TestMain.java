package io.apibrew.client.test;

import io.apibrew.client.ApiException;
import io.apibrew.client.Client;
import io.apibrew.client.nano.NanoService;
import io.apibrew.client.nano.impl.NanoServiceImpl;
import io.apibrew.client.nano.model.Code;
import io.apibrew.client.storage.StorageService;
import io.apibrew.client.storage.impl.StorageServiceImpl;
import io.apibrew.client.storage.model.StorageObject;

import java.io.File;

public class TestMain {

    public static void main(String[] args) {
        Client client = Client.newClient();

//        StorageService storageService = new StorageServiceImpl(client, "http://localhost:8080/local");
//
//        StorageObject object = storageService.repository().create(new StorageObject());
//
//        storageService.uploadFile(object.getId(), new File("/Users/taleh/Projects/apibrew/apibrew/sdk/java/src/main/java/io/apibrew/client/test/TestMain.java"));
//
//        System.out.println(new String(storageService.downloadBytes(object.getId())));


        NanoService nanoService = new NanoServiceImpl(client);

        try {
            nanoService.deploy("test-97.js", Code.Language.JAVASCRIPT, "console.log('Hello World!');", true);
        } catch (ApiException e) {
            e.printStackTrace();
        }
    }
}
