# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [model/audit.proto](#model_audit-proto)
    - [AuditData](#model-AuditData)
  
- [model/common.proto](#model_common-proto)
    - [MapAnyWrap](#model-MapAnyWrap)
    - [MapAnyWrap.ContentEntry](#model-MapAnyWrap-ContentEntry)
  
- [model/record.proto](#model_record-proto)
    - [Record](#model-Record)
    - [Record.PropertiesEntry](#model-Record-PropertiesEntry)
  
- [model/query.proto](#model_query-proto)
    - [BooleanExpression](#model-BooleanExpression)
    - [CompoundBooleanExpression](#model-CompoundBooleanExpression)
    - [Expression](#model-Expression)
    - [PairExpression](#model-PairExpression)
    - [RefValue](#model-RefValue)
    - [RefValue.PropertiesEntry](#model-RefValue-PropertiesEntry)
    - [RegexMatchExpression](#model-RegexMatchExpression)
  
- [model/annotations.proto](#model_annotations-proto)
    - [File-level Extensions](#model_annotations-proto-extensions)
    - [File-level Extensions](#model_annotations-proto-extensions)
    - [File-level Extensions](#model_annotations-proto-extensions)
  
- [model/security.proto](#model_security-proto)
    - [SecurityConstraint](#model-SecurityConstraint)
    - [SecurityContext](#model-SecurityContext)
  
    - [OperationType](#model-OperationType)
    - [PermitType](#model-PermitType)
  
- [model/resource.proto](#model_resource-proto)
    - [Reference](#model-Reference)
    - [Resource](#model-Resource)
    - [Resource.AnnotationsEntry](#model-Resource-AnnotationsEntry)
    - [ResourceIndex](#model-ResourceIndex)
    - [ResourceIndex.AnnotationsEntry](#model-ResourceIndex-AnnotationsEntry)
    - [ResourceIndexProperty](#model-ResourceIndexProperty)
    - [ResourceProperty](#model-ResourceProperty)
    - [ResourceProperty.AnnotationsEntry](#model-ResourceProperty-AnnotationsEntry)
    - [ResourceSourceConfig](#model-ResourceSourceConfig)
  
    - [Order](#model-Order)
    - [ResourceIndexType](#model-ResourceIndexType)
    - [ResourceProperty.Type](#model-ResourceProperty-Type)
  
- [model/error.proto](#model_error-proto)
    - [Error](#model-Error)
    - [ErrorField](#model-ErrorField)
  
    - [ErrorCode](#model-ErrorCode)
  
- [ext/function.proto](#ext_function-proto)
    - [FunctionCallRequest](#ext-FunctionCallRequest)
    - [FunctionCallRequest.RequestEntry](#ext-FunctionCallRequest-RequestEntry)
    - [FunctionCallResponse](#ext-FunctionCallResponse)
    - [FunctionCallResponse.ResponseEntry](#ext-FunctionCallResponse-ResponseEntry)
  
    - [Function](#ext-Function)
  
- [model/batch.proto](#model_batch-proto)
    - [Batch](#model-Batch)
    - [BatchHeader](#model-BatchHeader)
    - [BatchHeader.AnnotationsEntry](#model-BatchHeader-AnnotationsEntry)
    - [BatchRecordsPart](#model-BatchRecordsPart)
  
    - [BatchHeader.BatchMode](#model-BatchHeader-BatchMode)
  
- [model/data-source.proto](#model_data-source-proto)
    - [DataSource](#model-DataSource)
    - [DataSourceCatalog](#model-DataSourceCatalog)
    - [DataSourceEntity](#model-DataSourceEntity)
    - [MongoParams](#model-MongoParams)
    - [MysqlParams](#model-MysqlParams)
    - [PostgresqlParams](#model-PostgresqlParams)
    - [RedisParams](#model-RedisParams)
    - [VirtualParams](#model-VirtualParams)
  
    - [DataSourceBackendType](#model-DataSourceBackendType)
    - [VirtualParams.Mode](#model-VirtualParams-Mode)
  
- [model/external.proto](#model_external-proto)
    - [ExternalCall](#model-ExternalCall)
    - [FunctionCall](#model-FunctionCall)
    - [HttpCall](#model-HttpCall)
  
- [model/extension.proto](#model_extension-proto)
    - [Extension](#model-Extension)
    - [Extension.After](#model-Extension-After)
    - [Extension.Before](#model-Extension-Before)
    - [Extension.Instead](#model-Extension-Instead)
  
- [model/hcl.proto](#model_hcl-proto)
- [model/user.proto](#model_user-proto)
    - [User](#model-User)
  
- [model/namespace.proto](#model_namespace-proto)
    - [Namespace](#model-Namespace)
  
- [model/init.proto](#model_init-proto)
    - [AppConfig](#model-AppConfig)
    - [InitData](#model-InitData)
  
- [model/resource-migration.proto](#model_resource-migration-proto)
    - [ResourceMigrationCreateIndex](#model-ResourceMigrationCreateIndex)
    - [ResourceMigrationCreateProperty](#model-ResourceMigrationCreateProperty)
    - [ResourceMigrationCreateResource](#model-ResourceMigrationCreateResource)
    - [ResourceMigrationDeleteIndex](#model-ResourceMigrationDeleteIndex)
    - [ResourceMigrationDeleteProperty](#model-ResourceMigrationDeleteProperty)
    - [ResourceMigrationDeleteResource](#model-ResourceMigrationDeleteResource)
    - [ResourceMigrationPlan](#model-ResourceMigrationPlan)
    - [ResourceMigrationStep](#model-ResourceMigrationStep)
    - [ResourceMigrationUpdateProperty](#model-ResourceMigrationUpdateProperty)
    - [ResourceMigrationUpdateResource](#model-ResourceMigrationUpdateResource)
  
- [model/token.proto](#model_token-proto)
    - [Token](#model-Token)
  
    - [TokenTerm](#model-TokenTerm)
  
- [model/watch.proto](#model_watch-proto)
    - [WatchMessage](#model-WatchMessage)
  
    - [EventType](#model-EventType)
  
- [stub/authentication.proto](#stub_authentication-proto)
    - [AuthenticationRequest](#stub-AuthenticationRequest)
    - [AuthenticationResponse](#stub-AuthenticationResponse)
    - [RenewTokenRequest](#stub-RenewTokenRequest)
    - [RenewTokenResponse](#stub-RenewTokenResponse)
  
    - [Authentication](#stub-Authentication)
  
- [stub/data-source.proto](#stub_data-source-proto)
    - [CreateDataSourceRequest](#stub-CreateDataSourceRequest)
    - [CreateDataSourceResponse](#stub-CreateDataSourceResponse)
    - [DeleteDataSourceRequest](#stub-DeleteDataSourceRequest)
    - [DeleteDataSourceResponse](#stub-DeleteDataSourceResponse)
    - [GetDataSourceRequest](#stub-GetDataSourceRequest)
    - [GetDataSourceResponse](#stub-GetDataSourceResponse)
    - [ListDataSourceRequest](#stub-ListDataSourceRequest)
    - [ListDataSourceResponse](#stub-ListDataSourceResponse)
    - [ListEntitiesRequest](#stub-ListEntitiesRequest)
    - [ListEntitiesResponse](#stub-ListEntitiesResponse)
    - [PrepareResourceFromEntityRequest](#stub-PrepareResourceFromEntityRequest)
    - [PrepareResourceFromEntityResponse](#stub-PrepareResourceFromEntityResponse)
    - [StatusRequest](#stub-StatusRequest)
    - [StatusResponse](#stub-StatusResponse)
    - [UpdateDataSourceRequest](#stub-UpdateDataSourceRequest)
    - [UpdateDataSourceResponse](#stub-UpdateDataSourceResponse)
  
    - [DataSource](#stub-DataSource)
  
- [stub/extension.proto](#stub_extension-proto)
    - [CreateExtensionRequest](#stub-CreateExtensionRequest)
    - [CreateExtensionResponse](#stub-CreateExtensionResponse)
    - [DeleteExtensionRequest](#stub-DeleteExtensionRequest)
    - [DeleteExtensionResponse](#stub-DeleteExtensionResponse)
    - [GetExtensionRequest](#stub-GetExtensionRequest)
    - [GetExtensionResponse](#stub-GetExtensionResponse)
    - [ListExtensionRequest](#stub-ListExtensionRequest)
    - [ListExtensionResponse](#stub-ListExtensionResponse)
    - [UpdateExtensionRequest](#stub-UpdateExtensionRequest)
    - [UpdateExtensionResponse](#stub-UpdateExtensionResponse)
  
    - [Extension](#stub-Extension)
  
- [stub/generic.proto](#stub_generic-proto)
    - [CreateRequest](#stub-CreateRequest)
    - [CreateRequest.AnnotationsEntry](#stub-CreateRequest-AnnotationsEntry)
    - [CreateResponse](#stub-CreateResponse)
    - [DeleteRequest](#stub-DeleteRequest)
    - [DeleteRequest.AnnotationsEntry](#stub-DeleteRequest-AnnotationsEntry)
    - [DeleteResponse](#stub-DeleteResponse)
    - [GetRequest](#stub-GetRequest)
    - [GetRequest.AnnotationsEntry](#stub-GetRequest-AnnotationsEntry)
    - [GetResponse](#stub-GetResponse)
    - [ListRequest](#stub-ListRequest)
    - [ListRequest.AnnotationsEntry](#stub-ListRequest-AnnotationsEntry)
    - [ListRequest.FiltersEntry](#stub-ListRequest-FiltersEntry)
    - [ListResponse](#stub-ListResponse)
    - [SearchRequest](#stub-SearchRequest)
    - [SearchRequest.AnnotationsEntry](#stub-SearchRequest-AnnotationsEntry)
    - [SearchResponse](#stub-SearchResponse)
    - [UpdateMultiRequest](#stub-UpdateMultiRequest)
    - [UpdateMultiRequest.AnnotationsEntry](#stub-UpdateMultiRequest-AnnotationsEntry)
    - [UpdateMultiRequest.PropertiesEntry](#stub-UpdateMultiRequest-PropertiesEntry)
    - [UpdateMultiResponse](#stub-UpdateMultiResponse)
    - [UpdateRequest](#stub-UpdateRequest)
    - [UpdateRequest.AnnotationsEntry](#stub-UpdateRequest-AnnotationsEntry)
    - [UpdateResponse](#stub-UpdateResponse)
  
    - [Generic](#stub-Generic)
  
- [stub/namespace.proto](#stub_namespace-proto)
    - [CreateNamespaceRequest](#stub-CreateNamespaceRequest)
    - [CreateNamespaceResponse](#stub-CreateNamespaceResponse)
    - [DeleteNamespaceRequest](#stub-DeleteNamespaceRequest)
    - [DeleteNamespaceResponse](#stub-DeleteNamespaceResponse)
    - [GetNamespaceRequest](#stub-GetNamespaceRequest)
    - [GetNamespaceResponse](#stub-GetNamespaceResponse)
    - [ListNamespaceRequest](#stub-ListNamespaceRequest)
    - [ListNamespaceResponse](#stub-ListNamespaceResponse)
    - [UpdateNamespaceRequest](#stub-UpdateNamespaceRequest)
    - [UpdateNamespaceResponse](#stub-UpdateNamespaceResponse)
  
    - [Namespace](#stub-Namespace)
  
- [stub/openapi.proto](#stub_openapi-proto)
- [stub/record.proto](#stub_record-proto)
    - [ApplyRecordRequest](#stub-ApplyRecordRequest)
    - [ApplyRecordRequest.AnnotationsEntry](#stub-ApplyRecordRequest-AnnotationsEntry)
    - [ApplyRecordResponse](#stub-ApplyRecordResponse)
    - [CreateRecordRequest](#stub-CreateRecordRequest)
    - [CreateRecordRequest.AnnotationsEntry](#stub-CreateRecordRequest-AnnotationsEntry)
    - [CreateRecordResponse](#stub-CreateRecordResponse)
    - [DeleteRecordRequest](#stub-DeleteRecordRequest)
    - [DeleteRecordRequest.AnnotationsEntry](#stub-DeleteRecordRequest-AnnotationsEntry)
    - [DeleteRecordResponse](#stub-DeleteRecordResponse)
    - [GetRecordRequest](#stub-GetRecordRequest)
    - [GetRecordRequest.AnnotationsEntry](#stub-GetRecordRequest-AnnotationsEntry)
    - [GetRecordResponse](#stub-GetRecordResponse)
    - [ListRecordRequest](#stub-ListRecordRequest)
    - [ListRecordRequest.AnnotationsEntry](#stub-ListRecordRequest-AnnotationsEntry)
    - [ListRecordRequest.FiltersEntry](#stub-ListRecordRequest-FiltersEntry)
    - [ListRecordResponse](#stub-ListRecordResponse)
    - [ReadStreamRequest](#stub-ReadStreamRequest)
    - [ReadStreamRequest.AnnotationsEntry](#stub-ReadStreamRequest-AnnotationsEntry)
    - [SearchRecordRequest](#stub-SearchRecordRequest)
    - [SearchRecordRequest.AnnotationsEntry](#stub-SearchRecordRequest-AnnotationsEntry)
    - [SearchRecordResponse](#stub-SearchRecordResponse)
    - [UpdateMultiRecordRequest](#stub-UpdateMultiRecordRequest)
    - [UpdateMultiRecordRequest.AnnotationsEntry](#stub-UpdateMultiRecordRequest-AnnotationsEntry)
    - [UpdateMultiRecordRequest.PropertiesEntry](#stub-UpdateMultiRecordRequest-PropertiesEntry)
    - [UpdateMultiRecordResponse](#stub-UpdateMultiRecordResponse)
    - [UpdateRecordRequest](#stub-UpdateRecordRequest)
    - [UpdateRecordRequest.AnnotationsEntry](#stub-UpdateRecordRequest-AnnotationsEntry)
    - [UpdateRecordResponse](#stub-UpdateRecordResponse)
    - [WriteStreamResponse](#stub-WriteStreamResponse)
  
    - [Record](#stub-Record)
  
- [stub/resource.proto](#stub_resource-proto)
    - [CreateResourceRequest](#stub-CreateResourceRequest)
    - [CreateResourceRequest.AnnotationsEntry](#stub-CreateResourceRequest-AnnotationsEntry)
    - [CreateResourceResponse](#stub-CreateResourceResponse)
    - [DeleteResourceRequest](#stub-DeleteResourceRequest)
    - [DeleteResourceRequest.AnnotationsEntry](#stub-DeleteResourceRequest-AnnotationsEntry)
    - [DeleteResourceResponse](#stub-DeleteResourceResponse)
    - [GetResourceByNameRequest](#stub-GetResourceByNameRequest)
    - [GetResourceByNameRequest.AnnotationsEntry](#stub-GetResourceByNameRequest-AnnotationsEntry)
    - [GetResourceByNameResponse](#stub-GetResourceByNameResponse)
    - [GetResourceRequest](#stub-GetResourceRequest)
    - [GetResourceRequest.AnnotationsEntry](#stub-GetResourceRequest-AnnotationsEntry)
    - [GetResourceResponse](#stub-GetResourceResponse)
    - [GetSystemResourceRequest](#stub-GetSystemResourceRequest)
    - [GetSystemResourceRequest.AnnotationsEntry](#stub-GetSystemResourceRequest-AnnotationsEntry)
    - [GetSystemResourceResponse](#stub-GetSystemResourceResponse)
    - [ListResourceRequest](#stub-ListResourceRequest)
    - [ListResourceRequest.AnnotationsEntry](#stub-ListResourceRequest-AnnotationsEntry)
    - [ListResourceResponse](#stub-ListResourceResponse)
    - [PrepareResourceMigrationPlanRequest](#stub-PrepareResourceMigrationPlanRequest)
    - [PrepareResourceMigrationPlanRequest.AnnotationsEntry](#stub-PrepareResourceMigrationPlanRequest-AnnotationsEntry)
    - [PrepareResourceMigrationPlanResponse](#stub-PrepareResourceMigrationPlanResponse)
    - [UpdateResourceRequest](#stub-UpdateResourceRequest)
    - [UpdateResourceRequest.AnnotationsEntry](#stub-UpdateResourceRequest-AnnotationsEntry)
    - [UpdateResourceResponse](#stub-UpdateResourceResponse)
  
    - [Resource](#stub-Resource)
  
- [stub/rest/record.proto](#stub_rest_record-proto)
    - [ApplyRecordRequest](#rest-ApplyRecordRequest)
    - [ApplyRecordRequest.AnnotationsEntry](#rest-ApplyRecordRequest-AnnotationsEntry)
    - [ApplyRecordRequest.PropertiesEntry](#rest-ApplyRecordRequest-PropertiesEntry)
    - [ApplyRecordResponse](#rest-ApplyRecordResponse)
    - [ApplyRecordResponse.PropertiesEntry](#rest-ApplyRecordResponse-PropertiesEntry)
    - [CreateRecordRequest](#rest-CreateRecordRequest)
    - [CreateRecordRequest.AnnotationsEntry](#rest-CreateRecordRequest-AnnotationsEntry)
    - [CreateRecordRequest.PropertiesEntry](#rest-CreateRecordRequest-PropertiesEntry)
    - [CreateRecordResponse](#rest-CreateRecordResponse)
    - [CreateRecordResponse.PropertiesEntry](#rest-CreateRecordResponse-PropertiesEntry)
    - [DeleteRecordRequest](#rest-DeleteRecordRequest)
    - [DeleteRecordRequest.AnnotationsEntry](#rest-DeleteRecordRequest-AnnotationsEntry)
    - [DeleteRecordResponse](#rest-DeleteRecordResponse)
    - [UpdateRecordRequest](#rest-UpdateRecordRequest)
    - [UpdateRecordRequest.AnnotationsEntry](#rest-UpdateRecordRequest-AnnotationsEntry)
    - [UpdateRecordRequest.PropertiesEntry](#rest-UpdateRecordRequest-PropertiesEntry)
    - [UpdateRecordResponse](#rest-UpdateRecordResponse)
    - [UpdateRecordResponse.PropertiesEntry](#rest-UpdateRecordResponse-PropertiesEntry)
  
    - [Record](#rest-Record)
  
- [stub/user.proto](#stub_user-proto)
    - [CreateUserRequest](#stub-CreateUserRequest)
    - [CreateUserResponse](#stub-CreateUserResponse)
    - [DeleteUserRequest](#stub-DeleteUserRequest)
    - [DeleteUserResponse](#stub-DeleteUserResponse)
    - [GetUserRequest](#stub-GetUserRequest)
    - [GetUserResponse](#stub-GetUserResponse)
    - [ListUserRequest](#stub-ListUserRequest)
    - [ListUserResponse](#stub-ListUserResponse)
    - [UpdateUserRequest](#stub-UpdateUserRequest)
    - [UpdateUserResponse](#stub-UpdateUserResponse)
  
    - [User](#stub-User)
  
- [stub/watch.proto](#stub_watch-proto)
    - [WatchRequest](#stub-WatchRequest)
  
    - [Watch](#stub-Watch)
  
- [Scalar Value Types](#scalar-value-types)



<a name="model_audit-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## model/audit.proto



<a name="model-AuditData"></a>

### AuditData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| created_on | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_on | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| created_by | [string](#string) |  |  |
| updated_by | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="model_common-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## model/common.proto



<a name="model-MapAnyWrap"></a>

### MapAnyWrap



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| content | [MapAnyWrap.ContentEntry](#model-MapAnyWrap-ContentEntry) | repeated |  |






<a name="model-MapAnyWrap-ContentEntry"></a>

### MapAnyWrap.ContentEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Any](#google-protobuf-Any) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="model_record-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## model/record.proto



<a name="model-Record"></a>

### Record



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | id; read only |
| properties | [Record.PropertiesEntry](#model-Record-PropertiesEntry) | repeated |  |
| propertiesPacked | [google.protobuf.Value](#google-protobuf-Value) | repeated |  |






<a name="model-Record-PropertiesEntry"></a>

### Record.PropertiesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Value](#google-protobuf-Value) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="model_query-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## model/query.proto



<a name="model-BooleanExpression"></a>

### BooleanExpression



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| and | [CompoundBooleanExpression](#model-CompoundBooleanExpression) |  | logical expressions |
| or | [CompoundBooleanExpression](#model-CompoundBooleanExpression) |  |  |
| not | [BooleanExpression](#model-BooleanExpression) |  |  |
| equal | [PairExpression](#model-PairExpression) |  | basic comparison |
| lessThan | [PairExpression](#model-PairExpression) |  |  |
| greaterThan | [PairExpression](#model-PairExpression) |  |  |
| lessThanOrEqual | [PairExpression](#model-PairExpression) |  |  |
| greaterThanOrEqual | [PairExpression](#model-PairExpression) |  |  |
| in | [PairExpression](#model-PairExpression) |  |  |
| isNull | [Expression](#model-Expression) |  |  |
| regexMatch | [RegexMatchExpression](#model-RegexMatchExpression) |  | other |






<a name="model-CompoundBooleanExpression"></a>

### CompoundBooleanExpression



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| expressions | [BooleanExpression](#model-BooleanExpression) | repeated |  |






<a name="model-Expression"></a>

### Expression



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| property | [string](#string) |  |  |
| value | [google.protobuf.Value](#google-protobuf-Value) |  |  |
| refValue | [RefValue](#model-RefValue) |  |  |






<a name="model-PairExpression"></a>

### PairExpression



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| left | [Expression](#model-Expression) |  |  |
| right | [Expression](#model-Expression) |  |  |






<a name="model-RefValue"></a>

### RefValue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| properties | [RefValue.PropertiesEntry](#model-RefValue-PropertiesEntry) | repeated |  |






<a name="model-RefValue-PropertiesEntry"></a>

### RefValue.PropertiesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Value](#google-protobuf-Value) |  |  |






<a name="model-RegexMatchExpression"></a>

### RegexMatchExpression



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pattern | [string](#string) |  |  |
| expression | [Expression](#model-Expression) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="model_annotations-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## model/annotations.proto


 <!-- end messages -->

 <!-- end enums -->


<a name="model_annotations-proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| hcl_block | string | .google.protobuf.FieldOptions | 1144 |  |
| hcl_ignore | bool | .google.protobuf.FieldOptions | 1146 |  |
| hcl_label | string | .google.protobuf.FieldOptions | 1145 |  |

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="model_security-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## model/security.proto



<a name="model-SecurityConstraint"></a>

### SecurityConstraint
SecurityConstraint is a rule


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace | [string](#string) |  | namespace name where it will be applied |
| resource | [string](#string) |  | resource name where it will be applied |
| property | [string](#string) |  | property name where it will be applied |
| before | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | before it is valid |
| after | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | after it is valid |
| principal | [string](#string) |  | username which it is applied to |
| recordIds | [string](#string) | repeated | list of record ids which it is applied to |
| operation | [OperationType](#model-OperationType) |  | operation name which it is applied to |
| permit | [PermitType](#model-PermitType) |  | permission |






<a name="model-SecurityContext"></a>

### SecurityContext



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| constraints | [SecurityConstraint](#model-SecurityConstraint) | repeated |  |





 <!-- end messages -->


<a name="model-OperationType"></a>

### OperationType


| Name | Number | Description |
| ---- | ------ | ----------- |
| OPERATION_TYPE_READ | 0 |  |
| OPERATION_TYPE_CREATE | 1 |  |
| OPERATION_TYPE_UPDATE | 2 |  |
| OPERATION_TYPE_DELETE | 3 |  |
| FULL | 4 |  |



<a name="model-PermitType"></a>

### PermitType


| Name | Number | Description |
| ---- | ------ | ----------- |
| PERMIT_TYPE_ALLOW | 0 |  |
| PERMIT_TYPE_REJECT | 1 |  |
| PERMIT_TYPE_UNKNOWN | 2 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="model_resource-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## model/resource.proto



<a name="model-Reference"></a>

### Reference



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| referencedResource | [string](#string) |  | referenced resource name |
| cascade | [bool](#bool) |  | if cascade is true, delete/update operations will be cascaded to back referenced resources |






<a name="model-Resource"></a>

### Resource



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | unique resource id; read only |
| name | [string](#string) |  | unique resource name, it is unique per namespace |
| namespace | [string](#string) |  | each resource is kept inside a namespace. One namespace can have multiple resources |
| sourceConfig | [ResourceSourceConfig](#model-ResourceSourceConfig) |  |  |
| properties | [ResourceProperty](#model-ResourceProperty) | repeated | list of properties of resource. This properties will be used by records of resource. Properties is columns on sql databases. For schemaless data structures properties is only managed by Data handler itself |
| indexes | [ResourceIndex](#model-ResourceIndex) | repeated | list of resource indexes. Its implementation is depending on data source backend and may not be supported by some backends. |
| securityContext | [SecurityContext](#model-SecurityContext) |  | security context is to apply ACL to resource property |
| virtual | [bool](#bool) |  | If virtual is true. Operations will not phisically affect datasource/backend. Virtual resources is for extension purposes. Their behaviors can be extended and altered. It can also be used to integrate 3rd party systems. |
| immutable | [bool](#bool) |  | if true, delete and update will not be allowed on this resource |
| abstract | [bool](#bool) |  | if abstract, resource is only available to internal and extension side operations |
| title | [string](#string) | optional |  |
| description | [string](#string) | optional |  |
| auditData | [AuditData](#model-AuditData) |  |  |
| version | [uint32](#uint32) |  |  |
| annotations | [Resource.AnnotationsEntry](#model-Resource-AnnotationsEntry) | repeated |  |






<a name="model-Resource-AnnotationsEntry"></a>

### Resource.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="model-ResourceIndex"></a>

### ResourceIndex



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| properties | [ResourceIndexProperty](#model-ResourceIndexProperty) | repeated | list of properties inside single index. Normally you will need only single property. Multi property will be needed for multi property indexes(for complex indexes) |
| indexType | [ResourceIndexType](#model-ResourceIndexType) |  | Index type(BTREE, HASH) |
| unique | [bool](#bool) |  | if true index will be unique index |
| annotations | [ResourceIndex.AnnotationsEntry](#model-ResourceIndex-AnnotationsEntry) | repeated |  |






<a name="model-ResourceIndex-AnnotationsEntry"></a>

### ResourceIndex.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="model-ResourceIndexProperty"></a>

### ResourceIndexProperty



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| order | [Order](#model-Order) |  |  |






<a name="model-ResourceProperty"></a>

### ResourceProperty
Resource properties is used to describe its schema. Each resource property is corresponding to a field in a record
Data handler is responsible to validate data according to property types. For example, when you call create record and
if you send 123.45 for int64


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) | optional |  |
| name | [string](#string) |  | property name |
| type | [ResourceProperty.Type](#model-ResourceProperty-Type) |  | type of property - see [property-types](#property-types) section |
| mapping | [string](#string) |  | mapping is like a column name, it is binding name to entity. For abstraction purposes property name is not used while communicating to resource backend. Instead mapping is used as a key of property |
| required | [bool](#bool) |  | this is to mark property as required |
| primary | [bool](#bool) |  | this is to mark property as primary. Primary properties is like a part of primary key. Primary property(s) is used in to identify record. |
| length | [uint32](#uint32) |  | length property is only valid and required for String typed properties |
| unique | [bool](#bool) |  |  |
| immutable | [bool](#bool) |  | immutable is to mark property as immutable. If marked, updates on this field on records will be discarded |
| securityContext | [SecurityContext](#model-SecurityContext) | optional | security context is to apply ACL to resource property |
| defaultValue | [google.protobuf.Value](#google-protobuf-Value) | optional | defaultValue is default value. |
| exampleValue | [google.protobuf.Value](#google-protobuf-Value) | optional | exampleValue is example value. It is an informative column |
| enumValues | [google.protobuf.Value](#google-protobuf-Value) | repeated | enumValues is used if property type is an enum |
| reference | [Reference](#model-Reference) | optional | reference property is only valid and required for Reference types. |
| subType | [ResourceProperty.Type](#model-ResourceProperty-Type) | optional | subType is used for complex types(list, map). For list, subType is element type. For map, it is value type(key type is always string) |
| title | [string](#string) | optional | It is an informative column |
| description | [string](#string) | optional | It is an informative column |
| annotations | [ResourceProperty.AnnotationsEntry](#model-ResourceProperty-AnnotationsEntry) | repeated |  |






<a name="model-ResourceProperty-AnnotationsEntry"></a>

### ResourceProperty.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="model-ResourceSourceConfig"></a>

### ResourceSourceConfig
source config is to configure resource and bind it to data-source and an entity inside data source.
An entity is like a table on sql databases or collection on mongodb etc.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataSource | [string](#string) |  | data source name: where resource structure and its data will be physically exists. Data source name is required if resource is not virtual |
| catalog | [string](#string) |  | catalog is like a folder/schema/database. It is changing from backend to backend. Basically it is for grouping entities |
| entity | [string](#string) |  | entity name an item on datasource backend where resource will be bound. For sql databases it is table name, for mongo it is collection name, etc. |





 <!-- end messages -->


<a name="model-Order"></a>

### Order


| Name | Number | Description |
| ---- | ------ | ----------- |
| ORDER_UNKNOWN | 0 |  |
| ORDER_ASC | 1 |  |
| ORDER_DESC | 2 |  |



<a name="model-ResourceIndexType"></a>

### ResourceIndexType


| Name | Number | Description |
| ---- | ------ | ----------- |
| BTREE | 0 |  |
| HASH | 1 |  |



<a name="model-ResourceProperty-Type"></a>

### ResourceProperty.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| BOOL | 0 |  |
| STRING | 1 |  |
| FLOAT32 | 2 |  |
| FLOAT64 | 3 |  |
| INT32 | 4 |  |
| INT64 | 5 |  |
| BYTES | 6 |  |
| UUID | 8 |  |
| DATE | 9 |  |
| TIME | 10 |  |
| TIMESTAMP | 11 |  |
| OBJECT | 12 |  |
| MAP | 13 |  |
| LIST | 14 |  |
| REFERENCE | 15 |  |
| ENUM | 16 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="model_error-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## model/error.proto



<a name="model-Error"></a>

### Error



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [ErrorCode](#model-ErrorCode) |  |  |
| message | [string](#string) |  |  |
| fields | [ErrorField](#model-ErrorField) | repeated |  |






<a name="model-ErrorField"></a>

### ErrorField



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| recordId | [string](#string) |  |  |
| property | [string](#string) |  |  |
| message | [string](#string) |  |  |
| value | [google.protobuf.Value](#google-protobuf-Value) |  |  |





 <!-- end messages -->


<a name="model-ErrorCode"></a>

### ErrorCode


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN_ERROR | 0 |  |
| RECORD_NOT_FOUND | 1 |  |
| UNABLE_TO_LOCATE_PRIMARY_KEY | 2 |  |
| INTERNAL_ERROR | 3 |  |
| PROPERTY_NOT_FOUND | 4 |  |
| RECORD_VALIDATION_ERROR | 5 |  |
| RESOURCE_VALIDATION_ERROR | 13 |  |
| AUTHENTICATION_FAILED | 6 |  |
| ALREADY_EXISTS | 7 |  |
| ACCESS_DENIED | 8 |  |
| BACKEND_ERROR | 9 |  |
| UNIQUE_VIOLATION | 10 |  |
| REFERENCE_VIOLATION | 11 |  |
| RESOURCE_NOT_FOUND | 12 |  |
| UNSUPPORTED_OPERATION | 14 |  |
| EXTERNAL_BACKEND_COMMUNICATION_ERROR | 15 |  |
| EXTERNAL_BACKEND_ERROR | 16 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="ext_function-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## ext/function.proto



<a name="ext-FunctionCallRequest"></a>

### FunctionCallRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| request | [FunctionCallRequest.RequestEntry](#ext-FunctionCallRequest-RequestEntry) | repeated |  |






<a name="ext-FunctionCallRequest-RequestEntry"></a>

### FunctionCallRequest.RequestEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Any](#google-protobuf-Any) |  |  |






<a name="ext-FunctionCallResponse"></a>

### FunctionCallResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response | [FunctionCallResponse.ResponseEntry](#ext-FunctionCallResponse-ResponseEntry) | repeated |  |






<a name="ext-FunctionCallResponse-ResponseEntry"></a>

### FunctionCallResponse.ResponseEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Any](#google-protobuf-Any) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="ext-Function"></a>

### Function


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| FunctionCall | [FunctionCallRequest](#ext-FunctionCallRequest) | [FunctionCallResponse](#ext-FunctionCallResponse) |  |

 <!-- end services -->



<a name="model_batch-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## model/batch.proto



<a name="model-Batch"></a>

### Batch



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [BatchHeader](#model-BatchHeader) |  |  |
| resources | [Resource](#model-Resource) | repeated |  |
| batchRecords | [BatchRecordsPart](#model-BatchRecordsPart) | repeated |  |






<a name="model-BatchHeader"></a>

### BatchHeader



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mode | [BatchHeader.BatchMode](#model-BatchHeader-BatchMode) |  |  |
| annotations | [BatchHeader.AnnotationsEntry](#model-BatchHeader-AnnotationsEntry) | repeated |  |






<a name="model-BatchHeader-AnnotationsEntry"></a>

### BatchHeader.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="model-BatchRecordsPart"></a>

### BatchRecordsPart



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| values | [google.protobuf.Value](#google-protobuf-Value) | repeated |  |





 <!-- end messages -->


<a name="model-BatchHeader-BatchMode"></a>

### BatchHeader.BatchMode


| Name | Number | Description |
| ---- | ------ | ----------- |
| CREATE | 0 |  |
| UPDATE | 1 |  |
| DELETE | 2 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="model_data-source-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## model/data-source.proto



<a name="model-DataSource"></a>

### DataSource



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | id; read only |
| backend | [DataSourceBackendType](#model-DataSourceBackendType) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| postgresqlParams | [PostgresqlParams](#model-PostgresqlParams) |  |  |
| mysqlParams | [MysqlParams](#model-MysqlParams) |  |  |
| virtualParams | [VirtualParams](#model-VirtualParams) |  |  |
| redisParams | [RedisParams](#model-RedisParams) |  |  |
| mongoParams | [MongoParams](#model-MongoParams) |  |  |
| auditData | [AuditData](#model-AuditData) |  | read only |
| version | [uint32](#uint32) |  | read only |






<a name="model-DataSourceCatalog"></a>

### DataSourceCatalog



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| entities | [DataSourceEntity](#model-DataSourceEntity) | repeated |  |






<a name="model-DataSourceEntity"></a>

### DataSourceEntity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| readOnly | [bool](#bool) |  |  |






<a name="model-MongoParams"></a>

### MongoParams



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uri | [string](#string) |  |  |
| dbName | [string](#string) |  |  |






<a name="model-MysqlParams"></a>

### MysqlParams



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |
| password | [string](#string) |  |  |
| host | [string](#string) |  |  |
| port | [uint32](#uint32) |  |  |
| dbName | [string](#string) |  |  |
| defaultSchema | [string](#string) |  |  |






<a name="model-PostgresqlParams"></a>

### PostgresqlParams



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |
| password | [string](#string) |  |  |
| host | [string](#string) |  |  |
| port | [uint32](#uint32) |  |  |
| dbName | [string](#string) |  |  |
| defaultSchema | [string](#string) |  |  |






<a name="model-RedisParams"></a>

### RedisParams



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| addr | [string](#string) |  |  |
| password | [string](#string) |  |  |
| db | [int32](#int32) |  |  |






<a name="model-VirtualParams"></a>

### VirtualParams



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mode | [VirtualParams.Mode](#model-VirtualParams-Mode) |  |  |





 <!-- end messages -->


<a name="model-DataSourceBackendType"></a>

### DataSourceBackendType


| Name | Number | Description |
| ---- | ------ | ----------- |
| POSTGRESQL | 0 |  |
| VIRTUAL | 1 |  |
| MYSQL | 2 |  |
| ORACLE | 3 |  |
| MONGODB | 4 |  |
| REDIS | 5 |  |



<a name="model-VirtualParams-Mode"></a>

### VirtualParams.Mode


| Name | Number | Description |
| ---- | ------ | ----------- |
| DISCARD | 0 |  |
| ERROR | 1 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="model_external-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## model/external.proto



<a name="model-ExternalCall"></a>

### ExternalCall



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| functionCall | [FunctionCall](#model-FunctionCall) |  |  |
| httpCall | [HttpCall](#model-HttpCall) |  |  |






<a name="model-FunctionCall"></a>

### FunctionCall



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| host | [string](#string) |  |  |
| functionName | [string](#string) |  |  |






<a name="model-HttpCall"></a>

### HttpCall



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uri | [string](#string) |  |  |
| method | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="model_extension-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## model/extension.proto



<a name="model-Extension"></a>

### Extension



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| before | [Extension.Before](#model-Extension-Before) |  |  |
| instead | [Extension.Instead](#model-Extension-Instead) |  |  |
| after | [Extension.After](#model-Extension-After) |  |  |
| auditData | [AuditData](#model-AuditData) |  |  |
| version | [uint32](#uint32) |  |  |






<a name="model-Extension-After"></a>

### Extension.After



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| all | [ExternalCall](#model-ExternalCall) |  |  |
| create | [ExternalCall](#model-ExternalCall) |  |  |
| update | [ExternalCall](#model-ExternalCall) |  |  |
| delete | [ExternalCall](#model-ExternalCall) |  |  |
| get | [ExternalCall](#model-ExternalCall) |  |  |
| list | [ExternalCall](#model-ExternalCall) |  |  |
| sync | [bool](#bool) |  | if true, it will wait |






<a name="model-Extension-Before"></a>

### Extension.Before



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| all | [ExternalCall](#model-ExternalCall) |  |  |
| create | [ExternalCall](#model-ExternalCall) |  |  |
| update | [ExternalCall](#model-ExternalCall) |  |  |
| delete | [ExternalCall](#model-ExternalCall) |  |  |
| get | [ExternalCall](#model-ExternalCall) |  |  |
| list | [ExternalCall](#model-ExternalCall) |  |  |
| sync | [bool](#bool) |  | if true, it will wait |






<a name="model-Extension-Instead"></a>

### Extension.Instead



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| all | [ExternalCall](#model-ExternalCall) |  |  |
| create | [ExternalCall](#model-ExternalCall) |  |  |
| update | [ExternalCall](#model-ExternalCall) |  |  |
| delete | [ExternalCall](#model-ExternalCall) |  |  |
| get | [ExternalCall](#model-ExternalCall) |  |  |
| list | [ExternalCall](#model-ExternalCall) |  |  |
| finalize | [bool](#bool) |  | if true, it will respond with the result of the call, otherwise it will pass to backend |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="model_hcl-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## model/hcl.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="model_user-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## model/user.proto



<a name="model-User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| username | [string](#string) |  | principal |
| password | [string](#string) |  |  |
| securityContext | [SecurityContext](#model-SecurityContext) |  |  |
| details | [google.protobuf.Struct](#google-protobuf-Struct) |  |  |
| signKey | [string](#string) |  | rsa pub key to create tokens |
| auditData | [AuditData](#model-AuditData) |  |  |
| version | [uint32](#uint32) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="model_namespace-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## model/namespace.proto



<a name="model-Namespace"></a>

### Namespace



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| details | [google.protobuf.Struct](#google-protobuf-Struct) |  |  |
| securityContext | [SecurityContext](#model-SecurityContext) |  |  |
| auditData | [AuditData](#model-AuditData) |  |  |
| version | [uint32](#uint32) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="model_init-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## model/init.proto



<a name="model-AppConfig"></a>

### AppConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| host | [string](#string) |  |  |
| port | [int32](#int32) |  |  |
| jwtPrivateKey | [string](#string) |  |  |
| jwtPublicKey | [string](#string) |  |  |
| disableAuthentication | [bool](#bool) |  |  |
| disableCache | [bool](#bool) |  |  |
| pluginsPath | [string](#string) |  |  |






<a name="model-InitData"></a>

### InitData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [AppConfig](#model-AppConfig) |  |  |
| systemDataSource | [DataSource](#model-DataSource) |  |  |
| systemNamespace | [Namespace](#model-Namespace) |  |  |
| initDataSources | [DataSource](#model-DataSource) | repeated |  |
| initNamespaces | [Namespace](#model-Namespace) | repeated |  |
| initUsers | [User](#model-User) | repeated |  |
| initResources | [Resource](#model-Resource) | repeated |  |
| initRecords | [Record](#model-Record) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="model_resource-migration-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## model/resource-migration.proto



<a name="model-ResourceMigrationCreateIndex"></a>

### ResourceMigrationCreateIndex



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [uint32](#uint32) |  |  |






<a name="model-ResourceMigrationCreateProperty"></a>

### ResourceMigrationCreateProperty



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| property | [string](#string) |  |  |






<a name="model-ResourceMigrationCreateResource"></a>

### ResourceMigrationCreateResource







<a name="model-ResourceMigrationDeleteIndex"></a>

### ResourceMigrationDeleteIndex



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| existingIndex | [uint32](#uint32) |  |  |






<a name="model-ResourceMigrationDeleteProperty"></a>

### ResourceMigrationDeleteProperty



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| existingProperty | [string](#string) |  |  |






<a name="model-ResourceMigrationDeleteResource"></a>

### ResourceMigrationDeleteResource







<a name="model-ResourceMigrationPlan"></a>

### ResourceMigrationPlan



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| existingResource | [Resource](#model-Resource) |  |  |
| currentResource | [Resource](#model-Resource) |  |  |
| steps | [ResourceMigrationStep](#model-ResourceMigrationStep) | repeated |  |






<a name="model-ResourceMigrationStep"></a>

### ResourceMigrationStep



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| createResource | [ResourceMigrationCreateResource](#model-ResourceMigrationCreateResource) |  |  |
| deleteResource | [ResourceMigrationDeleteResource](#model-ResourceMigrationDeleteResource) |  |  |
| updateResource | [ResourceMigrationUpdateResource](#model-ResourceMigrationUpdateResource) |  |  |
| createProperty | [ResourceMigrationCreateProperty](#model-ResourceMigrationCreateProperty) |  |  |
| deleteProperty | [ResourceMigrationDeleteProperty](#model-ResourceMigrationDeleteProperty) |  |  |
| updateProperty | [ResourceMigrationUpdateProperty](#model-ResourceMigrationUpdateProperty) |  |  |
| createIndex | [ResourceMigrationCreateIndex](#model-ResourceMigrationCreateIndex) |  |  |
| deleteIndex | [ResourceMigrationDeleteIndex](#model-ResourceMigrationDeleteIndex) |  |  |






<a name="model-ResourceMigrationUpdateProperty"></a>

### ResourceMigrationUpdateProperty



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| existingProperty | [string](#string) |  |  |
| property | [string](#string) |  |  |
| changedFields | [string](#string) | repeated |  |






<a name="model-ResourceMigrationUpdateResource"></a>

### ResourceMigrationUpdateResource



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| changedFields | [string](#string) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="model_token-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## model/token.proto



<a name="model-Token"></a>

### Token



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| term | [TokenTerm](#model-TokenTerm) |  | issue term |
| content | [string](#string) |  | jwt token |
| expiration | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | expiration time |





 <!-- end messages -->


<a name="model-TokenTerm"></a>

### TokenTerm


| Name | Number | Description |
| ---- | ------ | ----------- |
| SHORT | 0 | 1 minute |
| MIDDLE | 1 | 2 hours |
| LONG | 2 | 2 days |
| VERY_LONG | 3 | 2 years |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="model_watch-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## model/watch.proto



<a name="model-WatchMessage"></a>

### WatchMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| changes | [google.protobuf.Struct](#google-protobuf-Struct) |  |  |
| recordIds | [string](#string) | repeated |  |
| event | [EventType](#model-EventType) |  |  |
| event_on | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |





 <!-- end messages -->


<a name="model-EventType"></a>

### EventType


| Name | Number | Description |
| ---- | ------ | ----------- |
| CREATE | 0 |  |
| UPDATE | 1 |  |
| DELETE | 2 |  |
| GET | 3 |  |
| LIST | 4 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="stub_authentication-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stub/authentication.proto



<a name="stub-AuthenticationRequest"></a>

### AuthenticationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |
| password | [string](#string) |  |  |
| term | [model.TokenTerm](#model-TokenTerm) |  |  |






<a name="stub-AuthenticationResponse"></a>

### AuthenticationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [model.Token](#model-Token) |  |  |






<a name="stub-RenewTokenRequest"></a>

### RenewTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| term | [model.TokenTerm](#model-TokenTerm) |  |  |






<a name="stub-RenewTokenResponse"></a>

### RenewTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [model.Token](#model-Token) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="stub-Authentication"></a>

### Authentication
Authentication Service is for authentication related operations

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Authenticate | [AuthenticationRequest](#stub-AuthenticationRequest) | [AuthenticationResponse](#stub-AuthenticationResponse) |  |
| RenewToken | [RenewTokenRequest](#stub-RenewTokenRequest) | [RenewTokenResponse](#stub-RenewTokenResponse) |  |

 <!-- end services -->



<a name="stub_data-source-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stub/data-source.proto



<a name="stub-CreateDataSourceRequest"></a>

### CreateDataSourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| dataSources | [model.DataSource](#model-DataSource) | repeated |  |






<a name="stub-CreateDataSourceResponse"></a>

### CreateDataSourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataSources | [model.DataSource](#model-DataSource) | repeated |  |






<a name="stub-DeleteDataSourceRequest"></a>

### DeleteDataSourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| ids | [string](#string) | repeated |  |






<a name="stub-DeleteDataSourceResponse"></a>

### DeleteDataSourceResponse







<a name="stub-GetDataSourceRequest"></a>

### GetDataSourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| id | [string](#string) |  |  |






<a name="stub-GetDataSourceResponse"></a>

### GetDataSourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataSource | [model.DataSource](#model-DataSource) |  |  |






<a name="stub-ListDataSourceRequest"></a>

### ListDataSourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |






<a name="stub-ListDataSourceResponse"></a>

### ListDataSourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| content | [model.DataSource](#model-DataSource) | repeated |  |






<a name="stub-ListEntitiesRequest"></a>

### ListEntitiesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| id | [string](#string) |  |  |






<a name="stub-ListEntitiesResponse"></a>

### ListEntitiesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| catalogs | [model.DataSourceCatalog](#model-DataSourceCatalog) | repeated |  |






<a name="stub-PrepareResourceFromEntityRequest"></a>

### PrepareResourceFromEntityRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| id | [string](#string) |  |  |
| catalog | [string](#string) |  |  |
| entity | [string](#string) |  |  |






<a name="stub-PrepareResourceFromEntityResponse"></a>

### PrepareResourceFromEntityResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resource | [model.Resource](#model-Resource) |  |  |






<a name="stub-StatusRequest"></a>

### StatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| id | [string](#string) |  |  |






<a name="stub-StatusResponse"></a>

### StatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| connectionAlreadyInitiated | [bool](#bool) |  |  |
| testConnection | [bool](#bool) |  |  |






<a name="stub-UpdateDataSourceRequest"></a>

### UpdateDataSourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| dataSources | [model.DataSource](#model-DataSource) | repeated |  |






<a name="stub-UpdateDataSourceResponse"></a>

### UpdateDataSourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataSources | [model.DataSource](#model-DataSource) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="stub-DataSource"></a>

### DataSource
DataSource Service is for managing data sources

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateDataSourceRequest](#stub-CreateDataSourceRequest) | [CreateDataSourceResponse](#stub-CreateDataSourceResponse) |  |
| List | [ListDataSourceRequest](#stub-ListDataSourceRequest) | [ListDataSourceResponse](#stub-ListDataSourceResponse) |  |
| Update | [UpdateDataSourceRequest](#stub-UpdateDataSourceRequest) | [UpdateDataSourceResponse](#stub-UpdateDataSourceResponse) |  |
| Delete | [DeleteDataSourceRequest](#stub-DeleteDataSourceRequest) | [DeleteDataSourceResponse](#stub-DeleteDataSourceResponse) |  |
| Get | [GetDataSourceRequest](#stub-GetDataSourceRequest) | [GetDataSourceResponse](#stub-GetDataSourceResponse) |  |
| Status | [StatusRequest](#stub-StatusRequest) | [StatusResponse](#stub-StatusResponse) |  |
| ListEntities | [ListEntitiesRequest](#stub-ListEntitiesRequest) | [ListEntitiesResponse](#stub-ListEntitiesResponse) |  |
| PrepareResourceFromEntity | [PrepareResourceFromEntityRequest](#stub-PrepareResourceFromEntityRequest) | [PrepareResourceFromEntityResponse](#stub-PrepareResourceFromEntityResponse) |  |

 <!-- end services -->



<a name="stub_extension-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stub/extension.proto



<a name="stub-CreateExtensionRequest"></a>

### CreateExtensionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| extensions | [model.Extension](#model-Extension) | repeated |  |






<a name="stub-CreateExtensionResponse"></a>

### CreateExtensionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| extensions | [model.Extension](#model-Extension) | repeated |  |






<a name="stub-DeleteExtensionRequest"></a>

### DeleteExtensionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| ids | [string](#string) | repeated |  |






<a name="stub-DeleteExtensionResponse"></a>

### DeleteExtensionResponse







<a name="stub-GetExtensionRequest"></a>

### GetExtensionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| id | [string](#string) |  |  |






<a name="stub-GetExtensionResponse"></a>

### GetExtensionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| extension | [model.Extension](#model-Extension) |  |  |






<a name="stub-ListExtensionRequest"></a>

### ListExtensionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |






<a name="stub-ListExtensionResponse"></a>

### ListExtensionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| content | [model.Extension](#model-Extension) | repeated |  |






<a name="stub-UpdateExtensionRequest"></a>

### UpdateExtensionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| extensions | [model.Extension](#model-Extension) | repeated |  |






<a name="stub-UpdateExtensionResponse"></a>

### UpdateExtensionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| extensions | [model.Extension](#model-Extension) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="stub-Extension"></a>

### Extension
Extension Service is for managing extensions

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| List | [ListExtensionRequest](#stub-ListExtensionRequest) | [ListExtensionResponse](#stub-ListExtensionResponse) |  |
| Get | [GetExtensionRequest](#stub-GetExtensionRequest) | [GetExtensionResponse](#stub-GetExtensionResponse) |  |
| Create | [CreateExtensionRequest](#stub-CreateExtensionRequest) | [CreateExtensionResponse](#stub-CreateExtensionResponse) |  |
| Update | [UpdateExtensionRequest](#stub-UpdateExtensionRequest) | [UpdateExtensionResponse](#stub-UpdateExtensionResponse) |  |
| Delete | [DeleteExtensionRequest](#stub-DeleteExtensionRequest) | [DeleteExtensionResponse](#stub-DeleteExtensionResponse) |  |

 <!-- end services -->



<a name="stub_generic-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stub/generic.proto



<a name="stub-CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| items | [google.protobuf.Any](#google-protobuf-Any) | repeated |  |
| ignoreIfExists | [bool](#bool) |  |  |
| annotations | [CreateRequest.AnnotationsEntry](#stub-CreateRequest-AnnotationsEntry) | repeated |  |






<a name="stub-CreateRequest-AnnotationsEntry"></a>

### CreateRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [google.protobuf.Any](#google-protobuf-Any) | repeated |  |
| inserted | [bool](#bool) | repeated |  |






<a name="stub-DeleteRequest"></a>

### DeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| id | [string](#string) | repeated |  |
| ids | [string](#string) | repeated |  |
| annotations | [DeleteRequest.AnnotationsEntry](#stub-DeleteRequest-AnnotationsEntry) | repeated |  |






<a name="stub-DeleteRequest-AnnotationsEntry"></a>

### DeleteRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-DeleteResponse"></a>

### DeleteResponse







<a name="stub-GetRequest"></a>

### GetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| id | [string](#string) |  |  |
| annotations | [GetRequest.AnnotationsEntry](#stub-GetRequest-AnnotationsEntry) | repeated |  |






<a name="stub-GetRequest-AnnotationsEntry"></a>

### GetRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [google.protobuf.Any](#google-protobuf-Any) |  |  |






<a name="stub-ListRequest"></a>

### ListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| filters | [ListRequest.FiltersEntry](#stub-ListRequest-FiltersEntry) | repeated |  |
| limit | [uint32](#uint32) |  |  |
| offset | [uint64](#uint64) |  |  |
| useHistory | [bool](#bool) |  |  |
| resolveReferences | [string](#string) | repeated |  |
| annotations | [ListRequest.AnnotationsEntry](#stub-ListRequest-AnnotationsEntry) | repeated |  |






<a name="stub-ListRequest-AnnotationsEntry"></a>

### ListRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-ListRequest-FiltersEntry"></a>

### ListRequest.FiltersEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [uint32](#uint32) |  |  |
| content | [google.protobuf.Any](#google-protobuf-Any) | repeated |  |






<a name="stub-SearchRequest"></a>

### SearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| query | [model.BooleanExpression](#model-BooleanExpression) |  |  |
| limit | [uint32](#uint32) |  |  |
| offset | [uint64](#uint64) |  |  |
| useHistory | [bool](#bool) |  |  |
| resolveReferences | [string](#string) | repeated |  |
| annotations | [SearchRequest.AnnotationsEntry](#stub-SearchRequest-AnnotationsEntry) | repeated |  |






<a name="stub-SearchRequest-AnnotationsEntry"></a>

### SearchRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-SearchResponse"></a>

### SearchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [uint32](#uint32) |  |  |
| content | [google.protobuf.Any](#google-protobuf-Any) | repeated |  |






<a name="stub-UpdateMultiRequest"></a>

### UpdateMultiRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| query | [model.BooleanExpression](#model-BooleanExpression) |  |  |
| properties | [UpdateMultiRequest.PropertiesEntry](#stub-UpdateMultiRequest-PropertiesEntry) | repeated |  |
| annotations | [UpdateMultiRequest.AnnotationsEntry](#stub-UpdateMultiRequest-AnnotationsEntry) | repeated |  |






<a name="stub-UpdateMultiRequest-AnnotationsEntry"></a>

### UpdateMultiRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-UpdateMultiRequest-PropertiesEntry"></a>

### UpdateMultiRequest.PropertiesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Value](#google-protobuf-Value) |  |  |






<a name="stub-UpdateMultiResponse"></a>

### UpdateMultiResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [google.protobuf.Any](#google-protobuf-Any) | repeated |  |






<a name="stub-UpdateRequest"></a>

### UpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| items | [google.protobuf.Any](#google-protobuf-Any) | repeated |  |
| checkVersion | [bool](#bool) |  |  |
| annotations | [UpdateRequest.AnnotationsEntry](#stub-UpdateRequest-AnnotationsEntry) | repeated |  |






<a name="stub-UpdateRequest-AnnotationsEntry"></a>

### UpdateRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [google.protobuf.Any](#google-protobuf-Any) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="stub-Generic"></a>

### Generic


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRequest](#stub-CreateRequest) | [CreateResponse](#stub-CreateResponse) |  |
| Update | [UpdateRequest](#stub-UpdateRequest) | [UpdateResponse](#stub-UpdateResponse) |  |
| UpdateMulti | [UpdateMultiRequest](#stub-UpdateMultiRequest) | [UpdateMultiResponse](#stub-UpdateMultiResponse) |  |
| Delete | [DeleteRequest](#stub-DeleteRequest) | [DeleteResponse](#stub-DeleteResponse) |  |
| List | [ListRequest](#stub-ListRequest) | [ListResponse](#stub-ListResponse) |  |
| Search | [SearchRequest](#stub-SearchRequest) | [SearchResponse](#stub-SearchResponse) |  |
| Get | [GetRequest](#stub-GetRequest) | [GetResponse](#stub-GetResponse) |  |

 <!-- end services -->



<a name="stub_namespace-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stub/namespace.proto



<a name="stub-CreateNamespaceRequest"></a>

### CreateNamespaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| Namespaces | [model.Namespace](#model-Namespace) | repeated |  |






<a name="stub-CreateNamespaceResponse"></a>

### CreateNamespaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Namespaces | [model.Namespace](#model-Namespace) | repeated |  |






<a name="stub-DeleteNamespaceRequest"></a>

### DeleteNamespaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| ids | [string](#string) | repeated |  |






<a name="stub-DeleteNamespaceResponse"></a>

### DeleteNamespaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Namespaces | [model.Namespace](#model-Namespace) | repeated |  |






<a name="stub-GetNamespaceRequest"></a>

### GetNamespaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| id | [string](#string) |  |  |






<a name="stub-GetNamespaceResponse"></a>

### GetNamespaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Namespace | [model.Namespace](#model-Namespace) |  |  |






<a name="stub-ListNamespaceRequest"></a>

### ListNamespaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |






<a name="stub-ListNamespaceResponse"></a>

### ListNamespaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| content | [model.Namespace](#model-Namespace) | repeated |  |






<a name="stub-UpdateNamespaceRequest"></a>

### UpdateNamespaceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| Namespaces | [model.Namespace](#model-Namespace) | repeated |  |






<a name="stub-UpdateNamespaceResponse"></a>

### UpdateNamespaceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Namespaces | [model.Namespace](#model-Namespace) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="stub-Namespace"></a>

### Namespace
Namespace Service is for managing namespaces

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateNamespaceRequest](#stub-CreateNamespaceRequest) | [CreateNamespaceResponse](#stub-CreateNamespaceResponse) |  |
| List | [ListNamespaceRequest](#stub-ListNamespaceRequest) | [ListNamespaceResponse](#stub-ListNamespaceResponse) |  |
| Update | [UpdateNamespaceRequest](#stub-UpdateNamespaceRequest) | [UpdateNamespaceResponse](#stub-UpdateNamespaceResponse) |  |
| Delete | [DeleteNamespaceRequest](#stub-DeleteNamespaceRequest) | [DeleteNamespaceResponse](#stub-DeleteNamespaceResponse) |  |
| Get | [GetNamespaceRequest](#stub-GetNamespaceRequest) | [GetNamespaceResponse](#stub-GetNamespaceResponse) |  |

 <!-- end services -->



<a name="stub_openapi-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stub/openapi.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="stub_record-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stub/record.proto



<a name="stub-ApplyRecordRequest"></a>

### ApplyRecordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| record | [model.Record](#model-Record) |  | Rest Only |
| records | [model.Record](#model-Record) | repeated |  |
| checkVersion | [bool](#bool) |  |  |
| annotations | [ApplyRecordRequest.AnnotationsEntry](#stub-ApplyRecordRequest-AnnotationsEntry) | repeated |  |






<a name="stub-ApplyRecordRequest-AnnotationsEntry"></a>

### ApplyRecordRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-ApplyRecordResponse"></a>

### ApplyRecordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| record | [model.Record](#model-Record) |  | Rest Only |
| records | [model.Record](#model-Record) | repeated |  |






<a name="stub-CreateRecordRequest"></a>

### CreateRecordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| record | [model.Record](#model-Record) |  |  |
| records | [model.Record](#model-Record) | repeated |  |
| ignoreIfExists | [bool](#bool) |  |  |
| annotations | [CreateRecordRequest.AnnotationsEntry](#stub-CreateRecordRequest-AnnotationsEntry) | repeated |  |






<a name="stub-CreateRecordRequest-AnnotationsEntry"></a>

### CreateRecordRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-CreateRecordResponse"></a>

### CreateRecordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| record | [model.Record](#model-Record) |  | Rest Only |
| records | [model.Record](#model-Record) | repeated |  |
| inserted | [bool](#bool) | repeated |  |






<a name="stub-DeleteRecordRequest"></a>

### DeleteRecordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| id | [string](#string) |  | Rest Only |
| ids | [string](#string) | repeated |  |
| annotations | [DeleteRecordRequest.AnnotationsEntry](#stub-DeleteRecordRequest-AnnotationsEntry) | repeated |  |






<a name="stub-DeleteRecordRequest-AnnotationsEntry"></a>

### DeleteRecordRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-DeleteRecordResponse"></a>

### DeleteRecordResponse







<a name="stub-GetRecordRequest"></a>

### GetRecordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| id | [string](#string) |  |  |
| annotations | [GetRecordRequest.AnnotationsEntry](#stub-GetRecordRequest-AnnotationsEntry) | repeated |  |






<a name="stub-GetRecordRequest-AnnotationsEntry"></a>

### GetRecordRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-GetRecordResponse"></a>

### GetRecordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| record | [model.Record](#model-Record) |  |  |






<a name="stub-ListRecordRequest"></a>

### ListRecordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| filters | [ListRecordRequest.FiltersEntry](#stub-ListRecordRequest-FiltersEntry) | repeated |  |
| limit | [uint32](#uint32) |  |  |
| offset | [uint64](#uint64) |  |  |
| useHistory | [bool](#bool) |  |  |
| resolveReferences | [string](#string) | repeated |  |
| annotations | [ListRecordRequest.AnnotationsEntry](#stub-ListRecordRequest-AnnotationsEntry) | repeated |  |






<a name="stub-ListRecordRequest-AnnotationsEntry"></a>

### ListRecordRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-ListRecordRequest-FiltersEntry"></a>

### ListRecordRequest.FiltersEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-ListRecordResponse"></a>

### ListRecordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [uint32](#uint32) |  |  |
| content | [model.Record](#model-Record) | repeated |  |






<a name="stub-ReadStreamRequest"></a>

### ReadStreamRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| query | [model.BooleanExpression](#model-BooleanExpression) |  |  |
| limit | [uint32](#uint32) |  |  |
| offset | [uint64](#uint64) |  |  |
| useHistory | [bool](#bool) |  |  |
| resolveReferences | [string](#string) | repeated |  |
| useTransaction | [bool](#bool) |  |  |
| packRecords | [bool](#bool) |  |  |
| annotations | [ReadStreamRequest.AnnotationsEntry](#stub-ReadStreamRequest-AnnotationsEntry) | repeated |  |






<a name="stub-ReadStreamRequest-AnnotationsEntry"></a>

### ReadStreamRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-SearchRecordRequest"></a>

### SearchRecordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| query | [model.BooleanExpression](#model-BooleanExpression) |  |  |
| limit | [uint32](#uint32) |  |  |
| offset | [uint64](#uint64) |  |  |
| useHistory | [bool](#bool) |  |  |
| resolveReferences | [string](#string) | repeated |  |
| annotations | [SearchRecordRequest.AnnotationsEntry](#stub-SearchRecordRequest-AnnotationsEntry) | repeated |  |






<a name="stub-SearchRecordRequest-AnnotationsEntry"></a>

### SearchRecordRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-SearchRecordResponse"></a>

### SearchRecordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [uint32](#uint32) |  |  |
| content | [model.Record](#model-Record) | repeated |  |






<a name="stub-UpdateMultiRecordRequest"></a>

### UpdateMultiRecordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| query | [model.BooleanExpression](#model-BooleanExpression) |  |  |
| properties | [UpdateMultiRecordRequest.PropertiesEntry](#stub-UpdateMultiRecordRequest-PropertiesEntry) | repeated |  |
| annotations | [UpdateMultiRecordRequest.AnnotationsEntry](#stub-UpdateMultiRecordRequest-AnnotationsEntry) | repeated |  |






<a name="stub-UpdateMultiRecordRequest-AnnotationsEntry"></a>

### UpdateMultiRecordRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-UpdateMultiRecordRequest-PropertiesEntry"></a>

### UpdateMultiRecordRequest.PropertiesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Value](#google-protobuf-Value) |  |  |






<a name="stub-UpdateMultiRecordResponse"></a>

### UpdateMultiRecordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| record | [model.Record](#model-Record) |  | Rest Only |
| records | [model.Record](#model-Record) | repeated |  |






<a name="stub-UpdateRecordRequest"></a>

### UpdateRecordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| record | [model.Record](#model-Record) |  | Rest Only |
| records | [model.Record](#model-Record) | repeated |  |
| checkVersion | [bool](#bool) |  |  |
| annotations | [UpdateRecordRequest.AnnotationsEntry](#stub-UpdateRecordRequest-AnnotationsEntry) | repeated |  |






<a name="stub-UpdateRecordRequest-AnnotationsEntry"></a>

### UpdateRecordRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-UpdateRecordResponse"></a>

### UpdateRecordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| record | [model.Record](#model-Record) |  | Rest Only |
| records | [model.Record](#model-Record) | repeated |  |






<a name="stub-WriteStreamResponse"></a>

### WriteStreamResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) | repeated |  |
| created | [bool](#bool) | repeated |  |
| updated | [bool](#bool) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="stub-Record"></a>

### Record
Record service is an abstract service for records of all resources. You can do CRUD like operations with Record service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRecordRequest](#stub-CreateRecordRequest) | [CreateRecordResponse](#stub-CreateRecordResponse) |  |
| Update | [UpdateRecordRequest](#stub-UpdateRecordRequest) | [UpdateRecordResponse](#stub-UpdateRecordResponse) |  |
| Apply | [ApplyRecordRequest](#stub-ApplyRecordRequest) | [ApplyRecordResponse](#stub-ApplyRecordResponse) |  |
| UpdateMulti | [UpdateMultiRecordRequest](#stub-UpdateMultiRecordRequest) | [UpdateMultiRecordResponse](#stub-UpdateMultiRecordResponse) |  |
| Delete | [DeleteRecordRequest](#stub-DeleteRecordRequest) | [DeleteRecordResponse](#stub-DeleteRecordResponse) |  |
| List | [ListRecordRequest](#stub-ListRecordRequest) | [ListRecordResponse](#stub-ListRecordResponse) |  |
| Search | [SearchRecordRequest](#stub-SearchRecordRequest) | [SearchRecordResponse](#stub-SearchRecordResponse) |  |
| ReadStream | [ReadStreamRequest](#stub-ReadStreamRequest) | [.model.Record](#model-Record) stream |  |
| WriteStream | [.model.Record](#model-Record) stream | [WriteStreamResponse](#stub-WriteStreamResponse) |  |
| Get | [GetRecordRequest](#stub-GetRecordRequest) | [GetRecordResponse](#stub-GetRecordResponse) |  |

 <!-- end services -->



<a name="stub_resource-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stub/resource.proto



<a name="stub-CreateResourceRequest"></a>

### CreateResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| resources | [model.Resource](#model-Resource) | repeated |  |
| doMigration | [bool](#bool) |  |  |
| forceMigration | [bool](#bool) |  |  |
| annotations | [CreateResourceRequest.AnnotationsEntry](#stub-CreateResourceRequest-AnnotationsEntry) | repeated |  |






<a name="stub-CreateResourceRequest-AnnotationsEntry"></a>

### CreateResourceRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-CreateResourceResponse"></a>

### CreateResourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resources | [model.Resource](#model-Resource) | repeated |  |






<a name="stub-DeleteResourceRequest"></a>

### DeleteResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| ids | [string](#string) | repeated |  |
| doMigration | [bool](#bool) |  |  |
| forceMigration | [bool](#bool) |  |  |
| annotations | [DeleteResourceRequest.AnnotationsEntry](#stub-DeleteResourceRequest-AnnotationsEntry) | repeated |  |






<a name="stub-DeleteResourceRequest-AnnotationsEntry"></a>

### DeleteResourceRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-DeleteResourceResponse"></a>

### DeleteResourceResponse







<a name="stub-GetResourceByNameRequest"></a>

### GetResourceByNameRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| name | [string](#string) |  |  |
| annotations | [GetResourceByNameRequest.AnnotationsEntry](#stub-GetResourceByNameRequest-AnnotationsEntry) | repeated |  |






<a name="stub-GetResourceByNameRequest-AnnotationsEntry"></a>

### GetResourceByNameRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-GetResourceByNameResponse"></a>

### GetResourceByNameResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resource | [model.Resource](#model-Resource) |  |  |






<a name="stub-GetResourceRequest"></a>

### GetResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| id | [string](#string) |  |  |
| annotations | [GetResourceRequest.AnnotationsEntry](#stub-GetResourceRequest-AnnotationsEntry) | repeated |  |






<a name="stub-GetResourceRequest-AnnotationsEntry"></a>

### GetResourceRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-GetResourceResponse"></a>

### GetResourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resource | [model.Resource](#model-Resource) |  |  |






<a name="stub-GetSystemResourceRequest"></a>

### GetSystemResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| name | [string](#string) |  |  |
| annotations | [GetSystemResourceRequest.AnnotationsEntry](#stub-GetSystemResourceRequest-AnnotationsEntry) | repeated |  |






<a name="stub-GetSystemResourceRequest-AnnotationsEntry"></a>

### GetSystemResourceRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-GetSystemResourceResponse"></a>

### GetSystemResourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resource | [model.Resource](#model-Resource) |  |  |






<a name="stub-ListResourceRequest"></a>

### ListResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| annotations | [ListResourceRequest.AnnotationsEntry](#stub-ListResourceRequest-AnnotationsEntry) | repeated |  |






<a name="stub-ListResourceRequest-AnnotationsEntry"></a>

### ListResourceRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-ListResourceResponse"></a>

### ListResourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resources | [model.Resource](#model-Resource) | repeated |  |






<a name="stub-PrepareResourceMigrationPlanRequest"></a>

### PrepareResourceMigrationPlanRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| prepareFromDataSource | [bool](#bool) |  |  |
| resources | [model.Resource](#model-Resource) | repeated |  |
| annotations | [PrepareResourceMigrationPlanRequest.AnnotationsEntry](#stub-PrepareResourceMigrationPlanRequest-AnnotationsEntry) | repeated |  |






<a name="stub-PrepareResourceMigrationPlanRequest-AnnotationsEntry"></a>

### PrepareResourceMigrationPlanRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-PrepareResourceMigrationPlanResponse"></a>

### PrepareResourceMigrationPlanResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| plans | [model.ResourceMigrationPlan](#model-ResourceMigrationPlan) | repeated |  |






<a name="stub-UpdateResourceRequest"></a>

### UpdateResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| resources | [model.Resource](#model-Resource) | repeated |  |
| doMigration | [bool](#bool) |  |  |
| forceMigration | [bool](#bool) |  |  |
| annotations | [UpdateResourceRequest.AnnotationsEntry](#stub-UpdateResourceRequest-AnnotationsEntry) | repeated |  |






<a name="stub-UpdateResourceRequest-AnnotationsEntry"></a>

### UpdateResourceRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="stub-UpdateResourceResponse"></a>

### UpdateResourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resources | [model.Resource](#model-Resource) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="stub-Resource"></a>

### Resource
Resource service is for managing resources

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateResourceRequest](#stub-CreateResourceRequest) | [CreateResourceResponse](#stub-CreateResourceResponse) |  |
| Update | [UpdateResourceRequest](#stub-UpdateResourceRequest) | [UpdateResourceResponse](#stub-UpdateResourceResponse) |  |
| Delete | [DeleteResourceRequest](#stub-DeleteResourceRequest) | [DeleteResourceResponse](#stub-DeleteResourceResponse) |  |
| List | [ListResourceRequest](#stub-ListResourceRequest) | [ListResourceResponse](#stub-ListResourceResponse) |  |
| PrepareResourceMigrationPlan | [PrepareResourceMigrationPlanRequest](#stub-PrepareResourceMigrationPlanRequest) | [PrepareResourceMigrationPlanResponse](#stub-PrepareResourceMigrationPlanResponse) |  |
| Get | [GetResourceRequest](#stub-GetResourceRequest) | [GetResourceResponse](#stub-GetResourceResponse) |  |
| GetByName | [GetResourceByNameRequest](#stub-GetResourceByNameRequest) | [GetResourceByNameResponse](#stub-GetResourceByNameResponse) |  |
| GetSystemResource | [GetSystemResourceRequest](#stub-GetSystemResourceRequest) | [GetSystemResourceResponse](#stub-GetSystemResourceResponse) |  |

 <!-- end services -->



<a name="stub_rest_record-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stub/rest/record.proto



<a name="rest-ApplyRecordRequest"></a>

### ApplyRecordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| id | [string](#string) |  | id; read only |
| properties | [ApplyRecordRequest.PropertiesEntry](#rest-ApplyRecordRequest-PropertiesEntry) | repeated |  |
| checkVersion | [bool](#bool) |  |  |
| annotations | [ApplyRecordRequest.AnnotationsEntry](#rest-ApplyRecordRequest-AnnotationsEntry) | repeated |  |






<a name="rest-ApplyRecordRequest-AnnotationsEntry"></a>

### ApplyRecordRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="rest-ApplyRecordRequest-PropertiesEntry"></a>

### ApplyRecordRequest.PropertiesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Value](#google-protobuf-Value) |  |  |






<a name="rest-ApplyRecordResponse"></a>

### ApplyRecordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| properties | [ApplyRecordResponse.PropertiesEntry](#rest-ApplyRecordResponse-PropertiesEntry) | repeated |  |






<a name="rest-ApplyRecordResponse-PropertiesEntry"></a>

### ApplyRecordResponse.PropertiesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Value](#google-protobuf-Value) |  |  |






<a name="rest-CreateRecordRequest"></a>

### CreateRecordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| properties | [CreateRecordRequest.PropertiesEntry](#rest-CreateRecordRequest-PropertiesEntry) | repeated |  |
| annotations | [CreateRecordRequest.AnnotationsEntry](#rest-CreateRecordRequest-AnnotationsEntry) | repeated |  |






<a name="rest-CreateRecordRequest-AnnotationsEntry"></a>

### CreateRecordRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="rest-CreateRecordRequest-PropertiesEntry"></a>

### CreateRecordRequest.PropertiesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Value](#google-protobuf-Value) |  |  |






<a name="rest-CreateRecordResponse"></a>

### CreateRecordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | id; read only |
| properties | [CreateRecordResponse.PropertiesEntry](#rest-CreateRecordResponse-PropertiesEntry) | repeated |  |






<a name="rest-CreateRecordResponse-PropertiesEntry"></a>

### CreateRecordResponse.PropertiesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Value](#google-protobuf-Value) |  |  |






<a name="rest-DeleteRecordRequest"></a>

### DeleteRecordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| id | [string](#string) |  | Rest Only |
| annotations | [DeleteRecordRequest.AnnotationsEntry](#rest-DeleteRecordRequest-AnnotationsEntry) | repeated |  |






<a name="rest-DeleteRecordRequest-AnnotationsEntry"></a>

### DeleteRecordRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="rest-DeleteRecordResponse"></a>

### DeleteRecordResponse







<a name="rest-UpdateRecordRequest"></a>

### UpdateRecordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| id | [string](#string) |  | id; read only |
| properties | [UpdateRecordRequest.PropertiesEntry](#rest-UpdateRecordRequest-PropertiesEntry) | repeated |  |
| checkVersion | [bool](#bool) |  |  |
| annotations | [UpdateRecordRequest.AnnotationsEntry](#rest-UpdateRecordRequest-AnnotationsEntry) | repeated |  |






<a name="rest-UpdateRecordRequest-AnnotationsEntry"></a>

### UpdateRecordRequest.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="rest-UpdateRecordRequest-PropertiesEntry"></a>

### UpdateRecordRequest.PropertiesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Value](#google-protobuf-Value) |  |  |






<a name="rest-UpdateRecordResponse"></a>

### UpdateRecordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| properties | [UpdateRecordResponse.PropertiesEntry](#rest-UpdateRecordResponse-PropertiesEntry) | repeated |  |






<a name="rest-UpdateRecordResponse-PropertiesEntry"></a>

### UpdateRecordResponse.PropertiesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Value](#google-protobuf-Value) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="rest-Record"></a>

### Record
Record service is an abstract service for records of all resources. You can do CRUD like operations with Record service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateRecordRequest](#rest-CreateRecordRequest) | [CreateRecordResponse](#rest-CreateRecordResponse) |  |
| Update | [UpdateRecordRequest](#rest-UpdateRecordRequest) | [UpdateRecordResponse](#rest-UpdateRecordResponse) |  |
| Apply | [ApplyRecordRequest](#rest-ApplyRecordRequest) | [ApplyRecordResponse](#rest-ApplyRecordResponse) |  |
| Delete | [DeleteRecordRequest](#rest-DeleteRecordRequest) | [DeleteRecordResponse](#rest-DeleteRecordResponse) |  |

 <!-- end services -->



<a name="stub_user-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stub/user.proto



<a name="stub-CreateUserRequest"></a>

### CreateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| user | [model.User](#model-User) |  |  |
| users | [model.User](#model-User) | repeated |  |






<a name="stub-CreateUserResponse"></a>

### CreateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [model.User](#model-User) |  |  |
| users | [model.User](#model-User) | repeated |  |






<a name="stub-DeleteUserRequest"></a>

### DeleteUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| id | [string](#string) |  |  |
| ids | [string](#string) | repeated |  |






<a name="stub-DeleteUserResponse"></a>

### DeleteUserResponse







<a name="stub-GetUserRequest"></a>

### GetUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| id | [string](#string) |  |  |






<a name="stub-GetUserResponse"></a>

### GetUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [model.User](#model-User) |  |  |






<a name="stub-ListUserRequest"></a>

### ListUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| limit | [uint32](#uint32) |  |  |
| offset | [uint64](#uint64) |  |  |






<a name="stub-ListUserResponse"></a>

### ListUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| content | [model.User](#model-User) | repeated |  |






<a name="stub-UpdateUserRequest"></a>

### UpdateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| user | [model.User](#model-User) |  |  |
| users | [model.User](#model-User) | repeated |  |






<a name="stub-UpdateUserResponse"></a>

### UpdateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [model.User](#model-User) |  |  |
| users | [model.User](#model-User) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="stub-User"></a>

### User
User service is for managing users

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Create | [CreateUserRequest](#stub-CreateUserRequest) | [CreateUserResponse](#stub-CreateUserResponse) |  |
| Update | [UpdateUserRequest](#stub-UpdateUserRequest) | [UpdateUserResponse](#stub-UpdateUserResponse) |  |
| Delete | [DeleteUserRequest](#stub-DeleteUserRequest) | [DeleteUserResponse](#stub-DeleteUserResponse) |  |
| List | [ListUserRequest](#stub-ListUserRequest) | [ListUserResponse](#stub-ListUserResponse) |  |
| Get | [GetUserRequest](#stub-GetUserRequest) | [GetUserResponse](#stub-GetUserResponse) |  |

 <!-- end services -->



<a name="stub_watch-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## stub/watch.proto



<a name="stub-WatchRequest"></a>

### WatchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| resource | [string](#string) |  |  |
| query | [model.BooleanExpression](#model-BooleanExpression) |  |  |
| events | [model.EventType](#model-EventType) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="stub-Watch"></a>

### Watch
Watch service watching operations on records

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Watch | [WatchRequest](#stub-WatchRequest) | [.model.WatchMessage](#model-WatchMessage) stream | Sends a greeting |

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

