export declare const RecordResource: {
    resource: string;
    namespace: string;
};
export interface Record {
    id: string;
    properties: object;
    packedProperties?: object[];
}
export declare const RecordName = "Record";
export declare const RecordIdName = "Id";
export declare const RecordPropertiesName = "Properties";
export declare const RecordPackedPropertiesName = "PackedProperties";
