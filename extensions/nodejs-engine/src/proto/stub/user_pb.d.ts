// package: stub
// file: stub/user.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as model_error_pb from "../model/error_pb";
import * as model_user_pb from "../model/user_pb";
import * as model_query_pb from "../model/query_pb";


export class CreateUserRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): CreateUserRequest;

    hasUser(): boolean;
    clearUser(): void;
    getUser(): model_user_pb.User | undefined;
    setUser(value?: model_user_pb.User): CreateUserRequest;
    clearUsersList(): void;
    getUsersList(): Array<model_user_pb.User>;
    setUsersList(value: Array<model_user_pb.User>): CreateUserRequest;
    addUsers(value?: model_user_pb.User, index?: number): model_user_pb.User;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateUserRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateUserRequest): CreateUserRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateUserRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateUserRequest;
    static deserializeBinaryFromReader(message: CreateUserRequest, reader: jspb.BinaryReader): CreateUserRequest;
}

export namespace CreateUserRequest {
    export type AsObject = {
        token: string,
        user?: model_user_pb.User.AsObject,
        usersList: Array<model_user_pb.User.AsObject>,
    }
}

export class CreateUserResponse extends jspb.Message { 

    hasUser(): boolean;
    clearUser(): void;
    getUser(): model_user_pb.User | undefined;
    setUser(value?: model_user_pb.User): CreateUserResponse;
    clearUsersList(): void;
    getUsersList(): Array<model_user_pb.User>;
    setUsersList(value: Array<model_user_pb.User>): CreateUserResponse;
    addUsers(value?: model_user_pb.User, index?: number): model_user_pb.User;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateUserResponse.AsObject;
    static toObject(includeInstance: boolean, msg: CreateUserResponse): CreateUserResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateUserResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateUserResponse;
    static deserializeBinaryFromReader(message: CreateUserResponse, reader: jspb.BinaryReader): CreateUserResponse;
}

export namespace CreateUserResponse {
    export type AsObject = {
        user?: model_user_pb.User.AsObject,
        usersList: Array<model_user_pb.User.AsObject>,
    }
}

export class UpdateUserRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): UpdateUserRequest;

    hasUser(): boolean;
    clearUser(): void;
    getUser(): model_user_pb.User | undefined;
    setUser(value?: model_user_pb.User): UpdateUserRequest;
    clearUsersList(): void;
    getUsersList(): Array<model_user_pb.User>;
    setUsersList(value: Array<model_user_pb.User>): UpdateUserRequest;
    addUsers(value?: model_user_pb.User, index?: number): model_user_pb.User;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateUserRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateUserRequest): UpdateUserRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateUserRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateUserRequest;
    static deserializeBinaryFromReader(message: UpdateUserRequest, reader: jspb.BinaryReader): UpdateUserRequest;
}

export namespace UpdateUserRequest {
    export type AsObject = {
        token: string,
        user?: model_user_pb.User.AsObject,
        usersList: Array<model_user_pb.User.AsObject>,
    }
}

export class UpdateUserResponse extends jspb.Message { 

    hasUser(): boolean;
    clearUser(): void;
    getUser(): model_user_pb.User | undefined;
    setUser(value?: model_user_pb.User): UpdateUserResponse;
    clearUsersList(): void;
    getUsersList(): Array<model_user_pb.User>;
    setUsersList(value: Array<model_user_pb.User>): UpdateUserResponse;
    addUsers(value?: model_user_pb.User, index?: number): model_user_pb.User;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateUserResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateUserResponse): UpdateUserResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateUserResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateUserResponse;
    static deserializeBinaryFromReader(message: UpdateUserResponse, reader: jspb.BinaryReader): UpdateUserResponse;
}

export namespace UpdateUserResponse {
    export type AsObject = {
        user?: model_user_pb.User.AsObject,
        usersList: Array<model_user_pb.User.AsObject>,
    }
}

export class DeleteUserRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): DeleteUserRequest;
    getId(): string;
    setId(value: string): DeleteUserRequest;
    clearIdsList(): void;
    getIdsList(): Array<string>;
    setIdsList(value: Array<string>): DeleteUserRequest;
    addIds(value: string, index?: number): string;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteUserRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteUserRequest): DeleteUserRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteUserRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteUserRequest;
    static deserializeBinaryFromReader(message: DeleteUserRequest, reader: jspb.BinaryReader): DeleteUserRequest;
}

export namespace DeleteUserRequest {
    export type AsObject = {
        token: string,
        id: string,
        idsList: Array<string>,
    }
}

export class DeleteUserResponse extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteUserResponse.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteUserResponse): DeleteUserResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteUserResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteUserResponse;
    static deserializeBinaryFromReader(message: DeleteUserResponse, reader: jspb.BinaryReader): DeleteUserResponse;
}

export namespace DeleteUserResponse {
    export type AsObject = {
    }
}

export class ListUserRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): ListUserRequest;
    getLimit(): number;
    setLimit(value: number): ListUserRequest;
    getOffset(): number;
    setOffset(value: number): ListUserRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListUserRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListUserRequest): ListUserRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListUserRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListUserRequest;
    static deserializeBinaryFromReader(message: ListUserRequest, reader: jspb.BinaryReader): ListUserRequest;
}

export namespace ListUserRequest {
    export type AsObject = {
        token: string,
        limit: number,
        offset: number,
    }
}

export class ListUserResponse extends jspb.Message { 
    clearContentList(): void;
    getContentList(): Array<model_user_pb.User>;
    setContentList(value: Array<model_user_pb.User>): ListUserResponse;
    addContent(value?: model_user_pb.User, index?: number): model_user_pb.User;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListUserResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ListUserResponse): ListUserResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListUserResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListUserResponse;
    static deserializeBinaryFromReader(message: ListUserResponse, reader: jspb.BinaryReader): ListUserResponse;
}

export namespace ListUserResponse {
    export type AsObject = {
        contentList: Array<model_user_pb.User.AsObject>,
    }
}

export class GetUserRequest extends jspb.Message { 
    getToken(): string;
    setToken(value: string): GetUserRequest;
    getId(): string;
    setId(value: string): GetUserRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetUserRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetUserRequest): GetUserRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetUserRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetUserRequest;
    static deserializeBinaryFromReader(message: GetUserRequest, reader: jspb.BinaryReader): GetUserRequest;
}

export namespace GetUserRequest {
    export type AsObject = {
        token: string,
        id: string,
    }
}

export class GetUserResponse extends jspb.Message { 

    hasUser(): boolean;
    clearUser(): void;
    getUser(): model_user_pb.User | undefined;
    setUser(value?: model_user_pb.User): GetUserResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetUserResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetUserResponse): GetUserResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetUserResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetUserResponse;
    static deserializeBinaryFromReader(message: GetUserResponse, reader: jspb.BinaryReader): GetUserResponse;
}

export namespace GetUserResponse {
    export type AsObject = {
        user?: model_user_pb.User.AsObject,
    }
}
