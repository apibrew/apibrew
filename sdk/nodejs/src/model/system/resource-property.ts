import { Resource } from "./resource";

export const ResourcePropertyResource = {
    resource: "ResourceProperty",
    namespace: "system",
};

// Sub Types

// Resource Type
export interface ResourceProperty {
    id: string;
version: number;
createdBy: string;
updatedBy?: string;
createdOn: string;
updatedOn?: string;
name: string;
type: string;
typeRef?: string;
mapping: string;
primary: boolean;
required: boolean;
unique: boolean;
immutable: boolean;
length: number;
resource: Resource;
item?: object;
properties?: object;
referenceResource?: Resource;
referenceCascade?: boolean;
backReferenceProperty?: boolean;
defaultValue?: object;
enumValues?: object;
exampleValue?: object;
title?: string;
description?: string;
annotations?: object;

}
// Resource and Property Names
export const ResourcePropertyName = "ResourceProperty";

export const ResourcePropertyIdName = "Id";

export const ResourcePropertyVersionName = "Version";

export const ResourcePropertyCreatedByName = "CreatedBy";

export const ResourcePropertyUpdatedByName = "UpdatedBy";

export const ResourcePropertyCreatedOnName = "CreatedOn";

export const ResourcePropertyUpdatedOnName = "UpdatedOn";

export const ResourcePropertyNameName = "Name";

export const ResourcePropertyTypeName = "Type";

export const ResourcePropertyTypeRefName = "TypeRef";

export const ResourcePropertyMappingName = "Mapping";

export const ResourcePropertyPrimaryName = "Primary";

export const ResourcePropertyRequiredName = "Required";

export const ResourcePropertyUniqueName = "Unique";

export const ResourcePropertyImmutableName = "Immutable";

export const ResourcePropertyLengthName = "Length";

export const ResourcePropertyResourceName = "Resource";

export const ResourcePropertyItemName = "Item";

export const ResourcePropertyPropertiesName = "Properties";

export const ResourcePropertyReferenceResourceName = "ReferenceResource";

export const ResourcePropertyReferenceCascadeName = "ReferenceCascade";

export const ResourcePropertyBackReferencePropertyName = "BackReferenceProperty";

export const ResourcePropertyDefaultValueName = "DefaultValue";

export const ResourcePropertyEnumValuesName = "EnumValues";

export const ResourcePropertyExampleValueName = "ExampleValue";

export const ResourcePropertyTitleName = "Title";

export const ResourcePropertyDescriptionName = "Description";

export const ResourcePropertyAnnotationsName = "Annotations";


