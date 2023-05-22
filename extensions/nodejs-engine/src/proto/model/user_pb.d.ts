// package: model
// file: model/user.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";
import * as model_audit_pb from "../model/audit_pb";
import * as model_common_pb from "../model/common_pb";
import * as model_security_pb from "../model/security_pb";
import * as model_annotations_pb from "../model/annotations_pb";

export class User extends jspb.Message { 
    getId(): string;
    setId(value: string): User;
    getUsername(): string;
    setUsername(value: string): User;
    getPassword(): string;
    setPassword(value: string): User;

    hasSecuritycontext(): boolean;
    clearSecuritycontext(): void;
    getSecuritycontext(): model_security_pb.SecurityContext | undefined;
    setSecuritycontext(value?: model_security_pb.SecurityContext): User;

    hasDetails(): boolean;
    clearDetails(): void;
    getDetails(): google_protobuf_struct_pb.Struct | undefined;
    setDetails(value?: google_protobuf_struct_pb.Struct): User;
    getSignkey(): string;
    setSignkey(value: string): User;

    hasAuditdata(): boolean;
    clearAuditdata(): void;
    getAuditdata(): model_audit_pb.AuditData | undefined;
    setAuditdata(value?: model_audit_pb.AuditData): User;
    getVersion(): number;
    setVersion(value: number): User;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): User.AsObject;
    static toObject(includeInstance: boolean, msg: User): User.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): User;
    static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
    export type AsObject = {
        id: string,
        username: string,
        password: string,
        securitycontext?: model_security_pb.SecurityContext.AsObject,
        details?: google_protobuf_struct_pb.Struct.AsObject,
        signkey: string,
        auditdata?: model_audit_pb.AuditData.AsObject,
        version: number,
    }
}
