// package: model
// file: model/error.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";

export class ErrorField extends jspb.Message { 
    getRecordid(): string;
    setRecordid(value: string): ErrorField;
    getProperty(): string;
    setProperty(value: string): ErrorField;
    getMessage(): string;
    setMessage(value: string): ErrorField;

    hasValue(): boolean;
    clearValue(): void;
    getValue(): google_protobuf_struct_pb.Value | undefined;
    setValue(value?: google_protobuf_struct_pb.Value): ErrorField;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ErrorField.AsObject;
    static toObject(includeInstance: boolean, msg: ErrorField): ErrorField.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ErrorField, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ErrorField;
    static deserializeBinaryFromReader(message: ErrorField, reader: jspb.BinaryReader): ErrorField;
}

export namespace ErrorField {
    export type AsObject = {
        recordid: string,
        property: string,
        message: string,
        value?: google_protobuf_struct_pb.Value.AsObject,
    }
}

export class Error extends jspb.Message { 
    getCode(): ErrorCode;
    setCode(value: ErrorCode): Error;
    getMessage(): string;
    setMessage(value: string): Error;
    clearFieldsList(): void;
    getFieldsList(): Array<ErrorField>;
    setFieldsList(value: Array<ErrorField>): Error;
    addFields(value?: ErrorField, index?: number): ErrorField;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Error.AsObject;
    static toObject(includeInstance: boolean, msg: Error): Error.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Error, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Error;
    static deserializeBinaryFromReader(message: Error, reader: jspb.BinaryReader): Error;
}

export namespace Error {
    export type AsObject = {
        code: ErrorCode,
        message: string,
        fieldsList: Array<ErrorField.AsObject>,
    }
}

export enum ErrorCode {
    UNKNOWN_ERROR = 0,
    RECORD_NOT_FOUND = 1,
    UNABLE_TO_LOCATE_PRIMARY_KEY = 2,
    INTERNAL_ERROR = 3,
    PROPERTY_NOT_FOUND = 4,
    RECORD_VALIDATION_ERROR = 5,
    RESOURCE_VALIDATION_ERROR = 13,
    AUTHENTICATION_FAILED = 6,
    ALREADY_EXISTS = 7,
    ACCESS_DENIED = 8,
    BACKEND_ERROR = 9,
    UNIQUE_VIOLATION = 10,
    REFERENCE_VIOLATION = 11,
    RESOURCE_NOT_FOUND = 12,
    UNSUPPORTED_OPERATION = 14,
    EXTERNAL_BACKEND_COMMUNICATION_ERROR = 15,
    EXTERNAL_BACKEND_ERROR = 16,
}
