export interface CrudFormOptions {
    label: string;
    value: string;
    selected: boolean;
    disabled: boolean;
    visible: boolean;
    readOnly: boolean;
    defaultValue: string;

}

export interface CrudFormConfigFieldType {
    name: string;
    label: string;
    type: string;
    required: boolean;
    readOnly: boolean;
    visible: boolean;
    defaultValue: string;
    options: boolean[];

}

export interface CrudFormConfig {
    fields: boolean[];

}

export interface CrudGridColumnConfig {
    header: string;
    visible: boolean;
    width: number;
    sortable: boolean;
    filterable: boolean;

}

export interface CrudGridConfig {
    defaultColumnConfig: CrudGridColumnConfig;
    columns: CrudGridColumnConfig[];

}


export interface Crud {
    id: string;
    resource: string;
    namespace: string;
    gridConfig: CrudGridConfig;
    formConfig: CrudFormConfig;
    version: number;

}


export interface AppDesignerBoard {
    id: string;
    name: string;
    description: string;
    resourceSelector: string[];
    resourceVisuals: {}[];
    version: number;

}
    

