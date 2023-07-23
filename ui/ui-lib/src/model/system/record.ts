


export const RecordResource = {
    resource: "Record",
    namespace: "system",
};

// Sub Types

// Resource Type
export interface Record {
    id: string;
properties: object;
packedProperties?: object[];

}
// Resource and Property Names
export const RecordName = "Record";

export const RecordIdName = "Id";

export const RecordPropertiesName = "Properties";

export const RecordPackedPropertiesName = "PackedProperties";


