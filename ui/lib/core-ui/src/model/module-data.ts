interface ExportOptions {
    [key: string]: any
}

export interface ModuleData {
    name: string
    package: string
    exports: ExportOptions
}
