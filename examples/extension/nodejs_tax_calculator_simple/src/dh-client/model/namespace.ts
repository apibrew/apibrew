/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 0.0.0
 * source: model/namespace.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./audit";
import * as dependency_2 from "./common";
import * as dependency_3 from "./security";
import * as dependency_4 from "./../google/protobuf/struct";
import * as pb_1 from "google-protobuf";
export class Namespace extends pb_1.Message {
    #one_of_decls: number[][] = [];
    constructor(data?: any[] | {
        id?: string;
        name?: string;
        description?: string;
        details?: dependency_4.Struct;
        securityContext?: dependency_3.SecurityContext;
        auditData?: dependency_1.AuditData;
        version?: number;
    }) {
        super();
        pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
        if (!Array.isArray(data) && typeof data == "object") {
            if ("id" in data && data.id != undefined) {
                this.id = data.id;
            }
            if ("name" in data && data.name != undefined) {
                this.name = data.name;
            }
            if ("description" in data && data.description != undefined) {
                this.description = data.description;
            }
            if ("details" in data && data.details != undefined) {
                this.details = data.details;
            }
            if ("securityContext" in data && data.securityContext != undefined) {
                this.securityContext = data.securityContext;
            }
            if ("auditData" in data && data.auditData != undefined) {
                this.auditData = data.auditData;
            }
            if ("version" in data && data.version != undefined) {
                this.version = data.version;
            }
        }
    }
    get id() {
        return pb_1.Message.getFieldWithDefault(this, 1, "") as string;
    }
    set id(value: string) {
        pb_1.Message.setField(this, 1, value);
    }
    get name() {
        return pb_1.Message.getFieldWithDefault(this, 2, "") as string;
    }
    set name(value: string) {
        pb_1.Message.setField(this, 2, value);
    }
    get description() {
        return pb_1.Message.getFieldWithDefault(this, 4, "") as string;
    }
    set description(value: string) {
        pb_1.Message.setField(this, 4, value);
    }
    get details() {
        return pb_1.Message.getWrapperField(this, dependency_4.Struct, 5) as dependency_4.Struct;
    }
    set details(value: dependency_4.Struct) {
        pb_1.Message.setWrapperField(this, 5, value);
    }
    get hasDetails() {
        return pb_1.Message.getField(this, 5) != null;
    }
    get securityContext() {
        return pb_1.Message.getWrapperField(this, dependency_3.SecurityContext, 6) as dependency_3.SecurityContext;
    }
    set securityContext(value: dependency_3.SecurityContext) {
        pb_1.Message.setWrapperField(this, 6, value);
    }
    get hasSecurityContext() {
        return pb_1.Message.getField(this, 6) != null;
    }
    get auditData() {
        return pb_1.Message.getWrapperField(this, dependency_1.AuditData, 101) as dependency_1.AuditData;
    }
    set auditData(value: dependency_1.AuditData) {
        pb_1.Message.setWrapperField(this, 101, value);
    }
    get hasAuditData() {
        return pb_1.Message.getField(this, 101) != null;
    }
    get version() {
        return pb_1.Message.getFieldWithDefault(this, 102, 0) as number;
    }
    set version(value: number) {
        pb_1.Message.setField(this, 102, value);
    }
    static fromObject(data: {
        id?: string;
        name?: string;
        description?: string;
        details?: ReturnType<typeof dependency_4.Struct.prototype.toObject>;
        securityContext?: ReturnType<typeof dependency_3.SecurityContext.prototype.toObject>;
        auditData?: ReturnType<typeof dependency_1.AuditData.prototype.toObject>;
        version?: number;
    }): Namespace {
        const message = new Namespace({});
        if (data.id != null) {
            message.id = data.id;
        }
        if (data.name != null) {
            message.name = data.name;
        }
        if (data.description != null) {
            message.description = data.description;
        }
        if (data.details != null) {
            message.details = dependency_4.Struct.fromObject(data.details);
        }
        if (data.securityContext != null) {
            message.securityContext = dependency_3.SecurityContext.fromObject(data.securityContext);
        }
        if (data.auditData != null) {
            message.auditData = dependency_1.AuditData.fromObject(data.auditData);
        }
        if (data.version != null) {
            message.version = data.version;
        }
        return message;
    }
    toObject() {
        const data: {
            id?: string;
            name?: string;
            description?: string;
            details?: ReturnType<typeof dependency_4.Struct.prototype.toObject>;
            securityContext?: ReturnType<typeof dependency_3.SecurityContext.prototype.toObject>;
            auditData?: ReturnType<typeof dependency_1.AuditData.prototype.toObject>;
            version?: number;
        } = {};
        if (this.id != null) {
            data.id = this.id;
        }
        if (this.name != null) {
            data.name = this.name;
        }
        if (this.description != null) {
            data.description = this.description;
        }
        if (this.details != null) {
            data.details = this.details.toObject();
        }
        if (this.securityContext != null) {
            data.securityContext = this.securityContext.toObject();
        }
        if (this.auditData != null) {
            data.auditData = this.auditData.toObject();
        }
        if (this.version != null) {
            data.version = this.version;
        }
        return data;
    }
    serialize(): Uint8Array;
    serialize(w: pb_1.BinaryWriter): void;
    serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
        const writer = w || new pb_1.BinaryWriter();
        if (this.id.length)
            writer.writeString(1, this.id);
        if (this.name.length)
            writer.writeString(2, this.name);
        if (this.description.length)
            writer.writeString(4, this.description);
        if (this.hasDetails)
            writer.writeMessage(5, this.details, () => this.details.serialize(writer));
        if (this.hasSecurityContext)
            writer.writeMessage(6, this.securityContext, () => this.securityContext.serialize(writer));
        if (this.hasAuditData)
            writer.writeMessage(101, this.auditData, () => this.auditData.serialize(writer));
        if (this.version != 0)
            writer.writeUint32(102, this.version);
        if (!w)
            return writer.getResultBuffer();
    }
    static deserialize(bytes: Uint8Array | pb_1.BinaryReader): Namespace {
        const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new Namespace();
        while (reader.nextField()) {
            if (reader.isEndGroup())
                break;
            switch (reader.getFieldNumber()) {
                case 1:
                    message.id = reader.readString();
                    break;
                case 2:
                    message.name = reader.readString();
                    break;
                case 4:
                    message.description = reader.readString();
                    break;
                case 5:
                    reader.readMessage(message.details, () => message.details = dependency_4.Struct.deserialize(reader));
                    break;
                case 6:
                    reader.readMessage(message.securityContext, () => message.securityContext = dependency_3.SecurityContext.deserialize(reader));
                    break;
                case 101:
                    reader.readMessage(message.auditData, () => message.auditData = dependency_1.AuditData.deserialize(reader));
                    break;
                case 102:
                    message.version = reader.readUint32();
                    break;
                default: reader.skipField();
            }
        }
        return message;
    }
    serializeBinary(): Uint8Array {
        return this.serialize();
    }
    static deserializeBinary(bytes: Uint8Array): Namespace {
        return Namespace.deserialize(bytes);
    }
}
