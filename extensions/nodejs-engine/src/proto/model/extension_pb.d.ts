// package: model
// file: model/extension.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as model_audit_pb from "../model/audit_pb";
import * as model_common_pb from "../model/common_pb";
import * as model_external_pb from "../model/external_pb";
import * as model_annotations_pb from "../model/annotations_pb";
import * as model_event_pb from "../model/event_pb";

export class Extension extends jspb.Message { 
    getId(): string;
    setId(value: string): Extension;
    getName(): string;
    setName(value: string): Extension;
    getDescription(): string;
    setDescription(value: string): Extension;

    hasSelector(): boolean;
    clearSelector(): void;
    getSelector(): model_event_pb.EventSelector | undefined;
    setSelector(value?: model_event_pb.EventSelector): Extension;
    getOrder(): number;
    setOrder(value: number): Extension;
    getFinalizes(): boolean;
    setFinalizes(value: boolean): Extension;
    getSync(): boolean;
    setSync(value: boolean): Extension;
    getResponds(): boolean;
    setResponds(value: boolean): Extension;

    hasCall(): boolean;
    clearCall(): void;
    getCall(): model_external_pb.ExternalCall | undefined;
    setCall(value?: model_external_pb.ExternalCall): Extension;

    hasAuditdata(): boolean;
    clearAuditdata(): void;
    getAuditdata(): model_audit_pb.AuditData | undefined;
    setAuditdata(value?: model_audit_pb.AuditData): Extension;
    getVersion(): number;
    setVersion(value: number): Extension;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Extension.AsObject;
    static toObject(includeInstance: boolean, msg: Extension): Extension.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Extension, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Extension;
    static deserializeBinaryFromReader(message: Extension, reader: jspb.BinaryReader): Extension;
}

export namespace Extension {
    export type AsObject = {
        id: string,
        name: string,
        description: string,
        selector?: model_event_pb.EventSelector.AsObject,
        order: number,
        finalizes: boolean,
        sync: boolean,
        responds: boolean,
        call?: model_external_pb.ExternalCall.AsObject,
        auditdata?: model_audit_pb.AuditData.AsObject,
        version: number,

        annotationsMap: Array<[string, string]>,
    }
}
