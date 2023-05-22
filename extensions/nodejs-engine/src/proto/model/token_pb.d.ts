// package: model
// file: model/token.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class Token extends jspb.Message { 
    getTerm(): TokenTerm;
    setTerm(value: TokenTerm): Token;
    getContent(): string;
    setContent(value: string): Token;

    hasExpiration(): boolean;
    clearExpiration(): void;
    getExpiration(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setExpiration(value?: google_protobuf_timestamp_pb.Timestamp): Token;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Token.AsObject;
    static toObject(includeInstance: boolean, msg: Token): Token.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Token, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Token;
    static deserializeBinaryFromReader(message: Token, reader: jspb.BinaryReader): Token;
}

export namespace Token {
    export type AsObject = {
        term: TokenTerm,
        content: string,
        expiration?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    }
}

export enum TokenTerm {
    SHORT = 0,
    MIDDLE = 1,
    LONG = 2,
    VERY_LONG = 3,
}
