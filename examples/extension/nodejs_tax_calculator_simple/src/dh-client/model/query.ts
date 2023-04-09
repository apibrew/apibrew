/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 0.0.0
 * source: model/query.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./../google/protobuf/struct";
import * as pb_1 from "google-protobuf";
export class CompoundBooleanExpression extends pb_1.Message {
    #one_of_decls: number[][] = [];
    constructor(data?: any[] | {
        expressions?: BooleanExpression[];
    }) {
        super();
        pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [1], this.#one_of_decls);
        if (!Array.isArray(data) && typeof data == "object") {
            if ("expressions" in data && data.expressions != undefined) {
                this.expressions = data.expressions;
            }
        }
    }
    get expressions() {
        return pb_1.Message.getRepeatedWrapperField(this, BooleanExpression, 1) as BooleanExpression[];
    }
    set expressions(value: BooleanExpression[]) {
        pb_1.Message.setRepeatedWrapperField(this, 1, value);
    }
    static fromObject(data: {
        expressions?: ReturnType<typeof BooleanExpression.prototype.toObject>[];
    }): CompoundBooleanExpression {
        const message = new CompoundBooleanExpression({});
        if (data.expressions != null) {
            message.expressions = data.expressions.map(item => BooleanExpression.fromObject(item));
        }
        return message;
    }
    toObject() {
        const data: {
            expressions?: ReturnType<typeof BooleanExpression.prototype.toObject>[];
        } = {};
        if (this.expressions != null) {
            data.expressions = this.expressions.map((item: BooleanExpression) => item.toObject());
        }
        return data;
    }
    serialize(): Uint8Array;
    serialize(w: pb_1.BinaryWriter): void;
    serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
        const writer = w || new pb_1.BinaryWriter();
        if (this.expressions.length)
            writer.writeRepeatedMessage(1, this.expressions, (item: BooleanExpression) => item.serialize(writer));
        if (!w)
            return writer.getResultBuffer();
    }
    static deserialize(bytes: Uint8Array | pb_1.BinaryReader): CompoundBooleanExpression {
        const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new CompoundBooleanExpression();
        while (reader.nextField()) {
            if (reader.isEndGroup())
                break;
            switch (reader.getFieldNumber()) {
                case 1:
                    reader.readMessage(message.expressions, () => pb_1.Message.addToRepeatedWrapperField(message, 1, BooleanExpression.deserialize(reader), BooleanExpression));
                    break;
                default: reader.skipField();
            }
        }
        return message;
    }
    serializeBinary(): Uint8Array {
        return this.serialize();
    }
    static deserializeBinary(bytes: Uint8Array): CompoundBooleanExpression {
        return CompoundBooleanExpression.deserialize(bytes);
    }
}
export class RefValue extends pb_1.Message {
    #one_of_decls: number[][] = [];
    constructor(data?: any[] | {
        namespace?: string;
        resource?: string;
        properties?: Map<string, dependency_1.Value>;
    }) {
        super();
        pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
        if (!Array.isArray(data) && typeof data == "object") {
            if ("namespace" in data && data.namespace != undefined) {
                this.namespace = data.namespace;
            }
            if ("resource" in data && data.resource != undefined) {
                this.resource = data.resource;
            }
            if ("properties" in data && data.properties != undefined) {
                this.properties = data.properties;
            }
        }
        if (!this.properties)
            this.properties = new Map();
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
    get properties() {
        return pb_1.Message.getField(this, 3) as any as Map<string, dependency_1.Value>;
    }
    set properties(value: Map<string, dependency_1.Value>) {
        pb_1.Message.setField(this, 3, value as any);
    }
    static fromObject(data: {
        namespace?: string;
        resource?: string;
        properties?: {
            [key: string]: ReturnType<typeof dependency_1.Value.prototype.toObject>;
        };
    }): RefValue {
        const message = new RefValue({});
        if (data.namespace != null) {
            message.namespace = data.namespace;
        }
        if (data.resource != null) {
            message.resource = data.resource;
        }
        if (typeof data.properties == "object") {
            message.properties = new Map(Object.entries(data.properties).map(([key, value]) => [key, dependency_1.Value.fromObject(value)]));
        }
        return message;
    }
    toObject() {
        const data: {
            namespace?: string;
            resource?: string;
            properties?: {
                [key: string]: ReturnType<typeof dependency_1.Value.prototype.toObject>;
            };
        } = {};
        if (this.namespace != null) {
            data.namespace = this.namespace;
        }
        if (this.resource != null) {
            data.resource = this.resource;
        }
        if (this.properties != null) {
            data.properties = (Object.fromEntries)((Array.from)(this.properties).map(([key, value]) => [key, value.toObject()]));
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
        for (const [key, value] of this.properties) {
            writer.writeMessage(3, this.properties, () => {
                writer.writeString(1, key);
                writer.writeMessage(2, value, () => value.serialize(writer));
            });
        }
        if (!w)
            return writer.getResultBuffer();
    }
    static deserialize(bytes: Uint8Array | pb_1.BinaryReader): RefValue {
        const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new RefValue();
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
                    reader.readMessage(message, () => pb_1.Map.deserializeBinary(message.properties as any, reader, reader.readString, () => {
                        let value;
                        reader.readMessage(message, () => value = dependency_1.Value.deserialize(reader));
                        return value;
                    }));
                    break;
                default: reader.skipField();
            }
        }
        return message;
    }
    serializeBinary(): Uint8Array {
        return this.serialize();
    }
    static deserializeBinary(bytes: Uint8Array): RefValue {
        return RefValue.deserialize(bytes);
    }
}
export class Expression extends pb_1.Message {
    #one_of_decls: number[][] = [[1, 3, 4]];
    constructor(data?: any[] | ({} & (({
        property?: string;
        value?: never;
        refValue?: never;
    } | {
        property?: never;
        value?: dependency_1.Value;
        refValue?: never;
    } | {
        property?: never;
        value?: never;
        refValue?: RefValue;
    })))) {
        super();
        pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
        if (!Array.isArray(data) && typeof data == "object") {
            if ("property" in data && data.property != undefined) {
                this.property = data.property;
            }
            if ("value" in data && data.value != undefined) {
                this.value = data.value;
            }
            if ("refValue" in data && data.refValue != undefined) {
                this.refValue = data.refValue;
            }
        }
    }
    get property() {
        return pb_1.Message.getFieldWithDefault(this, 1, "") as string;
    }
    set property(value: string) {
        pb_1.Message.setOneofField(this, 1, this.#one_of_decls[0], value);
    }
    get hasProperty() {
        return pb_1.Message.getField(this, 1) != null;
    }
    get value() {
        return pb_1.Message.getWrapperField(this, dependency_1.Value, 3) as dependency_1.Value;
    }
    set value(value: dependency_1.Value) {
        pb_1.Message.setOneofWrapperField(this, 3, this.#one_of_decls[0], value);
    }
    get hasValue() {
        return pb_1.Message.getField(this, 3) != null;
    }
    get refValue() {
        return pb_1.Message.getWrapperField(this, RefValue, 4) as RefValue;
    }
    set refValue(value: RefValue) {
        pb_1.Message.setOneofWrapperField(this, 4, this.#one_of_decls[0], value);
    }
    get hasRefValue() {
        return pb_1.Message.getField(this, 4) != null;
    }
    get expression() {
        const cases: {
            [index: number]: "none" | "property" | "value" | "refValue";
        } = {
            0: "none",
            1: "property",
            3: "value",
            4: "refValue"
        };
        return cases[pb_1.Message.computeOneofCase(this, [1, 3, 4])];
    }
    static fromObject(data: {
        property?: string;
        value?: ReturnType<typeof dependency_1.Value.prototype.toObject>;
        refValue?: ReturnType<typeof RefValue.prototype.toObject>;
    }): Expression {
        const message = new Expression({});
        if (data.property != null) {
            message.property = data.property;
        }
        if (data.value != null) {
            message.value = dependency_1.Value.fromObject(data.value);
        }
        if (data.refValue != null) {
            message.refValue = RefValue.fromObject(data.refValue);
        }
        return message;
    }
    toObject() {
        const data: {
            property?: string;
            value?: ReturnType<typeof dependency_1.Value.prototype.toObject>;
            refValue?: ReturnType<typeof RefValue.prototype.toObject>;
        } = {};
        if (this.property != null) {
            data.property = this.property;
        }
        if (this.value != null) {
            data.value = this.value.toObject();
        }
        if (this.refValue != null) {
            data.refValue = this.refValue.toObject();
        }
        return data;
    }
    serialize(): Uint8Array;
    serialize(w: pb_1.BinaryWriter): void;
    serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
        const writer = w || new pb_1.BinaryWriter();
        if (this.hasProperty)
            writer.writeString(1, this.property);
        if (this.hasValue)
            writer.writeMessage(3, this.value, () => this.value.serialize(writer));
        if (this.hasRefValue)
            writer.writeMessage(4, this.refValue, () => this.refValue.serialize(writer));
        if (!w)
            return writer.getResultBuffer();
    }
    static deserialize(bytes: Uint8Array | pb_1.BinaryReader): Expression {
        const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new Expression();
        while (reader.nextField()) {
            if (reader.isEndGroup())
                break;
            switch (reader.getFieldNumber()) {
                case 1:
                    message.property = reader.readString();
                    break;
                case 3:
                    reader.readMessage(message.value, () => message.value = dependency_1.Value.deserialize(reader));
                    break;
                case 4:
                    reader.readMessage(message.refValue, () => message.refValue = RefValue.deserialize(reader));
                    break;
                default: reader.skipField();
            }
        }
        return message;
    }
    serializeBinary(): Uint8Array {
        return this.serialize();
    }
    static deserializeBinary(bytes: Uint8Array): Expression {
        return Expression.deserialize(bytes);
    }
}
export class PairExpression extends pb_1.Message {
    #one_of_decls: number[][] = [];
    constructor(data?: any[] | {
        left?: Expression;
        right?: Expression;
    }) {
        super();
        pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
        if (!Array.isArray(data) && typeof data == "object") {
            if ("left" in data && data.left != undefined) {
                this.left = data.left;
            }
            if ("right" in data && data.right != undefined) {
                this.right = data.right;
            }
        }
    }
    get left() {
        return pb_1.Message.getWrapperField(this, Expression, 1) as Expression;
    }
    set left(value: Expression) {
        pb_1.Message.setWrapperField(this, 1, value);
    }
    get hasLeft() {
        return pb_1.Message.getField(this, 1) != null;
    }
    get right() {
        return pb_1.Message.getWrapperField(this, Expression, 2) as Expression;
    }
    set right(value: Expression) {
        pb_1.Message.setWrapperField(this, 2, value);
    }
    get hasRight() {
        return pb_1.Message.getField(this, 2) != null;
    }
    static fromObject(data: {
        left?: ReturnType<typeof Expression.prototype.toObject>;
        right?: ReturnType<typeof Expression.prototype.toObject>;
    }): PairExpression {
        const message = new PairExpression({});
        if (data.left != null) {
            message.left = Expression.fromObject(data.left);
        }
        if (data.right != null) {
            message.right = Expression.fromObject(data.right);
        }
        return message;
    }
    toObject() {
        const data: {
            left?: ReturnType<typeof Expression.prototype.toObject>;
            right?: ReturnType<typeof Expression.prototype.toObject>;
        } = {};
        if (this.left != null) {
            data.left = this.left.toObject();
        }
        if (this.right != null) {
            data.right = this.right.toObject();
        }
        return data;
    }
    serialize(): Uint8Array;
    serialize(w: pb_1.BinaryWriter): void;
    serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
        const writer = w || new pb_1.BinaryWriter();
        if (this.hasLeft)
            writer.writeMessage(1, this.left, () => this.left.serialize(writer));
        if (this.hasRight)
            writer.writeMessage(2, this.right, () => this.right.serialize(writer));
        if (!w)
            return writer.getResultBuffer();
    }
    static deserialize(bytes: Uint8Array | pb_1.BinaryReader): PairExpression {
        const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new PairExpression();
        while (reader.nextField()) {
            if (reader.isEndGroup())
                break;
            switch (reader.getFieldNumber()) {
                case 1:
                    reader.readMessage(message.left, () => message.left = Expression.deserialize(reader));
                    break;
                case 2:
                    reader.readMessage(message.right, () => message.right = Expression.deserialize(reader));
                    break;
                default: reader.skipField();
            }
        }
        return message;
    }
    serializeBinary(): Uint8Array {
        return this.serialize();
    }
    static deserializeBinary(bytes: Uint8Array): PairExpression {
        return PairExpression.deserialize(bytes);
    }
}
export class RegexMatchExpression extends pb_1.Message {
    #one_of_decls: number[][] = [];
    constructor(data?: any[] | {
        pattern?: string;
        expression?: Expression;
    }) {
        super();
        pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
        if (!Array.isArray(data) && typeof data == "object") {
            if ("pattern" in data && data.pattern != undefined) {
                this.pattern = data.pattern;
            }
            if ("expression" in data && data.expression != undefined) {
                this.expression = data.expression;
            }
        }
    }
    get pattern() {
        return pb_1.Message.getFieldWithDefault(this, 1, "") as string;
    }
    set pattern(value: string) {
        pb_1.Message.setField(this, 1, value);
    }
    get expression() {
        return pb_1.Message.getWrapperField(this, Expression, 2) as Expression;
    }
    set expression(value: Expression) {
        pb_1.Message.setWrapperField(this, 2, value);
    }
    get hasExpression() {
        return pb_1.Message.getField(this, 2) != null;
    }
    static fromObject(data: {
        pattern?: string;
        expression?: ReturnType<typeof Expression.prototype.toObject>;
    }): RegexMatchExpression {
        const message = new RegexMatchExpression({});
        if (data.pattern != null) {
            message.pattern = data.pattern;
        }
        if (data.expression != null) {
            message.expression = Expression.fromObject(data.expression);
        }
        return message;
    }
    toObject() {
        const data: {
            pattern?: string;
            expression?: ReturnType<typeof Expression.prototype.toObject>;
        } = {};
        if (this.pattern != null) {
            data.pattern = this.pattern;
        }
        if (this.expression != null) {
            data.expression = this.expression.toObject();
        }
        return data;
    }
    serialize(): Uint8Array;
    serialize(w: pb_1.BinaryWriter): void;
    serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
        const writer = w || new pb_1.BinaryWriter();
        if (this.pattern.length)
            writer.writeString(1, this.pattern);
        if (this.hasExpression)
            writer.writeMessage(2, this.expression, () => this.expression.serialize(writer));
        if (!w)
            return writer.getResultBuffer();
    }
    static deserialize(bytes: Uint8Array | pb_1.BinaryReader): RegexMatchExpression {
        const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new RegexMatchExpression();
        while (reader.nextField()) {
            if (reader.isEndGroup())
                break;
            switch (reader.getFieldNumber()) {
                case 1:
                    message.pattern = reader.readString();
                    break;
                case 2:
                    reader.readMessage(message.expression, () => message.expression = Expression.deserialize(reader));
                    break;
                default: reader.skipField();
            }
        }
        return message;
    }
    serializeBinary(): Uint8Array {
        return this.serialize();
    }
    static deserializeBinary(bytes: Uint8Array): RegexMatchExpression {
        return RegexMatchExpression.deserialize(bytes);
    }
}
export class BooleanExpression extends pb_1.Message {
    #one_of_decls: number[][] = [[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11]];
    constructor(data?: any[] | ({} & (({
        and?: CompoundBooleanExpression;
        or?: never;
        not?: never;
        equal?: never;
        lessThan?: never;
        greaterThan?: never;
        lessThanOrEqual?: never;
        greaterThanOrEqual?: never;
        in?: never;
        isNull?: never;
        regexMatch?: never;
    } | {
        and?: never;
        or?: CompoundBooleanExpression;
        not?: never;
        equal?: never;
        lessThan?: never;
        greaterThan?: never;
        lessThanOrEqual?: never;
        greaterThanOrEqual?: never;
        in?: never;
        isNull?: never;
        regexMatch?: never;
    } | {
        and?: never;
        or?: never;
        not?: BooleanExpression;
        equal?: never;
        lessThan?: never;
        greaterThan?: never;
        lessThanOrEqual?: never;
        greaterThanOrEqual?: never;
        in?: never;
        isNull?: never;
        regexMatch?: never;
    } | {
        and?: never;
        or?: never;
        not?: never;
        equal?: PairExpression;
        lessThan?: never;
        greaterThan?: never;
        lessThanOrEqual?: never;
        greaterThanOrEqual?: never;
        in?: never;
        isNull?: never;
        regexMatch?: never;
    } | {
        and?: never;
        or?: never;
        not?: never;
        equal?: never;
        lessThan?: PairExpression;
        greaterThan?: never;
        lessThanOrEqual?: never;
        greaterThanOrEqual?: never;
        in?: never;
        isNull?: never;
        regexMatch?: never;
    } | {
        and?: never;
        or?: never;
        not?: never;
        equal?: never;
        lessThan?: never;
        greaterThan?: PairExpression;
        lessThanOrEqual?: never;
        greaterThanOrEqual?: never;
        in?: never;
        isNull?: never;
        regexMatch?: never;
    } | {
        and?: never;
        or?: never;
        not?: never;
        equal?: never;
        lessThan?: never;
        greaterThan?: never;
        lessThanOrEqual?: PairExpression;
        greaterThanOrEqual?: never;
        in?: never;
        isNull?: never;
        regexMatch?: never;
    } | {
        and?: never;
        or?: never;
        not?: never;
        equal?: never;
        lessThan?: never;
        greaterThan?: never;
        lessThanOrEqual?: never;
        greaterThanOrEqual?: PairExpression;
        in?: never;
        isNull?: never;
        regexMatch?: never;
    } | {
        and?: never;
        or?: never;
        not?: never;
        equal?: never;
        lessThan?: never;
        greaterThan?: never;
        lessThanOrEqual?: never;
        greaterThanOrEqual?: never;
        in?: PairExpression;
        isNull?: never;
        regexMatch?: never;
    } | {
        and?: never;
        or?: never;
        not?: never;
        equal?: never;
        lessThan?: never;
        greaterThan?: never;
        lessThanOrEqual?: never;
        greaterThanOrEqual?: never;
        in?: never;
        isNull?: Expression;
        regexMatch?: never;
    } | {
        and?: never;
        or?: never;
        not?: never;
        equal?: never;
        lessThan?: never;
        greaterThan?: never;
        lessThanOrEqual?: never;
        greaterThanOrEqual?: never;
        in?: never;
        isNull?: never;
        regexMatch?: RegexMatchExpression;
    })))) {
        super();
        pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
        if (!Array.isArray(data) && typeof data == "object") {
            if ("and" in data && data.and != undefined) {
                this.and = data.and;
            }
            if ("or" in data && data.or != undefined) {
                this.or = data.or;
            }
            if ("not" in data && data.not != undefined) {
                this.not = data.not;
            }
            if ("equal" in data && data.equal != undefined) {
                this.equal = data.equal;
            }
            if ("lessThan" in data && data.lessThan != undefined) {
                this.lessThan = data.lessThan;
            }
            if ("greaterThan" in data && data.greaterThan != undefined) {
                this.greaterThan = data.greaterThan;
            }
            if ("lessThanOrEqual" in data && data.lessThanOrEqual != undefined) {
                this.lessThanOrEqual = data.lessThanOrEqual;
            }
            if ("greaterThanOrEqual" in data && data.greaterThanOrEqual != undefined) {
                this.greaterThanOrEqual = data.greaterThanOrEqual;
            }
            if ("in" in data && data.in != undefined) {
                this.in = data.in;
            }
            if ("isNull" in data && data.isNull != undefined) {
                this.isNull = data.isNull;
            }
            if ("regexMatch" in data && data.regexMatch != undefined) {
                this.regexMatch = data.regexMatch;
            }
        }
    }
    get and() {
        return pb_1.Message.getWrapperField(this, CompoundBooleanExpression, 1) as CompoundBooleanExpression;
    }
    set and(value: CompoundBooleanExpression) {
        pb_1.Message.setOneofWrapperField(this, 1, this.#one_of_decls[0], value);
    }
    get hasAnd() {
        return pb_1.Message.getField(this, 1) != null;
    }
    get or() {
        return pb_1.Message.getWrapperField(this, CompoundBooleanExpression, 2) as CompoundBooleanExpression;
    }
    set or(value: CompoundBooleanExpression) {
        pb_1.Message.setOneofWrapperField(this, 2, this.#one_of_decls[0], value);
    }
    get hasOr() {
        return pb_1.Message.getField(this, 2) != null;
    }
    get not() {
        return pb_1.Message.getWrapperField(this, BooleanExpression, 3) as BooleanExpression;
    }
    set not(value: BooleanExpression) {
        pb_1.Message.setOneofWrapperField(this, 3, this.#one_of_decls[0], value);
    }
    get hasNot() {
        return pb_1.Message.getField(this, 3) != null;
    }
    get equal() {
        return pb_1.Message.getWrapperField(this, PairExpression, 4) as PairExpression;
    }
    set equal(value: PairExpression) {
        pb_1.Message.setOneofWrapperField(this, 4, this.#one_of_decls[0], value);
    }
    get hasEqual() {
        return pb_1.Message.getField(this, 4) != null;
    }
    get lessThan() {
        return pb_1.Message.getWrapperField(this, PairExpression, 5) as PairExpression;
    }
    set lessThan(value: PairExpression) {
        pb_1.Message.setOneofWrapperField(this, 5, this.#one_of_decls[0], value);
    }
    get hasLessThan() {
        return pb_1.Message.getField(this, 5) != null;
    }
    get greaterThan() {
        return pb_1.Message.getWrapperField(this, PairExpression, 6) as PairExpression;
    }
    set greaterThan(value: PairExpression) {
        pb_1.Message.setOneofWrapperField(this, 6, this.#one_of_decls[0], value);
    }
    get hasGreaterThan() {
        return pb_1.Message.getField(this, 6) != null;
    }
    get lessThanOrEqual() {
        return pb_1.Message.getWrapperField(this, PairExpression, 7) as PairExpression;
    }
    set lessThanOrEqual(value: PairExpression) {
        pb_1.Message.setOneofWrapperField(this, 7, this.#one_of_decls[0], value);
    }
    get hasLessThanOrEqual() {
        return pb_1.Message.getField(this, 7) != null;
    }
    get greaterThanOrEqual() {
        return pb_1.Message.getWrapperField(this, PairExpression, 8) as PairExpression;
    }
    set greaterThanOrEqual(value: PairExpression) {
        pb_1.Message.setOneofWrapperField(this, 8, this.#one_of_decls[0], value);
    }
    get hasGreaterThanOrEqual() {
        return pb_1.Message.getField(this, 8) != null;
    }
    get in() {
        return pb_1.Message.getWrapperField(this, PairExpression, 9) as PairExpression;
    }
    set in(value: PairExpression) {
        pb_1.Message.setOneofWrapperField(this, 9, this.#one_of_decls[0], value);
    }
    get hasIn() {
        return pb_1.Message.getField(this, 9) != null;
    }
    get isNull() {
        return pb_1.Message.getWrapperField(this, Expression, 10) as Expression;
    }
    set isNull(value: Expression) {
        pb_1.Message.setOneofWrapperField(this, 10, this.#one_of_decls[0], value);
    }
    get hasIsNull() {
        return pb_1.Message.getField(this, 10) != null;
    }
    get regexMatch() {
        return pb_1.Message.getWrapperField(this, RegexMatchExpression, 11) as RegexMatchExpression;
    }
    set regexMatch(value: RegexMatchExpression) {
        pb_1.Message.setOneofWrapperField(this, 11, this.#one_of_decls[0], value);
    }
    get hasRegexMatch() {
        return pb_1.Message.getField(this, 11) != null;
    }
    get expression() {
        const cases: {
            [index: number]: "none" | "and" | "or" | "not" | "equal" | "lessThan" | "greaterThan" | "lessThanOrEqual" | "greaterThanOrEqual" | "in" | "isNull" | "regexMatch";
        } = {
            0: "none",
            1: "and",
            2: "or",
            3: "not",
            4: "equal",
            5: "lessThan",
            6: "greaterThan",
            7: "lessThanOrEqual",
            8: "greaterThanOrEqual",
            9: "in",
            10: "isNull",
            11: "regexMatch"
        };
        return cases[pb_1.Message.computeOneofCase(this, [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11])];
    }
    static fromObject(data: {
        and?: ReturnType<typeof CompoundBooleanExpression.prototype.toObject>;
        or?: ReturnType<typeof CompoundBooleanExpression.prototype.toObject>;
        not?: ReturnType<typeof BooleanExpression.prototype.toObject>;
        equal?: ReturnType<typeof PairExpression.prototype.toObject>;
        lessThan?: ReturnType<typeof PairExpression.prototype.toObject>;
        greaterThan?: ReturnType<typeof PairExpression.prototype.toObject>;
        lessThanOrEqual?: ReturnType<typeof PairExpression.prototype.toObject>;
        greaterThanOrEqual?: ReturnType<typeof PairExpression.prototype.toObject>;
        in?: ReturnType<typeof PairExpression.prototype.toObject>;
        isNull?: ReturnType<typeof Expression.prototype.toObject>;
        regexMatch?: ReturnType<typeof RegexMatchExpression.prototype.toObject>;
    }): BooleanExpression {
        const message = new BooleanExpression({});
        if (data.and != null) {
            message.and = CompoundBooleanExpression.fromObject(data.and);
        }
        if (data.or != null) {
            message.or = CompoundBooleanExpression.fromObject(data.or);
        }
        if (data.not != null) {
            message.not = BooleanExpression.fromObject(data.not);
        }
        if (data.equal != null) {
            message.equal = PairExpression.fromObject(data.equal);
        }
        if (data.lessThan != null) {
            message.lessThan = PairExpression.fromObject(data.lessThan);
        }
        if (data.greaterThan != null) {
            message.greaterThan = PairExpression.fromObject(data.greaterThan);
        }
        if (data.lessThanOrEqual != null) {
            message.lessThanOrEqual = PairExpression.fromObject(data.lessThanOrEqual);
        }
        if (data.greaterThanOrEqual != null) {
            message.greaterThanOrEqual = PairExpression.fromObject(data.greaterThanOrEqual);
        }
        if (data.in != null) {
            message.in = PairExpression.fromObject(data.in);
        }
        if (data.isNull != null) {
            message.isNull = Expression.fromObject(data.isNull);
        }
        if (data.regexMatch != null) {
            message.regexMatch = RegexMatchExpression.fromObject(data.regexMatch);
        }
        return message;
    }
    toObject() {
        const data: {
            and?: ReturnType<typeof CompoundBooleanExpression.prototype.toObject>;
            or?: ReturnType<typeof CompoundBooleanExpression.prototype.toObject>;
            not?: ReturnType<typeof BooleanExpression.prototype.toObject>;
            equal?: ReturnType<typeof PairExpression.prototype.toObject>;
            lessThan?: ReturnType<typeof PairExpression.prototype.toObject>;
            greaterThan?: ReturnType<typeof PairExpression.prototype.toObject>;
            lessThanOrEqual?: ReturnType<typeof PairExpression.prototype.toObject>;
            greaterThanOrEqual?: ReturnType<typeof PairExpression.prototype.toObject>;
            in?: ReturnType<typeof PairExpression.prototype.toObject>;
            isNull?: ReturnType<typeof Expression.prototype.toObject>;
            regexMatch?: ReturnType<typeof RegexMatchExpression.prototype.toObject>;
        } = {};
        if (this.and != null) {
            data.and = this.and.toObject();
        }
        if (this.or != null) {
            data.or = this.or.toObject();
        }
        if (this.not != null) {
            data.not = this.not.toObject();
        }
        if (this.equal != null) {
            data.equal = this.equal.toObject();
        }
        if (this.lessThan != null) {
            data.lessThan = this.lessThan.toObject();
        }
        if (this.greaterThan != null) {
            data.greaterThan = this.greaterThan.toObject();
        }
        if (this.lessThanOrEqual != null) {
            data.lessThanOrEqual = this.lessThanOrEqual.toObject();
        }
        if (this.greaterThanOrEqual != null) {
            data.greaterThanOrEqual = this.greaterThanOrEqual.toObject();
        }
        if (this.in != null) {
            data.in = this.in.toObject();
        }
        if (this.isNull != null) {
            data.isNull = this.isNull.toObject();
        }
        if (this.regexMatch != null) {
            data.regexMatch = this.regexMatch.toObject();
        }
        return data;
    }
    serialize(): Uint8Array;
    serialize(w: pb_1.BinaryWriter): void;
    serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
        const writer = w || new pb_1.BinaryWriter();
        if (this.hasAnd)
            writer.writeMessage(1, this.and, () => this.and.serialize(writer));
        if (this.hasOr)
            writer.writeMessage(2, this.or, () => this.or.serialize(writer));
        if (this.hasNot)
            writer.writeMessage(3, this.not, () => this.not.serialize(writer));
        if (this.hasEqual)
            writer.writeMessage(4, this.equal, () => this.equal.serialize(writer));
        if (this.hasLessThan)
            writer.writeMessage(5, this.lessThan, () => this.lessThan.serialize(writer));
        if (this.hasGreaterThan)
            writer.writeMessage(6, this.greaterThan, () => this.greaterThan.serialize(writer));
        if (this.hasLessThanOrEqual)
            writer.writeMessage(7, this.lessThanOrEqual, () => this.lessThanOrEqual.serialize(writer));
        if (this.hasGreaterThanOrEqual)
            writer.writeMessage(8, this.greaterThanOrEqual, () => this.greaterThanOrEqual.serialize(writer));
        if (this.hasIn)
            writer.writeMessage(9, this.in, () => this.in.serialize(writer));
        if (this.hasIsNull)
            writer.writeMessage(10, this.isNull, () => this.isNull.serialize(writer));
        if (this.hasRegexMatch)
            writer.writeMessage(11, this.regexMatch, () => this.regexMatch.serialize(writer));
        if (!w)
            return writer.getResultBuffer();
    }
    static deserialize(bytes: Uint8Array | pb_1.BinaryReader): BooleanExpression {
        const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new BooleanExpression();
        while (reader.nextField()) {
            if (reader.isEndGroup())
                break;
            switch (reader.getFieldNumber()) {
                case 1:
                    reader.readMessage(message.and, () => message.and = CompoundBooleanExpression.deserialize(reader));
                    break;
                case 2:
                    reader.readMessage(message.or, () => message.or = CompoundBooleanExpression.deserialize(reader));
                    break;
                case 3:
                    reader.readMessage(message.not, () => message.not = BooleanExpression.deserialize(reader));
                    break;
                case 4:
                    reader.readMessage(message.equal, () => message.equal = PairExpression.deserialize(reader));
                    break;
                case 5:
                    reader.readMessage(message.lessThan, () => message.lessThan = PairExpression.deserialize(reader));
                    break;
                case 6:
                    reader.readMessage(message.greaterThan, () => message.greaterThan = PairExpression.deserialize(reader));
                    break;
                case 7:
                    reader.readMessage(message.lessThanOrEqual, () => message.lessThanOrEqual = PairExpression.deserialize(reader));
                    break;
                case 8:
                    reader.readMessage(message.greaterThanOrEqual, () => message.greaterThanOrEqual = PairExpression.deserialize(reader));
                    break;
                case 9:
                    reader.readMessage(message.in, () => message.in = PairExpression.deserialize(reader));
                    break;
                case 10:
                    reader.readMessage(message.isNull, () => message.isNull = Expression.deserialize(reader));
                    break;
                case 11:
                    reader.readMessage(message.regexMatch, () => message.regexMatch = RegexMatchExpression.deserialize(reader));
                    break;
                default: reader.skipField();
            }
        }
        return message;
    }
    serializeBinary(): Uint8Array {
        return this.serialize();
    }
    static deserializeBinary(bytes: Uint8Array): BooleanExpression {
        return BooleanExpression.deserialize(bytes);
    }
}
