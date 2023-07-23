

import { Namespace } from "./namespace";

import { DataSource } from "./data-source";


export const ResourceResource = {
    resource: "Resource",
    namespace: "system",
};

// Sub Types

// Resource Type
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
types?: object;
immutable: boolean;
abstract: boolean;
dataSource?: DataSource;
entity?: string;
catalog?: string;
annotations?: object;
indexes?: object;
title?: string;
description?: string;

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

export const ResourceTypesName = "Types";

export const ResourceImmutableName = "Immutable";

export const ResourceAbstractName = "Abstract";

export const ResourceDataSourceName = "DataSource";

export const ResourceEntityName = "Entity";

export const ResourceCatalogName = "Catalog";

export const ResourceAnnotationsName = "Annotations";

export const ResourceIndexesName = "Indexes";

export const ResourceTitleName = "Title";

export const ResourceDescriptionName = "Description";


