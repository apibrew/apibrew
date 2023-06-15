export interface Root {
    file: File[]
}

export interface File {
    name: string
    package: string
    dependency: string[]
    extension?: Extension[]
    options: Options
    syntax: string
    bufExtension: BufExtension
    messageType?: MessageType[]
    enumType?: EnumType2[]
    service?: Service[]
}

export interface Extension {
    name: string
    number: number
    label: string
    type: string
    typeName?: string
    extendee: string
    jsonName: string
    proto3Optional?: boolean
}

export interface Options {
    goPackage: string
    javaPackage?: string
    javaOuterClassname?: string
    javaMultipleFiles?: boolean
    objcClassPrefix?: string
    "[gnostic.openapi.v3.document]"?: GnosticOpenapiV3Document
}

export interface GnosticOpenapiV3Document {
    info: Info
    components: Components
    security: Security[]
}

export interface Info {
    title: string
    description: string
    contact: Contact
    license: License
    version: string
}

export interface Contact {
    name: string
    url: string
    email: string
}

export interface License {
    name: string
    url: string
}

export interface Components {
    securitySchemes: SecuritySchemes
}

export interface SecuritySchemes {
    additionalProperties: AdditionalProperty[]
}

export interface AdditionalProperty {
    name: string
    value: Value
}

export interface Value {
    securityScheme: SecurityScheme
}

export interface SecurityScheme {
    type: string
    scheme: string
    bearerFormat: string
}

export interface Security {
    additionalProperties: AdditionalProperty2[]
}

export interface AdditionalProperty2 {
    name: string
    value: Value2
}

export interface Value2 {}

export interface BufExtension {
    isImport: boolean
    isSyntaxUnspecified: boolean
    unusedDependency?: number[]
}

export interface MessageType {
    name: string
    field?: Field[]
    nestedType?: NestedType[]
    oneofDecl?: OneofDecl[]
    options?: Options4
    enumType?: EnumType[]
}

export interface Field {
    name: string
    number: number
    label: string
    type: string
    typeName?: string
    jsonName: string
    oneofIndex?: number
    proto3Optional?: boolean
    options?: Options2
}

export interface Options2 {
    "[model.hcl_ignore]"?: boolean
    "[model.propertyType]"?: string
    "[model.propertyUnique]"?: boolean
    "[model.hcl_block]"?: string
    "[model.hcl_label]"?: string
    "[model.propertyAnnotations]"?: ModelPropertyAnnotations[]
    "[model.propertyMapping]"?: string
}

export interface ModelPropertyAnnotations {
    name: string
    value: string
}

export interface NestedType {
    name: string
    field: Field2[]
    options?: Options3
}

export interface Field2 {
    name: string
    number: number
    label: string
    type: string
    typeName?: string
    jsonName: string
}

export interface Options3 {
    mapEntry: boolean
}

export interface OneofDecl {
    name: string
}

export interface Options4 {
    "[model.resourceDataSource]": string
    "[model.resourceEntity]"?: string
    "[model.resourceName]"?: string
    "[model.resourceNamespace]"?: string
    "[model.securityContextDisallowAll]"?: boolean
}

export interface EnumType {
    name: string
    value: Value3[]
}

export interface Value3 {
    name: string
    number: number
}

export interface EnumType2 {
    name: string
    value: Value4[]
}

export interface Value4 {
    name: string
    number: number
}

export interface Service {
    name: string
    method: Method[]
}

export interface Method {
    name: string
    inputType: string
    outputType: string
    options: Options5
    serverStreaming?: boolean
    clientStreaming?: boolean
}

export interface Options5 {
    "[google.api.http]"?: GoogleApiHttp
    "[gnostic.openapi.v3.operation]"?: GnosticOpenapiV3Operation
    "[stub.resources]"?: StubResources
}

export interface GoogleApiHttp {
    post?: string
    additionalBindings?: AdditionalBinding[]
    put?: string
    delete?: string
    get?: string
    responseBody?: string
    body?: string
    patch?: string
}

export interface AdditionalBinding {
    post?: string
    body?: string
    responseBody?: string
    put?: string
    delete?: string
}

export interface GnosticOpenapiV3Operation {
    tags: string[]
    summary: string
    description: string
    operationId: string
}

export interface StubResources {
    operation: Operation[]
}

export interface Operation {
    abc: string
}
