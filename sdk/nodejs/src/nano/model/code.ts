
export interface Code {
    id: string
    name: string
    language: Language
    content: string
    contentFormat: ContentFormat
    annotations?: { [key: string]: string }
    version: number
    auditData?: AuditData
}

export const CodeEntityInfo = {
    namespace: "nano",
    resource: "Code",
    restPath: "nano-code",
}

export interface AuditData {
    createdBy: string
    updatedBy: string
    createdOn: string | Date
    updatedOn: string | Date
}

export enum Language {
    PYTHON = "PYTHON",
    JAVASCRIPT = "JAVASCRIPT",
}

export enum ContentFormat {
    TEXT = "TEXT",
    TAR = "TAR",
    TAR_GZ = "TAR_GZ",
}


