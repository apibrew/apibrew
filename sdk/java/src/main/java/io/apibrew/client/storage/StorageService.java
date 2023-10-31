package io.apibrew.client.storage;

import io.apibrew.client.Repository;
import io.apibrew.client.storage.model.StorageObject;

import java.io.File;
import java.io.InputStream;
import java.io.OutputStream;
import java.util.UUID;

public interface StorageService {
    public Repository<StorageObject> repository();

    public void uploadBytes(UUID id, byte[] data, String fileName);

    public void uploadBytes(UUID id, byte[] data, String fileName, String mimeType);

    public void uploadFile(UUID id, File file);

    public void uploadFile(UUID id, File file, String mimeType);

    public void uploadStream(UUID id, InputStream value, String fileName);

    public void uploadStream(UUID id, InputStream value, String fileName, String mimeType);

    public byte[] downloadBytes(UUID id);
    public File downloadFile(UUID id, String destinationPath);
}
