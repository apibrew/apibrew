/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 0.0.0
 * source: model/watch.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./../google/protobuf/struct";
import * as dependency_2 from "./../google/protobuf/timestamp";
import * as pb_1 from "google-protobuf";
export enum EventType {
    CREATE = 0,
    UPDATE = 1,
    DELETE = 2,
    GET = 3,
    LIST = 4
}
export class WatchMessage extends pb_1.Message {
    #one_of_decls: number[][] = [];
    constructor(data?: any[] | {
        changes?: dependency_1.Struct;
        recordIds?: string[];
        event?: EventType;
        eventOn?: dependency_2.Timestamp;
    }) {
        super();
        pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [1], this.#one_of_decls);
        if (!Array.isArray(data) && typeof data == "object") {
            if ("changes" in data && data.changes != undefined) {
                this.changes = data.changes;
            }
            if ("recordIds" in data && data.recordIds != undefined) {
                this.recordIds = data.recordIds;
            }
            if ("event" in data && data.event != undefined) {
                this.event = data.event;
            }
            if ("eventOn" in data && data.eventOn != undefined) {
                this.eventOn = data.eventOn;
            }
        }
    }
    get changes() {
        return pb_1.Message.getWrapperField(this, dependency_1.Struct, 4) as dependency_1.Struct;
    }
    set changes(value: dependency_1.Struct) {
        pb_1.Message.setWrapperField(this, 4, value);
    }
    get hasChanges() {
        return pb_1.Message.getField(this, 4) != null;
    }
    get recordIds() {
        return pb_1.Message.getFieldWithDefault(this, 1, []) as string[];
    }
    set recordIds(value: string[]) {
        pb_1.Message.setField(this, 1, value);
    }
    get event() {
        return pb_1.Message.getFieldWithDefault(this, 2, EventType.CREATE) as EventType;
    }
    set event(value: EventType) {
        pb_1.Message.setField(this, 2, value);
    }
    get eventOn() {
        return pb_1.Message.getWrapperField(this, dependency_2.Timestamp, 3) as dependency_2.Timestamp;
    }
    set eventOn(value: dependency_2.Timestamp) {
        pb_1.Message.setWrapperField(this, 3, value);
    }
    get hasEventOn() {
        return pb_1.Message.getField(this, 3) != null;
    }
    static fromObject(data: {
        changes?: ReturnType<typeof dependency_1.Struct.prototype.toObject>;
        recordIds?: string[];
        event?: EventType;
        eventOn?: ReturnType<typeof dependency_2.Timestamp.prototype.toObject>;
    }): WatchMessage {
        const message = new WatchMessage({});
        if (data.changes != null) {
            message.changes = dependency_1.Struct.fromObject(data.changes);
        }
        if (data.recordIds != null) {
            message.recordIds = data.recordIds;
        }
        if (data.event != null) {
            message.event = data.event;
        }
        if (data.eventOn != null) {
            message.eventOn = dependency_2.Timestamp.fromObject(data.eventOn);
        }
        return message;
    }
    toObject() {
        const data: {
            changes?: ReturnType<typeof dependency_1.Struct.prototype.toObject>;
            recordIds?: string[];
            event?: EventType;
            eventOn?: ReturnType<typeof dependency_2.Timestamp.prototype.toObject>;
        } = {};
        if (this.changes != null) {
            data.changes = this.changes.toObject();
        }
        if (this.recordIds != null) {
            data.recordIds = this.recordIds;
        }
        if (this.event != null) {
            data.event = this.event;
        }
        if (this.eventOn != null) {
            data.eventOn = this.eventOn.toObject();
        }
        return data;
    }
    serialize(): Uint8Array;
    serialize(w: pb_1.BinaryWriter): void;
    serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
        const writer = w || new pb_1.BinaryWriter();
        if (this.hasChanges)
            writer.writeMessage(4, this.changes, () => this.changes.serialize(writer));
        if (this.recordIds.length)
            writer.writeRepeatedString(1, this.recordIds);
        if (this.event != EventType.CREATE)
            writer.writeEnum(2, this.event);
        if (this.hasEventOn)
            writer.writeMessage(3, this.eventOn, () => this.eventOn.serialize(writer));
        if (!w)
            return writer.getResultBuffer();
    }
    static deserialize(bytes: Uint8Array | pb_1.BinaryReader): WatchMessage {
        const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new WatchMessage();
        while (reader.nextField()) {
            if (reader.isEndGroup())
                break;
            switch (reader.getFieldNumber()) {
                case 4:
                    reader.readMessage(message.changes, () => message.changes = dependency_1.Struct.deserialize(reader));
                    break;
                case 1:
                    pb_1.Message.addToRepeatedField(message, 1, reader.readString());
                    break;
                case 2:
                    message.event = reader.readEnum();
                    break;
                case 3:
                    reader.readMessage(message.eventOn, () => message.eventOn = dependency_2.Timestamp.deserialize(reader));
                    break;
                default: reader.skipField();
            }
        }
        return message;
    }
    serializeBinary(): Uint8Array {
        return this.serialize();
    }
    static deserializeBinary(bytes: Uint8Array): WatchMessage {
        return WatchMessage.deserialize(bytes);
    }
}
