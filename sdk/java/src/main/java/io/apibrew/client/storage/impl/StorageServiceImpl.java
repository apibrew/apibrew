package io.apibrew.client.storage.impl;

import io.apibrew.client.Client;
import io.apibrew.client.Repository;
import io.apibrew.client.impl.ClientImpl;
import io.apibrew.client.storage.StorageService;
import io.apibrew.client.storage.model.StorageObject;
import kong.unirest.ContentType;
import kong.unirest.HttpResponse;
import kong.unirest.Unirest;
import lombok.RequiredArgsConstructor;

import java.io.File;
import java.io.InputStream;
import java.util.UUID;

@RequiredArgsConstructor
public class StorageServiceImpl implements StorageService {

    private final Client client;
    private final String baseUrl;

    @Override
    public Repository<StorageObject> repository() {
        return client.repo(StorageObject.class);
    }

    @Override
    public void uploadBytes(UUID id, byte[] data, String fileName) {
        HttpResponse<?> result = Unirest.post(baseUrl + "/" + id.toString())
                .header("Authorization", client.headers().get("Authorization"))
                .multiPartContent()
                .field("file", data, fileName)
                .asEmpty();

        ClientImpl.ensureResponseSuccess(result);
    }

    @Override
    public void uploadBytes(UUID id, byte[] data, String fileName, String mimeType) {
        HttpResponse<?> result = Unirest.post(baseUrl + "/" + id.toString())
                .header("Authorization", client.headers().get("Authorization"))
                .multiPartContent()
                .field("file", data, ContentType.create(mimeType), fileName)
                .asEmpty();

        ClientImpl.ensureResponseSuccess(result);
    }

    @Override
    public void uploadFile(UUID id, File file) {
        HttpResponse<?> result = Unirest.post(baseUrl + "/" + id.toString())
                .header("Authorization", client.headers().get("Authorization"))
                .multiPartContent()
                .field("file", file)
                .asEmpty();

        ClientImpl.ensureResponseSuccess(result);
    }

    @Override
    public void uploadFile(UUID id, File file, String mimeType) {
        HttpResponse<?> result = Unirest.post(baseUrl + "/" + id.toString())
                .header("Authorization", client.headers().get("Authorization"))
                .multiPartContent()
                .field("file", file, mimeType)
                .asEmpty();

        ClientImpl.ensureResponseSuccess(result);
    }

    @Override
    public void uploadStream(UUID id, InputStream value, String fileName) {
        HttpResponse<?> result = Unirest.post(baseUrl + "/" + id.toString())
                .header("Authorization", client.headers().get("Authorization"))
                .multiPartContent()
                .field("file", value, fileName)
                .asEmpty();

        ClientImpl.ensureResponseSuccess(result);
    }

    @Override
    public void uploadStream(UUID id, InputStream value, String fileName, String mimeType) {
        HttpResponse<?> result = Unirest.post(baseUrl + "/" + id.toString())
                .header("Authorization", client.headers().get("Authorization"))
                .multiPartContent()
                .field("file", value, ContentType.create(mimeType), fileName)
                .asEmpty();

        ClientImpl.ensureResponseSuccess(result);
    }

    @Override
    public byte[] downloadBytes(UUID id) {
        HttpResponse<byte[]> result = Unirest.get(baseUrl + "/" + id.toString())
                .headers(client.headers())
                .asBytes();

        ClientImpl.ensureResponseSuccess(result);

        return result.getBody();
    }

    @Override
    public File downloadFile(UUID id, String destinationPath) {
        HttpResponse<File> result = Unirest.get(baseUrl + "/" + id.toString())
                .headers(client.headers())
                .asFile(destinationPath);

        ClientImpl.ensureResponseSuccess(result);

        return result.getBody();
    }
}
