export interface AppDesignerBoardResourceVisual {
    resource: string;
    allowRecordsOnBoard: boolean;
    location: {
        x: number;
        y: number;

    };

}


export interface CrudFormConfigFieldType {
    name: string;
    label: string;
    type: string;
    required: boolean;
    readOnly: boolean;
    visible: boolean;
    defaultValue: string;

}

export interface CrudFormConfig {
    fields: CrudFormConfigFieldType[];

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


export const AppDesignerBoardName = "AppDesignerBoard";

export const AppDesignerBoardIdName = "Id";

export const AppDesignerBoardDescriptionName = "Description";

export const AppDesignerBoardNameName = "Name";

export const AppDesignerBoardVersionName = "Version";

export const AppDesignerBoardResourceVisualsName = "ResourceVisuals";

export const AppDesignerBoardResourceSelectorName = "ResourceSelector";

export interface AppDesignerBoard {
    id: string;
    description: string;
    name: string;
    version: number;
    resourceVisuals: AppDesignerBoardResourceVisual[];
    resourceSelector: string[];

}


export const CrudName = "Crud";

export const CrudIdName = "Id";

export const CrudResourceName = "Resource";

export const CrudNamespaceName = "Namespace";

export const CrudGridConfigName = "GridConfig";

export const CrudFormConfigName = "FormConfig";

export const CrudVersionName = "Version";

export interface Crud {
    id: string;
    resource: string;
    namespace: string;
    gridConfig: CrudGridConfig;
    formConfig: CrudFormConfig;
    version: number;

}
    

