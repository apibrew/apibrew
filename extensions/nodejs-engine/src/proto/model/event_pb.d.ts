// package: model
// file: model/event.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as model_resource_pb from "../model/resource_pb";
import * as model_record_pb from "../model/record_pb";
import * as model_query_pb from "../model/query_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class Event extends jspb.Message { 
    getId(): string;
    setId(value: string): Event;
    getAction(): Event.Action;
    setAction(value: Event.Action): Event;
    getActionsummary(): string;
    setActionsummary(value: string): Event;
    getActiondescription(): string;
    setActiondescription(value: string): Event;

    hasResource(): boolean;
    clearResource(): void;
    getResource(): model_resource_pb.Resource | undefined;
    setResource(value?: model_resource_pb.Resource): Event;
    clearRecordsList(): void;
    getRecordsList(): Array<model_record_pb.Record>;
    setRecordsList(value: Array<model_record_pb.Record>): Event;
    addRecords(value?: model_record_pb.Record, index?: number): model_record_pb.Record;
    clearIdsList(): void;
    getIdsList(): Array<string>;
    setIdsList(value: Array<string>): Event;
    addIds(value: string, index?: number): string;

    hasRecordsearchparams(): boolean;
    clearRecordsearchparams(): void;
    getRecordsearchparams(): Event.RecordSearchParams | undefined;
    setRecordsearchparams(value?: Event.RecordSearchParams): Event;
    getFinalizes(): boolean;
    setFinalizes(value: boolean): Event;
    getSync(): boolean;
    setSync(value: boolean): Event;

    hasTime(): boolean;
    clearTime(): void;
    getTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setTime(value?: google_protobuf_timestamp_pb.Timestamp): Event;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Event.AsObject;
    static toObject(includeInstance: boolean, msg: Event): Event.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Event, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Event;
    static deserializeBinaryFromReader(message: Event, reader: jspb.BinaryReader): Event;
}

export namespace Event {
    export type AsObject = {
        id: string,
        action: Event.Action,
        actionsummary: string,
        actiondescription: string,
        resource?: model_resource_pb.Resource.AsObject,
        recordsList: Array<model_record_pb.Record.AsObject>,
        idsList: Array<string>,
        recordsearchparams?: Event.RecordSearchParams.AsObject,
        finalizes: boolean,
        sync: boolean,
        time?: google_protobuf_timestamp_pb.Timestamp.AsObject,

        annotationsMap: Array<[string, string]>,
    }


    export class RecordSearchParams extends jspb.Message { 

        hasQuery(): boolean;
        clearQuery(): void;
        getQuery(): model_query_pb.BooleanExpression | undefined;
        setQuery(value?: model_query_pb.BooleanExpression): RecordSearchParams;
        getLimit(): number;
        setLimit(value: number): RecordSearchParams;
        getOffset(): number;
        setOffset(value: number): RecordSearchParams;
        clearResolvereferencesList(): void;
        getResolvereferencesList(): Array<string>;
        setResolvereferencesList(value: Array<string>): RecordSearchParams;
        addResolvereferences(value: string, index?: number): string;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): RecordSearchParams.AsObject;
        static toObject(includeInstance: boolean, msg: RecordSearchParams): RecordSearchParams.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: RecordSearchParams, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): RecordSearchParams;
        static deserializeBinaryFromReader(message: RecordSearchParams, reader: jspb.BinaryReader): RecordSearchParams;
    }

    export namespace RecordSearchParams {
        export type AsObject = {
            query?: model_query_pb.BooleanExpression.AsObject,
            limit: number,
            offset: number,
            resolvereferencesList: Array<string>,
        }
    }


    export enum Action {
    CREATE = 0,
    UPDATE = 1,
    DELETE = 2,
    GET = 3,
    LIST = 4,
    OPERATE = 5,
    }

}

export class EventSelector extends jspb.Message { 
    clearActionsList(): void;
    getActionsList(): Array<Event.Action>;
    setActionsList(value: Array<Event.Action>): EventSelector;
    addActions(value: Event.Action, index?: number): Event.Action;

    hasRecordselector(): boolean;
    clearRecordselector(): void;
    getRecordselector(): model_query_pb.BooleanExpression | undefined;
    setRecordselector(value?: model_query_pb.BooleanExpression): EventSelector;
    clearNamespacesList(): void;
    getNamespacesList(): Array<string>;
    setNamespacesList(value: Array<string>): EventSelector;
    addNamespaces(value: string, index?: number): string;
    clearResourcesList(): void;
    getResourcesList(): Array<string>;
    setResourcesList(value: Array<string>): EventSelector;
    addResources(value: string, index?: number): string;
    clearIdsList(): void;
    getIdsList(): Array<string>;
    setIdsList(value: Array<string>): EventSelector;
    addIds(value: string, index?: number): string;

    getAnnotationsMap(): jspb.Map<string, string>;
    clearAnnotationsMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): EventSelector.AsObject;
    static toObject(includeInstance: boolean, msg: EventSelector): EventSelector.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: EventSelector, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): EventSelector;
    static deserializeBinaryFromReader(message: EventSelector, reader: jspb.BinaryReader): EventSelector;
}

export namespace EventSelector {
    export type AsObject = {
        actionsList: Array<Event.Action>,
        recordselector?: model_query_pb.BooleanExpression.AsObject,
        namespacesList: Array<string>,
        resourcesList: Array<string>,
        idsList: Array<string>,

        annotationsMap: Array<[string, string]>,
    }
}
