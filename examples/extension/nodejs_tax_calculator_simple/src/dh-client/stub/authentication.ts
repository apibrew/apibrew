/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 0.0.0
 * source: stub/authentication.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./../model/error";
import * as dependency_2 from "./../model/token";
import * as dependency_3 from "./../google/api/annotations";
import * as dependency_4 from "./../openapiv3/annotations";
import * as pb_1 from "google-protobuf";
import * as grpc_1 from "@grpc/grpc-js";
export namespace stub {
    export class AuthenticationRequest extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            username?: string;
            password?: string;
            term?: dependency_2.model.TokenTerm;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("username" in data && data.username != undefined) {
                    this.username = data.username;
                }
                if ("password" in data && data.password != undefined) {
                    this.password = data.password;
                }
                if ("term" in data && data.term != undefined) {
                    this.term = data.term;
                }
            }
        }
        get username() {
            return pb_1.Message.getFieldWithDefault(this, 1, "") as string;
        }
        set username(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        get password() {
            return pb_1.Message.getFieldWithDefault(this, 2, "") as string;
        }
        set password(value: string) {
            pb_1.Message.setField(this, 2, value);
        }
        get term() {
            return pb_1.Message.getFieldWithDefault(this, 3, dependency_2.model.TokenTerm.SHORT) as dependency_2.model.TokenTerm;
        }
        set term(value: dependency_2.model.TokenTerm) {
            pb_1.Message.setField(this, 3, value);
        }
        static fromObject(data: {
            username?: string;
            password?: string;
            term?: dependency_2.model.TokenTerm;
        }): AuthenticationRequest {
            const message = new AuthenticationRequest({});
            if (data.username != null) {
                message.username = data.username;
            }
            if (data.password != null) {
                message.password = data.password;
            }
            if (data.term != null) {
                message.term = data.term;
            }
            return message;
        }
        toObject() {
            const data: {
                username?: string;
                password?: string;
                term?: dependency_2.model.TokenTerm;
            } = {};
            if (this.username != null) {
                data.username = this.username;
            }
            if (this.password != null) {
                data.password = this.password;
            }
            if (this.term != null) {
                data.term = this.term;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.username.length)
                writer.writeString(1, this.username);
            if (this.password.length)
                writer.writeString(2, this.password);
            if (this.term != dependency_2.model.TokenTerm.SHORT)
                writer.writeEnum(3, this.term);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): AuthenticationRequest {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new AuthenticationRequest();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.username = reader.readString();
                        break;
                    case 2:
                        message.password = reader.readString();
                        break;
                    case 3:
                        message.term = reader.readEnum();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): AuthenticationRequest {
            return AuthenticationRequest.deserialize(bytes);
        }
    }
    export class AuthenticationResponse extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            token?: dependency_2.model.Token;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("token" in data && data.token != undefined) {
                    this.token = data.token;
                }
            }
        }
        get token() {
            return pb_1.Message.getWrapperField(this, dependency_2.model.Token, 1) as dependency_2.model.Token;
        }
        set token(value: dependency_2.model.Token) {
            pb_1.Message.setWrapperField(this, 1, value);
        }
        get has_token() {
            return pb_1.Message.getField(this, 1) != null;
        }
        static fromObject(data: {
            token?: ReturnType<typeof dependency_2.model.Token.prototype.toObject>;
        }): AuthenticationResponse {
            const message = new AuthenticationResponse({});
            if (data.token != null) {
                message.token = dependency_2.model.Token.fromObject(data.token);
            }
            return message;
        }
        toObject() {
            const data: {
                token?: ReturnType<typeof dependency_2.model.Token.prototype.toObject>;
            } = {};
            if (this.token != null) {
                data.token = this.token.toObject();
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.has_token)
                writer.writeMessage(1, this.token, () => this.token.serialize(writer));
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): AuthenticationResponse {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new AuthenticationResponse();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        reader.readMessage(message.token, () => message.token = dependency_2.model.Token.deserialize(reader));
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): AuthenticationResponse {
            return AuthenticationResponse.deserialize(bytes);
        }
    }
    export class RenewTokenRequest extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            token?: string;
            term?: dependency_2.model.TokenTerm;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("token" in data && data.token != undefined) {
                    this.token = data.token;
                }
                if ("term" in data && data.term != undefined) {
                    this.term = data.term;
                }
            }
        }
        get token() {
            return pb_1.Message.getFieldWithDefault(this, 1, "") as string;
        }
        set token(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        get term() {
            return pb_1.Message.getFieldWithDefault(this, 2, dependency_2.model.TokenTerm.SHORT) as dependency_2.model.TokenTerm;
        }
        set term(value: dependency_2.model.TokenTerm) {
            pb_1.Message.setField(this, 2, value);
        }
        static fromObject(data: {
            token?: string;
            term?: dependency_2.model.TokenTerm;
        }): RenewTokenRequest {
            const message = new RenewTokenRequest({});
            if (data.token != null) {
                message.token = data.token;
            }
            if (data.term != null) {
                message.term = data.term;
            }
            return message;
        }
        toObject() {
            const data: {
                token?: string;
                term?: dependency_2.model.TokenTerm;
            } = {};
            if (this.token != null) {
                data.token = this.token;
            }
            if (this.term != null) {
                data.term = this.term;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.token.length)
                writer.writeString(1, this.token);
            if (this.term != dependency_2.model.TokenTerm.SHORT)
                writer.writeEnum(2, this.term);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): RenewTokenRequest {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new RenewTokenRequest();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.token = reader.readString();
                        break;
                    case 2:
                        message.term = reader.readEnum();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): RenewTokenRequest {
            return RenewTokenRequest.deserialize(bytes);
        }
    }
    export class RenewTokenResponse extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            token?: dependency_2.model.Token;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("token" in data && data.token != undefined) {
                    this.token = data.token;
                }
            }
        }
        get token() {
            return pb_1.Message.getWrapperField(this, dependency_2.model.Token, 1) as dependency_2.model.Token;
        }
        set token(value: dependency_2.model.Token) {
            pb_1.Message.setWrapperField(this, 1, value);
        }
        get has_token() {
            return pb_1.Message.getField(this, 1) != null;
        }
        static fromObject(data: {
            token?: ReturnType<typeof dependency_2.model.Token.prototype.toObject>;
        }): RenewTokenResponse {
            const message = new RenewTokenResponse({});
            if (data.token != null) {
                message.token = dependency_2.model.Token.fromObject(data.token);
            }
            return message;
        }
        toObject() {
            const data: {
                token?: ReturnType<typeof dependency_2.model.Token.prototype.toObject>;
            } = {};
            if (this.token != null) {
                data.token = this.token.toObject();
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.has_token)
                writer.writeMessage(1, this.token, () => this.token.serialize(writer));
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): RenewTokenResponse {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new RenewTokenResponse();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        reader.readMessage(message.token, () => message.token = dependency_2.model.Token.deserialize(reader));
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): RenewTokenResponse {
            return RenewTokenResponse.deserialize(bytes);
        }
    }
    interface GrpcUnaryServiceInterface<P, R> {
        (message: P, metadata: grpc_1.Metadata, options: grpc_1.CallOptions, callback: grpc_1.requestCallback<R>): grpc_1.ClientUnaryCall;
        (message: P, metadata: grpc_1.Metadata, callback: grpc_1.requestCallback<R>): grpc_1.ClientUnaryCall;
        (message: P, options: grpc_1.CallOptions, callback: grpc_1.requestCallback<R>): grpc_1.ClientUnaryCall;
        (message: P, callback: grpc_1.requestCallback<R>): grpc_1.ClientUnaryCall;
    }
    interface GrpcStreamServiceInterface<P, R> {
        (message: P, metadata: grpc_1.Metadata, options?: grpc_1.CallOptions): grpc_1.ClientReadableStream<R>;
        (message: P, options?: grpc_1.CallOptions): grpc_1.ClientReadableStream<R>;
    }
    interface GrpWritableServiceInterface<P, R> {
        (metadata: grpc_1.Metadata, options: grpc_1.CallOptions, callback: grpc_1.requestCallback<R>): grpc_1.ClientWritableStream<P>;
        (metadata: grpc_1.Metadata, callback: grpc_1.requestCallback<R>): grpc_1.ClientWritableStream<P>;
        (options: grpc_1.CallOptions, callback: grpc_1.requestCallback<R>): grpc_1.ClientWritableStream<P>;
        (callback: grpc_1.requestCallback<R>): grpc_1.ClientWritableStream<P>;
    }
    interface GrpcChunkServiceInterface<P, R> {
        (metadata: grpc_1.Metadata, options?: grpc_1.CallOptions): grpc_1.ClientDuplexStream<P, R>;
        (options?: grpc_1.CallOptions): grpc_1.ClientDuplexStream<P, R>;
    }
    interface GrpcPromiseServiceInterface<P, R> {
        (message: P, metadata: grpc_1.Metadata, options?: grpc_1.CallOptions): Promise<R>;
        (message: P, options?: grpc_1.CallOptions): Promise<R>;
    }
    export abstract class UnimplementedAuthenticationService {
        static definition = {
            Authenticate: {
                path: "/stub.Authentication/Authenticate",
                requestStream: false,
                responseStream: false,
                requestSerialize: (message: AuthenticationRequest) => Buffer.from(message.serialize()),
                requestDeserialize: (bytes: Buffer) => AuthenticationRequest.deserialize(new Uint8Array(bytes)),
                responseSerialize: (message: AuthenticationResponse) => Buffer.from(message.serialize()),
                responseDeserialize: (bytes: Buffer) => AuthenticationResponse.deserialize(new Uint8Array(bytes))
            },
            RenewToken: {
                path: "/stub.Authentication/RenewToken",
                requestStream: false,
                responseStream: false,
                requestSerialize: (message: RenewTokenRequest) => Buffer.from(message.serialize()),
                requestDeserialize: (bytes: Buffer) => RenewTokenRequest.deserialize(new Uint8Array(bytes)),
                responseSerialize: (message: RenewTokenResponse) => Buffer.from(message.serialize()),
                responseDeserialize: (bytes: Buffer) => RenewTokenResponse.deserialize(new Uint8Array(bytes))
            }
        };
        [method: string]: grpc_1.UntypedHandleCall;
        abstract Authenticate(call: grpc_1.ServerUnaryCall<AuthenticationRequest, AuthenticationResponse>, callback: grpc_1.sendUnaryData<AuthenticationResponse>): void;
        abstract RenewToken(call: grpc_1.ServerUnaryCall<RenewTokenRequest, RenewTokenResponse>, callback: grpc_1.sendUnaryData<RenewTokenResponse>): void;
    }
    export class AuthenticationClient extends grpc_1.makeGenericClientConstructor(UnimplementedAuthenticationService.definition, "Authentication", {}) {
        constructor(address: string, credentials: grpc_1.ChannelCredentials, options?: Partial<grpc_1.ChannelOptions>) {
            super(address, credentials, options);
        }
        Authenticate: GrpcUnaryServiceInterface<AuthenticationRequest, AuthenticationResponse> = (message: AuthenticationRequest, metadata: grpc_1.Metadata | grpc_1.CallOptions | grpc_1.requestCallback<AuthenticationResponse>, options?: grpc_1.CallOptions | grpc_1.requestCallback<AuthenticationResponse>, callback?: grpc_1.requestCallback<AuthenticationResponse>): grpc_1.ClientUnaryCall => {
            return super.Authenticate(message, metadata, options, callback);
        };
        RenewToken: GrpcUnaryServiceInterface<RenewTokenRequest, RenewTokenResponse> = (message: RenewTokenRequest, metadata: grpc_1.Metadata | grpc_1.CallOptions | grpc_1.requestCallback<RenewTokenResponse>, options?: grpc_1.CallOptions | grpc_1.requestCallback<RenewTokenResponse>, callback?: grpc_1.requestCallback<RenewTokenResponse>): grpc_1.ClientUnaryCall => {
            return super.RenewToken(message, metadata, options, callback);
        };
    }
}
