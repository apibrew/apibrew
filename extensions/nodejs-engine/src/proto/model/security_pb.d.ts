// package: model
// file: model/security.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as model_annotations_pb from "../model/annotations_pb";

export class SecurityConstraint extends jspb.Message { 
    getNamespace(): string;
    setNamespace(value: string): SecurityConstraint;
    getResource(): string;
    setResource(value: string): SecurityConstraint;
    getProperty(): string;
    setProperty(value: string): SecurityConstraint;

    hasBefore(): boolean;
    clearBefore(): void;
    getBefore(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setBefore(value?: google_protobuf_timestamp_pb.Timestamp): SecurityConstraint;

    hasAfter(): boolean;
    clearAfter(): void;
    getAfter(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setAfter(value?: google_protobuf_timestamp_pb.Timestamp): SecurityConstraint;
    getPrincipal(): string;
    setPrincipal(value: string): SecurityConstraint;
    clearRecordidsList(): void;
    getRecordidsList(): Array<string>;
    setRecordidsList(value: Array<string>): SecurityConstraint;
    addRecordids(value: string, index?: number): string;
    getOperation(): OperationType;
    setOperation(value: OperationType): SecurityConstraint;
    getPermit(): PermitType;
    setPermit(value: PermitType): SecurityConstraint;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SecurityConstraint.AsObject;
    static toObject(includeInstance: boolean, msg: SecurityConstraint): SecurityConstraint.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SecurityConstraint, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SecurityConstraint;
    static deserializeBinaryFromReader(message: SecurityConstraint, reader: jspb.BinaryReader): SecurityConstraint;
}

export namespace SecurityConstraint {
    export type AsObject = {
        namespace: string,
        resource: string,
        property: string,
        before?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        after?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        principal: string,
        recordidsList: Array<string>,
        operation: OperationType,
        permit: PermitType,
    }
}

export class SecurityContext extends jspb.Message { 
    clearConstraintsList(): void;
    getConstraintsList(): Array<SecurityConstraint>;
    setConstraintsList(value: Array<SecurityConstraint>): SecurityContext;
    addConstraints(value?: SecurityConstraint, index?: number): SecurityConstraint;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SecurityContext.AsObject;
    static toObject(includeInstance: boolean, msg: SecurityContext): SecurityContext.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SecurityContext, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SecurityContext;
    static deserializeBinaryFromReader(message: SecurityContext, reader: jspb.BinaryReader): SecurityContext;
}

export namespace SecurityContext {
    export type AsObject = {
        constraintsList: Array<SecurityConstraint.AsObject>,
    }
}

export enum OperationType {
    OPERATION_TYPE_READ = 0,
    OPERATION_TYPE_CREATE = 1,
    OPERATION_TYPE_UPDATE = 2,
    OPERATION_TYPE_DELETE = 3,
    FULL = 4,
}

export enum PermitType {
    PERMIT_TYPE_ALLOW = 0,
    PERMIT_TYPE_REJECT = 1,
    PERMIT_TYPE_UNKNOWN = 2,
}
