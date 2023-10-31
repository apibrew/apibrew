import {Repository} from "../repository";
import {StorageObject} from "./model/storage-object";

export interface StorageService {
    repository(): Repository<StorageObject>;

    uploadBytes(id: string, data: Buffer, fileName: string, mimeType?: string): Promise<void>;

    uploadFile(id: string, file: string, mimeType?: string): Promise<void>;

    downloadBytes(id: string): Promise<Buffer>;

    downloadFile(id: string, destinationPath: string): Promise<void>;
}
