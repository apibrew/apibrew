import { Namespace } from "./namespace";
import { DataSource } from "./data-source";
export declare const ResourceResource: {
    resource: string;
    namespace: string;
};
export interface Property {
    name: string;
    type: number;
    typeRef?: string;
    mapping: string;
    primary: boolean;
    required: boolean;
    unique: boolean;
    immutable: boolean;
    length: number;
    item?: Property;
    properties: Property[];
    reference?: Reference;
    defaultValue?: object;
    enumValues?: string[];
    exampleValue?: object;
    title?: string;
    description?: string;
    annotations?: object;
}
export interface SubType {
    name: string;
    properties: Property[];
}
export interface IndexProperty {
    name: string;
    order?: 'UNKNOWN' | 'ASC' | 'DESC';
}
export interface Index {
    properties?: IndexProperty[];
    indexType?: 'BTREE' | 'HASH';
    unique?: boolean;
    annotations?: object;
}
export interface Reference {
    resource?: Resource;
    cascade?: boolean;
    backReference?: string;
}
export interface Resource {
    id: string;
    version: number;
    createdBy: string;
    updatedBy?: string;
    createdOn: string;
    updatedOn?: string;
    name: string;
    namespace: Namespace;
    virtual: boolean;
    properties: Property[];
    indexes?: Index[];
    types?: SubType[];
    immutable: boolean;
    abstract: boolean;
    dataSource?: DataSource;
    entity?: string;
    catalog?: string;
    title?: string;
    description?: string;
    annotations?: object;
}
export declare const ResourceName = "Resource";
export declare const ResourceIdName = "Id";
export declare const ResourceVersionName = "Version";
export declare const ResourceCreatedByName = "CreatedBy";
export declare const ResourceUpdatedByName = "UpdatedBy";
export declare const ResourceCreatedOnName = "CreatedOn";
export declare const ResourceUpdatedOnName = "UpdatedOn";
export declare const ResourceNameName = "Name";
export declare const ResourceNamespaceName = "Namespace";
export declare const ResourceVirtualName = "Virtual";
export declare const ResourcePropertiesName = "Properties";
export declare const ResourceIndexesName = "Indexes";
export declare const ResourceTypesName = "Types";
export declare const ResourceImmutableName = "Immutable";
export declare const ResourceAbstractName = "Abstract";
export declare const ResourceDataSourceName = "DataSource";
export declare const ResourceEntityName = "Entity";
export declare const ResourceCatalogName = "Catalog";
export declare const ResourceTitleName = "Title";
export declare const ResourceDescriptionName = "Description";
export declare const ResourceAnnotationsName = "Annotations";
