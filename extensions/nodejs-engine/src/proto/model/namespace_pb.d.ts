// package: model
// file: model/namespace.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as model_audit_pb from "../model/audit_pb";
import * as model_common_pb from "../model/common_pb";
import * as model_security_pb from "../model/security_pb";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";

export class Namespace extends jspb.Message { 
    getId(): string;
    setId(value: string): Namespace;
    getName(): string;
    setName(value: string): Namespace;
    getDescription(): string;
    setDescription(value: string): Namespace;

    hasDetails(): boolean;
    clearDetails(): void;
    getDetails(): google_protobuf_struct_pb.Struct | undefined;
    setDetails(value?: google_protobuf_struct_pb.Struct): Namespace;

    hasSecuritycontext(): boolean;
    clearSecuritycontext(): void;
    getSecuritycontext(): model_security_pb.SecurityContext | undefined;
    setSecuritycontext(value?: model_security_pb.SecurityContext): Namespace;

    hasAuditdata(): boolean;
    clearAuditdata(): void;
    getAuditdata(): model_audit_pb.AuditData | undefined;
    setAuditdata(value?: model_audit_pb.AuditData): Namespace;
    getVersion(): number;
    setVersion(value: number): Namespace;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Namespace.AsObject;
    static toObject(includeInstance: boolean, msg: Namespace): Namespace.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Namespace, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Namespace;
    static deserializeBinaryFromReader(message: Namespace, reader: jspb.BinaryReader): Namespace;
}

export namespace Namespace {
    export type AsObject = {
        id: string,
        name: string,
        description: string,
        details?: google_protobuf_struct_pb.Struct.AsObject,
        securitycontext?: model_security_pb.SecurityContext.AsObject,
        auditdata?: model_audit_pb.AuditData.AsObject,
        version: number,
    }
}
