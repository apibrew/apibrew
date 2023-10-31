import {StorageService} from "../storage-service";
import * as Buffer from "buffer";
import {Repository} from "../../repository";
import {StorageObject, StorageObjectEntityInfo} from "../model/storage-object";
import {Client} from "../../client";
import axios from "axios";
import FormData from 'form-data'
import {ClientImpl} from "../../impl/client-impl";

export class StorageServiceImpl implements StorageService {

    constructor(private client: Client, private baseUrl: string) {
    }

    repository(): Repository<StorageObject> {
        return this.client.repo(StorageObjectEntityInfo);
    }

    async uploadBytes(id: string, data: Buffer, fileName: string): Promise<void> {
        const formData = new FormData();

        formData.append('file', data, fileName)

        const resp = await axios.postForm(this.baseUrl + '/' + id, formData, {
            headers: this.client.headers(),
            validateStatus: (status) => true,
        })

        ClientImpl.ensureResponseSuccess(resp)
    }

    downloadFile(id: string, destinationPath?: string): Promise<void> {
        return Promise.resolve(undefined);
    }

    async downloadBytes(id: string): Promise<Buffer> {
        const resp = await axios.get<Buffer>(this.baseUrl + '/' + id, {
            headers: this.client.headers(),
            validateStatus: (status) => true,
        })

        ClientImpl.ensureResponseSuccess(resp)

        return resp.data
    }

    uploadFile(id: string, file: string, mimeType?: string): Promise<void> {
        return Promise.resolve(undefined);
    }

}

// package io.apibrew.client.storage.impl;
//
// import io.apibrew.client.Client;
// import io.apibrew.client.Repository;
// import io.apibrew.client.impl.ClientImpl;
// import io.apibrew.client.storage.StorageService;
// import io.apibrew.client.storage.model.StorageObject;
// import kong.unirest.ContentType;
// import kong.unirest.HttpResponse;
// import kong.unirest.Unirest;
// import lombok.RequiredArgsConstructor;
//
// import java.io.File;
// import java.io.InputStream;
// import java.util.UUID;
//
// @RequiredArgsConstructor
// public class StorageServiceImpl implements StorageService {
//
//     private final Client client;
//     private final String baseUrl;
//
//     @Override
//     public Repository<StorageObject> repository() {
//         return client.repo(StorageObject.class);
//     }
//
//     @Override
//     public void uploadBytes(UUID id, byte[] data, String fileName) {
//         HttpResponse<?> result = Unirest.post(baseUrl + "/" + id.toString())
//                 .header("Authorization", client.headers().get("Authorization"))
//                 .multiPartContent()
//                 .field("file", data, fileName)
//                 .asEmpty();
//
//         ClientImpl.ensureResponseSuccess(result);
//     }
//
//     @Override
//     public void uploadBytes(UUID id, byte[] data, String fileName, String mimeType) {
//         HttpResponse<?> result = Unirest.post(baseUrl + "/" + id.toString())
//                 .header("Authorization", client.headers().get("Authorization"))
//                 .multiPartContent()
//                 .field("file", data, ContentType.create(mimeType), fileName)
//                 .asEmpty();
//
//         ClientImpl.ensureResponseSuccess(result);
//     }
//
//     @Override
//     public void uploadFile(UUID id, File file) {
//         HttpResponse<?> result = Unirest.post(baseUrl + "/" + id.toString())
//                 .header("Authorization", client.headers().get("Authorization"))
//                 .multiPartContent()
//                 .field("file", file)
//                 .asEmpty();
//
//         ClientImpl.ensureResponseSuccess(result);
//     }
//
//     @Override
//     public void uploadFile(UUID id, File file, String mimeType) {
//         HttpResponse<?> result = Unirest.post(baseUrl + "/" + id.toString())
//                 .header("Authorization", client.headers().get("Authorization"))
//                 .multiPartContent()
//                 .field("file", file, mimeType)
//                 .asEmpty();
//
//         ClientImpl.ensureResponseSuccess(result);
//     }
//
//     @Override
//     public void uploadStream(UUID id, InputStream value, String fileName) {
//         HttpResponse<?> result = Unirest.post(baseUrl + "/" + id.toString())
//                 .header("Authorization", client.headers().get("Authorization"))
//                 .multiPartContent()
//                 .field("file", value, fileName)
//                 .asEmpty();
//
//         ClientImpl.ensureResponseSuccess(result);
//     }
//
//     @Override
//     public void uploadStream(UUID id, InputStream value, String fileName, String mimeType) {
//         HttpResponse<?> result = Unirest.post(baseUrl + "/" + id.toString())
//                 .header("Authorization", client.headers().get("Authorization"))
//                 .multiPartContent()
//                 .field("file", value, ContentType.create(mimeType), fileName)
//                 .asEmpty();
//
//         ClientImpl.ensureResponseSuccess(result);
//     }
//
//     @Override
//     public byte[] downloadBytes(UUID id) {
//         HttpResponse<byte[]> result = Unirest.get(baseUrl + "/" + id.toString())
//                 .headers(client.headers())
//                 .asBytes();
//
//         ClientImpl.ensureResponseSuccess(result);
//
//         return result.getBody();
//     }
//
//     @Override
//     public File downloadFile(UUID id, String destinationPath) {
//         HttpResponse<File> result = Unirest.get(baseUrl + "/" + id.toString())
//                 .headers(client.headers())
//                 .asFile(destinationPath);
//
//         ClientImpl.ensureResponseSuccess(result);
//
//         return result.getBody();
//     }
// }