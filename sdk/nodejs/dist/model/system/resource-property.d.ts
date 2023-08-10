import { Resource } from "./resource";
export declare const ResourcePropertyResource: {
    resource: string;
    namespace: string;
};
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
export declare const ResourcePropertyName = "ResourceProperty";
export declare const ResourcePropertyIdName = "Id";
export declare const ResourcePropertyVersionName = "Version";
export declare const ResourcePropertyCreatedByName = "CreatedBy";
export declare const ResourcePropertyUpdatedByName = "UpdatedBy";
export declare const ResourcePropertyCreatedOnName = "CreatedOn";
export declare const ResourcePropertyUpdatedOnName = "UpdatedOn";
export declare const ResourcePropertyNameName = "Name";
export declare const ResourcePropertyTypeName = "Type";
export declare const ResourcePropertyTypeRefName = "TypeRef";
export declare const ResourcePropertyMappingName = "Mapping";
export declare const ResourcePropertyPrimaryName = "Primary";
export declare const ResourcePropertyRequiredName = "Required";
export declare const ResourcePropertyUniqueName = "Unique";
export declare const ResourcePropertyImmutableName = "Immutable";
export declare const ResourcePropertyLengthName = "Length";
export declare const ResourcePropertyResourceName = "Resource";
export declare const ResourcePropertyItemName = "Item";
export declare const ResourcePropertyPropertiesName = "Properties";
export declare const ResourcePropertyReferenceResourceName = "ReferenceResource";
export declare const ResourcePropertyReferenceCascadeName = "ReferenceCascade";
export declare const ResourcePropertyBackReferencePropertyName = "BackReferenceProperty";
export declare const ResourcePropertyDefaultValueName = "DefaultValue";
export declare const ResourcePropertyEnumValuesName = "EnumValues";
export declare const ResourcePropertyExampleValueName = "ExampleValue";
export declare const ResourcePropertyTitleName = "Title";
export declare const ResourcePropertyDescriptionName = "Description";
export declare const ResourcePropertyAnnotationsName = "Annotations";