// Constants

import {Type} from "../model/resource";

export const Enabled = "true";

// Resource flags
export const KeepHistory = "KeepHistory";
export const HistoryResource = "HistoryResource";
export const NormalizedResource = "NormalizedResource";
export const AutoCreated = "AutoCreated";
export const EnableAudit = "EnableAudit";
export const DisableVersion = "DisableVersion";
export const DisableBackup = "DisableBackup";

// Property params
export const SourceDef = "SourceDef";
export const SourceIdentity = "SourceIdentity";
export const SourceMatchKey = "SourceMatchKey";
export const Identity = "Identity";
export const SpecialProperty = "SpecialProperty";

// HCL
export const IsHclLabel = "IsHclLabel";
export const HclBlock = "HclBlock";

// Request annotations
export const IgnoreIfExists = "IgnoreIfExists";
export const CheckVersion = "CheckVersion";

// Security
export const AllowPublicAccess = "AllowPublicAccess";
export const AllowPublicReadAccess = "AllowPublicGetAccess";
export const AllowPublicCreateAccess = "AllowPublicCreateAccess";
export const AllowPublicUpdateAccess = "AllowPublicUpdateAccess";
export const AllowPublicDeleteAccess = "AllowPublicDeleteAccess";

// SQL
export const SQLType = "SQLType";
export const SQLUseTextType = "SQLUseTextType";
export const UseJoinTable = "UseJoinTable";

// Code generator
export const TypeName = "TypeName";
export const SelfContainedProperty = "SelfContainedProperty";
export const AllowEmptyPrimitive = "AllowEmptyPrimitive";
export const CommonType = "CommonType";

// REST API
export const RestApiDisabled = "RestApiDisabled";

// OpenAPI
export const OpenApiGroup = "OpenApiGroup";
export const OpenApiHide = "OpenApiHide";
export const OpenApiRestPath = "OpenApiRestPath";

// Service
export const ServiceKey = "ServiceKey";

// ExtensionId
export const ExtensionId = "ExtensionId";

// Bypass extensions
export const BypassExtensions = "BypassExtensions";

export const StudioSeparatePages = "StudioSeparatePages";

// Client allowed annotations
export const ClientAllowedAnnotations = {
    [BypassExtensions]: true,
};

export interface AnnotationDef {
    name: string
    type: Type
    defaultValue?: any
}

const annotations: AnnotationDef[] = [
    {name: Enabled, type: Type.BOOL},
    {name: KeepHistory, type: Type.BOOL},
    {name: HistoryResource, type: Type.BOOL},
    {name: NormalizedResource, type: Type.BOOL},
    {name: AutoCreated, type: Type.BOOL},
    {name: EnableAudit, type: Type.BOOL},
    {name: DisableVersion, type: Type.BOOL},
    {name: DisableBackup, type: Type.BOOL},
    {name: SourceDef, type: Type.BOOL},
    {name: SourceIdentity, type: Type.BOOL},
    {name: SourceMatchKey, type: Type.BOOL},
    {name: Identity, type: Type.BOOL},
    {name: SpecialProperty, type: Type.BOOL},
    {name: IsHclLabel, type: Type.BOOL},
    {name: HclBlock, type: Type.BOOL},
    {name: IgnoreIfExists, type: Type.BOOL},
    {name: CheckVersion, type: Type.BOOL},
    {name: AllowPublicAccess, type: Type.BOOL},
    {name: AllowPublicReadAccess, type: Type.BOOL},
    {name: AllowPublicCreateAccess, type: Type.BOOL},
    {name: AllowPublicUpdateAccess, type: Type.BOOL},
    {name: AllowPublicDeleteAccess, type: Type.BOOL},
    {name: SQLType, type: Type.BOOL},
    {name: SQLUseTextType, type: Type.BOOL},
    {name: UseJoinTable, type: Type.BOOL},
    {name: TypeName, type: Type.BOOL},
    {name: SelfContainedProperty, type: Type.BOOL},
    {name: AllowEmptyPrimitive, type: Type.BOOL},
    {name: CommonType, type: Type.BOOL},
    {name: RestApiDisabled, type: Type.BOOL},
    {name: OpenApiGroup, type: Type.BOOL},
    {name: OpenApiHide, type: Type.BOOL},
    {name: OpenApiRestPath, type: Type.BOOL},
    {name: ServiceKey, type: Type.BOOL},
    {name: ExtensionId, type: Type.BOOL},
    {name: BypassExtensions, type: Type.BOOL},
    {name: StudioSeparatePages, type: Type.BOOL}
];