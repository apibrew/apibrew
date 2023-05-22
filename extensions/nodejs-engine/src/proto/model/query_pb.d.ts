// package: model
// file: model/query.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";

export class CompoundBooleanExpression extends jspb.Message { 
    clearExpressionsList(): void;
    getExpressionsList(): Array<BooleanExpression>;
    setExpressionsList(value: Array<BooleanExpression>): CompoundBooleanExpression;
    addExpressions(value?: BooleanExpression, index?: number): BooleanExpression;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CompoundBooleanExpression.AsObject;
    static toObject(includeInstance: boolean, msg: CompoundBooleanExpression): CompoundBooleanExpression.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CompoundBooleanExpression, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CompoundBooleanExpression;
    static deserializeBinaryFromReader(message: CompoundBooleanExpression, reader: jspb.BinaryReader): CompoundBooleanExpression;
}

export namespace CompoundBooleanExpression {
    export type AsObject = {
        expressionsList: Array<BooleanExpression.AsObject>,
    }
}

export class RefValue extends jspb.Message { 
    getNamespace(): string;
    setNamespace(value: string): RefValue;
    getResource(): string;
    setResource(value: string): RefValue;

    getPropertiesMap(): jspb.Map<string, google_protobuf_struct_pb.Value>;
    clearPropertiesMap(): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RefValue.AsObject;
    static toObject(includeInstance: boolean, msg: RefValue): RefValue.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RefValue, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RefValue;
    static deserializeBinaryFromReader(message: RefValue, reader: jspb.BinaryReader): RefValue;
}

export namespace RefValue {
    export type AsObject = {
        namespace: string,
        resource: string,

        propertiesMap: Array<[string, google_protobuf_struct_pb.Value.AsObject]>,
    }
}

export class Expression extends jspb.Message { 

    hasProperty(): boolean;
    clearProperty(): void;
    getProperty(): string;
    setProperty(value: string): Expression;

    hasValue(): boolean;
    clearValue(): void;
    getValue(): google_protobuf_struct_pb.Value | undefined;
    setValue(value?: google_protobuf_struct_pb.Value): Expression;

    hasRefvalue(): boolean;
    clearRefvalue(): void;
    getRefvalue(): RefValue | undefined;
    setRefvalue(value?: RefValue): Expression;

    getExpressionCase(): Expression.ExpressionCase;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Expression.AsObject;
    static toObject(includeInstance: boolean, msg: Expression): Expression.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Expression, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Expression;
    static deserializeBinaryFromReader(message: Expression, reader: jspb.BinaryReader): Expression;
}

export namespace Expression {
    export type AsObject = {
        property: string,
        value?: google_protobuf_struct_pb.Value.AsObject,
        refvalue?: RefValue.AsObject,
    }

    export enum ExpressionCase {
        EXPRESSION_NOT_SET = 0,
        PROPERTY = 1,
        VALUE = 3,
        REFVALUE = 4,
    }

}

export class PairExpression extends jspb.Message { 

    hasLeft(): boolean;
    clearLeft(): void;
    getLeft(): Expression | undefined;
    setLeft(value?: Expression): PairExpression;

    hasRight(): boolean;
    clearRight(): void;
    getRight(): Expression | undefined;
    setRight(value?: Expression): PairExpression;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PairExpression.AsObject;
    static toObject(includeInstance: boolean, msg: PairExpression): PairExpression.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PairExpression, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PairExpression;
    static deserializeBinaryFromReader(message: PairExpression, reader: jspb.BinaryReader): PairExpression;
}

export namespace PairExpression {
    export type AsObject = {
        left?: Expression.AsObject,
        right?: Expression.AsObject,
    }
}

export class RegexMatchExpression extends jspb.Message { 
    getPattern(): string;
    setPattern(value: string): RegexMatchExpression;

