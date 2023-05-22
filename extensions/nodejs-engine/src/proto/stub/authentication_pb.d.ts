// package: stub
// file: stub/authentication.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as model_error_pb from "../model/error_pb";
import * as model_token_pb from "../model/token_pb";


export class AuthenticationRequest extends jspb.Message { 
    getUsername(): string;
    setUsername(value: string): AuthenticationRequest;
    getPassword(): string;
    setPassword(value: string): AuthenticationRequest;
    getTerm(): model_token_pb.TokenTerm;
    setTerm(value: model_token_pb.TokenTerm): AuthenticationRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AuthenticationRequest.AsObject;
    static toObject(includeInstance: boolean, msg: AuthenticationRequest): AuthenticationRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AuthenticationRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AuthenticationRequest;
    static deserializeBinaryFromReader(message: AuthenticationRequest, reader: jspb.BinaryReader): AuthenticationRequest;
}

export namespace AuthenticationRequest {
    export type AsObject = {
        username: string,
        password: string,
        term: model_token_pb.TokenTerm,
    }
}

export class AuthenticationResponse extends jspb.Message { 

    hasToken(): boolean;
    clearToken(): void;
    getToken(): model_token_pb.Token | undefined;
    setToken(value?: model_token_pb.Token): AuthenticationResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AuthenticationResponse.AsObject;
    static toObject(includeInstance: boolean, msg: AuthenticationResponse): AuthenticationResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AuthenticationResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AuthenticationResponse;
    static deserializeBinaryFromReader(message: AuthenticationResponse, reader: jspb.BinaryReader): AuthenticationResponse;
}

export namespace AuthenticationResponse {
    export type AsObject = {
        token?: model_token_pb.Token.AsObject,
    }
}

export class RenewTokenRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): RenewTokenRequest;
    getTerm(): model_token_pb.TokenTerm;
    setTerm(value: model_token_pb.TokenTerm): RenewTokenRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RenewTokenRequest.AsObject;
    static toObject(includeInstance: boolean, msg: RenewTokenRequest): RenewTokenRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RenewTokenRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RenewTokenRequest;
    static deserializeBinaryFromReader(message: RenewTokenRequest, reader: jspb.BinaryReader): RenewTokenRequest;
}

export namespace RenewTokenRequest {
    export type AsObject = {
        token: string,
        term: model_token_pb.TokenTerm,
    }
}

export class RenewTokenResponse extends jspb.Message { 

    hasToken(): boolean;
    clearToken(): void;
    getToken(): model_token_pb.Token | undefined;
    setToken(value?: model_token_pb.Token): RenewTokenResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RenewTokenResponse.AsObject;
    static toObject(includeInstance: boolean, msg: RenewTokenResponse): RenewTokenResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RenewTokenResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RenewTokenResponse;
    static deserializeBinaryFromReader(message: RenewTokenResponse, reader: jspb.BinaryReader): RenewTokenResponse;
}

export namespace RenewTokenResponse {
    export type AsObject = {
        token?: model_token_pb.Token.AsObject,
    }
}
