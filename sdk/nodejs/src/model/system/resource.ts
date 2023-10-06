

import { Namespace } from "./namespace";

import { DataSource } from "./data-source";


export const ResourceResource = {
    resource: "Resource",
    namespace: "system",
};

// Sub Types

export interface Property {
     name?: string;
     type: 'BOOL' | 'STRING' | 'FLOAT32' | 'FLOAT64' | 'INT32' | 'INT64' | 'BYTES' | 'UUID' | 'DATE' | 'TIME' | 'TIMESTAMP' | 'OBJECT' | 'MAP' | 'LIST' | 'REFERENCE' | 'ENUM' | 'STRUCT';
     typeRef?: string;
     primary: boolean;
     required: boolean;
     unique: boolean;
     immutable: boolean;
     length: number;
     item?: Property;
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
     title?: string;
     description?: string;
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

// Resource Type
export interface Resource {
    id: string;
version: number;
createdBy?: string;
updatedBy?: string;
createdOn?: string;
updatedOn?: string;
name: string;
namespace: Namespace;
virtual: boolean;
properties: Property[];
indexes?: Index[];
types?: SubType[];
immutable: boolean;
abstract: boolean;
checkReferences: boolean;
dataSource?: DataSource;
entity?: string;
catalog?: string;
title?: string;
description?: string;
annotations?: object;

}
// Resource and Property Names
export const ResourceName = "Resource";

export const ResourceIdName = "Id";

export const ResourceVersionName = "Version";

export const ResourceCreatedByName = "CreatedBy";

export const ResourceUpdatedByName = "UpdatedBy";

export const ResourceCreatedOnName = "CreatedOn";

export const ResourceUpdatedOnName = "UpdatedOn";

export const ResourceNameName = "Name";

export const ResourceNamespaceName = "Namespace";

export const ResourceVirtualName = "Virtual";

export const ResourcePropertiesName = "Properties";

export const ResourceIndexesName = "Indexes";

export const ResourceTypesName = "Types";

export const ResourceImmutableName = "Immutable";

export const ResourceAbstractName = "Abstract";

export const ResourceCheckReferencesName = "CheckReferences";

export const ResourceDataSourceName = "DataSource";

export const ResourceEntityName = "Entity";

export const ResourceCatalogName = "Catalog";

export const ResourceTitleName = "Title";

export const ResourceDescriptionName = "Description";

export const ResourceAnnotationsName = "Annotations";


