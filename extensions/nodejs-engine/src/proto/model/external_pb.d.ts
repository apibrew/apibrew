// package: model
// file: model/external.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as model_annotations_pb from "../model/annotations_pb";

export class FunctionCall extends jspb.Message { 
    getHost(): string;
    setHost(value: string): FunctionCall;
    getFunctionname(): string;
    setFunctionname(value: string): FunctionCall;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FunctionCall.AsObject;
    static toObject(includeInstance: boolean, msg: FunctionCall): FunctionCall.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FunctionCall, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FunctionCall;
    static deserializeBinaryFromReader(message: FunctionCall, reader: jspb.BinaryReader): FunctionCall;
}

export namespace FunctionCall {
    export type AsObject = {
        host: string,
        functionname: string,
    }
}

export class HttpCall extends jspb.Message { 
    getUri(): string;
    setUri(value: string): HttpCall;
    getMethod(): string;
    setMethod(value: string): HttpCall;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): HttpCall.AsObject;
    static toObject(includeInstance: boolean, msg: HttpCall): HttpCall.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: HttpCall, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): HttpCall;
    static deserializeBinaryFromReader(message: HttpCall, reader: jspb.BinaryReader): HttpCall;
}

export namespace HttpCall {
    export type AsObject = {
        uri: string,
        method: string,
    }
}

export class ExternalCall extends jspb.Message { 

    hasFunctioncall(): boolean;
    clearFunctioncall(): void;
    getFunctioncall(): FunctionCall | undefined;
    setFunctioncall(value?: FunctionCall): ExternalCall;

    hasHttpcall(): boolean;
    clearHttpcall(): void;
    getHttpcall(): HttpCall | undefined;
    setHttpcall(value?: HttpCall): ExternalCall;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ExternalCall.AsObject;
    static toObject(includeInstance: boolean, msg: ExternalCall): ExternalCall.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ExternalCall, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ExternalCall;
    static deserializeBinaryFromReader(message: ExternalCall, reader: jspb.BinaryReader): ExternalCall;
}

export namespace ExternalCall {
    export type AsObject = {
        functioncall?: FunctionCall.AsObject,
        httpcall?: HttpCall.AsObject,
    }
}