    hasExpression(): boolean;
    clearExpression(): void;
    getExpression(): Expression | undefined;
    setExpression(value?: Expression): RegexMatchExpression;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RegexMatchExpression.AsObject;
    static toObject(includeInstance: boolean, msg: RegexMatchExpression): RegexMatchExpression.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RegexMatchExpression, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RegexMatchExpression;
    static deserializeBinaryFromReader(message: RegexMatchExpression, reader: jspb.BinaryReader): RegexMatchExpression;
}

export namespace RegexMatchExpression {
    export type AsObject = {
        pattern: string,
        expression?: Expression.AsObject,
    }
}

export class BooleanExpression extends jspb.Message { 

    hasAnd(): boolean;
    clearAnd(): void;
    getAnd(): CompoundBooleanExpression | undefined;
    setAnd(value?: CompoundBooleanExpression): BooleanExpression;

    hasOr(): boolean;
    clearOr(): void;
    getOr(): CompoundBooleanExpression | undefined;
    setOr(value?: CompoundBooleanExpression): BooleanExpression;

    hasNot(): boolean;
    clearNot(): void;
    getNot(): BooleanExpression | undefined;
    setNot(value?: BooleanExpression): BooleanExpression;

    hasEqual(): boolean;
    clearEqual(): void;
    getEqual(): PairExpression | undefined;
    setEqual(value?: PairExpression): BooleanExpression;

    hasLessthan(): boolean;
    clearLessthan(): void;
    getLessthan(): PairExpression | undefined;
    setLessthan(value?: PairExpression): BooleanExpression;

    hasGreaterthan(): boolean;
    clearGreaterthan(): void;
    getGreaterthan(): PairExpression | undefined;
    setGreaterthan(value?: PairExpression): BooleanExpression;

    hasLessthanorequal(): boolean;
    clearLessthanorequal(): void;
    getLessthanorequal(): PairExpression | undefined;
    setLessthanorequal(value?: PairExpression): BooleanExpression;

    hasGreaterthanorequal(): boolean;
    clearGreaterthanorequal(): void;
    getGreaterthanorequal(): PairExpression | undefined;
    setGreaterthanorequal(value?: PairExpression): BooleanExpression;

    hasIn(): boolean;
    clearIn(): void;
    getIn(): PairExpression | undefined;
    setIn(value?: PairExpression): BooleanExpression;

    hasIsnull(): boolean;
    clearIsnull(): void;
    getIsnull(): Expression | undefined;
    setIsnull(value?: Expression): BooleanExpression;

    hasRegexmatch(): boolean;
    clearRegexmatch(): void;
    getRegexmatch(): RegexMatchExpression | undefined;
    setRegexmatch(value?: RegexMatchExpression): BooleanExpression;

    getExpressionCase(): BooleanExpression.ExpressionCase;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BooleanExpression.AsObject;
    static toObject(includeInstance: boolean, msg: BooleanExpression): BooleanExpression.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BooleanExpression, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BooleanExpression;
    static deserializeBinaryFromReader(message: BooleanExpression, reader: jspb.BinaryReader): BooleanExpression;
}

export namespace BooleanExpression {
    export type AsObject = {
        and?: CompoundBooleanExpression.AsObject,
        or?: CompoundBooleanExpression.AsObject,
        not?: BooleanExpression.AsObject,
        equal?: PairExpression.AsObject,
        lessthan?: PairExpression.AsObject,
        greaterthan?: PairExpression.AsObject,
        lessthanorequal?: PairExpression.AsObject,
        greaterthanorequal?: PairExpression.AsObject,
        pb_in?: PairExpression.AsObject,
        isnull?: Expression.AsObject,
        regexmatch?: RegexMatchExpression.AsObject,
    }

    export enum ExpressionCase {
        EXPRESSION_NOT_SET = 0,
        AND = 1,
        OR = 2,
        NOT = 3,
        EQUAL = 4,
        LESSTHAN = 5,
        GREATERTHAN = 6,
        LESSTHANOREQUAL = 7,
        GREATERTHANOREQUAL = 8,
        IN = 9,
        ISNULL = 10,
        REGEXMATCH = 11,
    }

}
