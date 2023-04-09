/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 0.0.0
 * source: model/security.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./../google/protobuf/timestamp";
import * as dependency_2 from "./annotations";
import * as pb_1 from "google-protobuf";
export enum OperationType {
    OPERATION_TYPE_READ = 0,
    OPERATION_TYPE_CREATE = 1,
    OPERATION_TYPE_UPDATE = 2,
    OPERATION_TYPE_DELETE = 3,
    FULL = 4
}
export enum PermitType {
    PERMIT_TYPE_ALLOW = 0,
    PERMIT_TYPE_REJECT = 1,
    PERMIT_TYPE_UNKNOWN = 2
}
export class SecurityConstraint extends pb_1.Message {
    #one_of_decls: number[][] = [];
    constructor(data?: any[] | {
        namespace?: string;
        resource?: string;
        property?: string;
        before?: dependency_1.Timestamp;
        after?: dependency_1.Timestamp;
        principal?: string;
        recordIds?: string[];
        operation?: OperationType;
        permit?: PermitType;
    }) {
        super();
        pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [8], this.#one_of_decls);
        if (!Array.isArray(data) && typeof data == "object") {
            if ("namespace" in data && data.namespace != undefined) {
                this.namespace = data.namespace;
            }
            if ("resource" in data && data.resource != undefined) {
                this.resource = data.resource;
            }
            if ("property" in data && data.property != undefined) {
                this.property = data.property;
            }
            if ("before" in data && data.before != undefined) {
                this.before = data.before;
            }
            if ("after" in data && data.after != undefined) {
                this.after = data.after;
            }
            if ("principal" in data && data.principal != undefined) {
                this.principal = data.principal;
            }
            if ("recordIds" in data && data.recordIds != undefined) {
                this.recordIds = data.recordIds;
            }
            if ("operation" in data && data.operation != undefined) {
                this.operation = data.operation;
            }
            if ("permit" in data && data.permit != undefined) {
                this.permit = data.permit;
            }
        }
    }
    get namespace() {
        return pb_1.Message.getFieldWithDefault(this, 1, "") as string;
    }
    set namespace(value: string) {
        pb_1.Message.setField(this, 1, value);
    }
    get resource() {
        return pb_1.Message.getFieldWithDefault(this, 2, "") as string;
    }
    set resource(value: string) {
        pb_1.Message.setField(this, 2, value);
    }
    get property() {
        return pb_1.Message.getFieldWithDefault(this, 3, "") as string;
    }
    set property(value: string) {
        pb_1.Message.setField(this, 3, value);
    }
    get before() {
        return pb_1.Message.getWrapperField(this, dependency_1.Timestamp, 5) as dependency_1.Timestamp;
    }
    set before(value: dependency_1.Timestamp) {
        pb_1.Message.setWrapperField(this, 5, value);
    }
    get hasBefore() {
        return pb_1.Message.getField(this, 5) != null;
    }
    get after() {
        return pb_1.Message.getWrapperField(this, dependency_1.Timestamp, 6) as dependency_1.Timestamp;
    }
    set after(value: dependency_1.Timestamp) {
        pb_1.Message.setWrapperField(this, 6, value);
    }
    get hasAfter() {
        return pb_1.Message.getField(this, 6) != null;
    }
    get principal() {
        return pb_1.Message.getFieldWithDefault(this, 7, "") as string;
    }
    set principal(value: string) {
        pb_1.Message.setField(this, 7, value);
    }
    get recordIds() {
        return pb_1.Message.getFieldWithDefault(this, 8, []) as string[];
    }
    set recordIds(value: string[]) {
        pb_1.Message.setField(this, 8, value);
    }
    get operation() {
        return pb_1.Message.getFieldWithDefault(this, 13, OperationType.OPERATION_TYPE_READ) as OperationType;
    }
    set operation(value: OperationType) {
        pb_1.Message.setField(this, 13, value);
    }
    get permit() {
        return pb_1.Message.getFieldWithDefault(this, 14, PermitType.PERMIT_TYPE_ALLOW) as PermitType;
    }
    set permit(value: PermitType) {
        pb_1.Message.setField(this, 14, value);
    }
    static fromObject(data: {
        namespace?: string;
        resource?: string;
        property?: string;
        before?: ReturnType<typeof dependency_1.Timestamp.prototype.toObject>;
        after?: ReturnType<typeof dependency_1.Timestamp.prototype.toObject>;
        principal?: string;
        recordIds?: string[];
        operation?: OperationType;
        permit?: PermitType;
    }): SecurityConstraint {
        const message = new SecurityConstraint({});
        if (data.namespace != null) {
            message.namespace = data.namespace;
        }
        if (data.resource != null) {
            message.resource = data.resource;
        }
        if (data.property != null) {
            message.property = data.property;
        }
        if (data.before != null) {
            message.before = dependency_1.Timestamp.fromObject(data.before);
        }
        if (data.after != null) {
            message.after = dependency_1.Timestamp.fromObject(data.after);
        }
        if (data.principal != null) {
            message.principal = data.principal;
        }
        if (data.recordIds != null) {
            message.recordIds = data.recordIds;
        }
        if (data.operation != null) {
            message.operation = data.operation;
        }
        if (data.permit != null) {
            message.permit = data.permit;
        }
        return message;
    }
    toObject() {
        const data: {
            namespace?: string;
            resource?: string;
            property?: string;
            before?: ReturnType<typeof dependency_1.Timestamp.prototype.toObject>;
            after?: ReturnType<typeof dependency_1.Timestamp.prototype.toObject>;
            principal?: string;
            recordIds?: string[];
            operation?: OperationType;
            permit?: PermitType;
        } = {};
        if (this.namespace != null) {
            data.namespace = this.namespace;
        }
        if (this.resource != null) {
            data.resource = this.resource;
        }
        if (this.property != null) {
            data.property = this.property;
        }
        if (this.before != null) {
            data.before = this.before.toObject();
        }
        if (this.after != null) {
            data.after = this.after.toObject();
        }
        if (this.principal != null) {
            data.principal = this.principal;
        }
        if (this.recordIds != null) {
            data.recordIds = this.recordIds;
        }
        if (this.operation != null) {
            data.operation = this.operation;
        }
        if (this.permit != null) {
            data.permit = this.permit;
        }
        return data;
    }
    serialize(): Uint8Array;
    serialize(w: pb_1.BinaryWriter): void;
    serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
        const writer = w || new pb_1.BinaryWriter();
        if (this.namespace.length)
            writer.writeString(1, this.namespace);
        if (this.resource.length)
            writer.writeString(2, this.resource);
        if (this.property.length)
            writer.writeString(3, this.property);
        if (this.hasBefore)
            writer.writeMessage(5, this.before, () => this.before.serialize(writer));
        if (this.hasAfter)
            writer.writeMessage(6, this.after, () => this.after.serialize(writer));
        if (this.principal.length)
            writer.writeString(7, this.principal);
        if (this.recordIds.length)
            writer.writeRepeatedString(8, this.recordIds);
        if (this.operation != OperationType.OPERATION_TYPE_READ)
            writer.writeEnum(13, this.operation);
        if (this.permit != PermitType.PERMIT_TYPE_ALLOW)
            writer.writeEnum(14, this.permit);
        if (!w)
            return writer.getResultBuffer();
    }
    static deserialize(bytes: Uint8Array | pb_1.BinaryReader): SecurityConstraint {
        const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new SecurityConstraint();
        while (reader.nextField()) {
            if (reader.isEndGroup())
                break;
            switch (reader.getFieldNumber()) {
                case 1:
                    message.namespace = reader.readString();
                    break;
                case 2:
                    message.resource = reader.readString();
                    break;
                case 3:
                    message.property = reader.readString();
                    break;
                case 5:
                    reader.readMessage(message.before, () => message.before = dependency_1.Timestamp.deserialize(reader));
                    break;
                case 6:
                    reader.readMessage(message.after, () => message.after = dependency_1.Timestamp.deserialize(reader));
                    break;
                case 7:
                    message.principal = reader.readString();
                    break;
                case 8:
                    pb_1.Message.addToRepeatedField(message, 8, reader.readString());
                    break;
                case 13:
                    message.operation = reader.readEnum();
                    break;
                case 14:
                    message.permit = reader.readEnum();
                    break;
                default: reader.skipField();
            }
        }
        return message;
    }
    serializeBinary(): Uint8Array {
        return this.serialize();
    }
    static deserializeBinary(bytes: Uint8Array): SecurityConstraint {
        return SecurityConstraint.deserialize(bytes);
    }
}
export class SecurityContext extends pb_1.Message {
    #one_of_decls: number[][] = [];
    constructor(data?: any[] | {
        constraints?: SecurityConstraint[];
    }) {
        super();
        pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [1], this.#one_of_decls);
        if (!Array.isArray(data) && typeof data == "object") {
            if ("constraints" in data && data.constraints != undefined) {
                this.constraints = data.constraints;
            }
        }
    }
    get constraints() {
        return pb_1.Message.getRepeatedWrapperField(this, SecurityConstraint, 1) as SecurityConstraint[];
    }
    set constraints(value: SecurityConstraint[]) {
        pb_1.Message.setRepeatedWrapperField(this, 1, value);
    }
    static fromObject(data: {
        constraints?: ReturnType<typeof SecurityConstraint.prototype.toObject>[];
    }): SecurityContext {
        const message = new SecurityContext({});
        if (data.constraints != null) {
            message.constraints = data.constraints.map(item => SecurityConstraint.fromObject(item));
        }
        return message;
    }
    toObject() {
        const data: {
            constraints?: ReturnType<typeof SecurityConstraint.prototype.toObject>[];
        } = {};
        if (this.constraints != null) {
            data.constraints = this.constraints.map((item: SecurityConstraint) => item.toObject());
        }
        return data;
    }
    serialize(): Uint8Array;
    serialize(w: pb_1.BinaryWriter): void;
    serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
        const writer = w || new pb_1.BinaryWriter();
        if (this.constraints.length)
            writer.writeRepeatedMessage(1, this.constraints, (item: SecurityConstraint) => item.serialize(writer));
        if (!w)
            return writer.getResultBuffer();
    }
    static deserialize(bytes: Uint8Array | pb_1.BinaryReader): SecurityContext {
        const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new SecurityContext();
        while (reader.nextField()) {
            if (reader.isEndGroup())
                break;
            switch (reader.getFieldNumber()) {
                case 1:
                    reader.readMessage(message.constraints, () => pb_1.Message.addToRepeatedWrapperField(message, 1, SecurityConstraint.deserialize(reader), SecurityConstraint));
                    break;
                default: reader.skipField();
            }
        }
        return message;
    }
    serializeBinary(): Uint8Array {
        return this.serialize();
    }
    static deserializeBinary(bytes: Uint8Array): SecurityContext {
        return SecurityContext.deserialize(bytes);
    }
}
