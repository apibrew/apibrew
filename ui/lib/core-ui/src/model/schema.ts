export interface AppDesignerBoardResourceVisual {
    resource: string;
    allowRecordsOnBoard: boolean;
    location: {
        x: number;
        y: number;

    };

}


export interface CrudFormItem {
    kind: string;
    propertyPath: string;
    title: string;
    component: string;
    children: CrudFormItem[];
    params: object;

}

export interface CrudFormConfig {
    children: CrudFormItem[];

}

export interface CrudGridConfig {

}


export const AppDesignerBoardName = "AppDesignerBoard";

export const AppDesignerBoardIdName = "Id";

export const AppDesignerBoardDescriptionName = "Description";

export const AppDesignerBoardNameName = "Name";

export const AppDesignerBoardNamespaceSelectorName = "NamespaceSelector";

export const AppDesignerBoardVersionName = "Version";

export const AppDesignerBoardResourceVisualsName = "ResourceVisuals";

export const AppDesignerBoardResourceSelectorName = "ResourceSelector";

export interface AppDesignerBoard {
    id: string;
    description: string;
    name: string;
    namespaceSelector: string[];
    version: number;
    resourceVisuals: AppDesignerBoardResourceVisual[];
    resourceSelector: string[];

}


export const ModuleName = "Module";

export const ModuleIdName = "Id";

export const ModuleNameName = "Name";

export const ModulePackageName = "Package";

export const ModuleDescriptionName = "Description";

export const ModuleSourceName = "Source";

export const ModuleComponentsName = "Components";

export const ModuleVersionName = "Version";

export interface Module {
    id: string;
    name: string;
    package: string;
    description: string;
    source: string;
    components: string[];
    version: number;

}


export const CrudName = "Crud";

export const CrudIdName = "Id";

export const CrudNameName = "Name";

export const CrudResourceName = "Resource";

export const CrudNamespaceName = "Namespace";

export const CrudGridConfigName = "GridConfig";

export const CrudFormConfigName = "FormConfig";

export const CrudVersionName = "Version";

export interface Crud {
    id: string;
    name: string;
    resource: string;
    namespace: string;
    gridConfig: CrudGridConfig;
    formConfig: CrudFormConfig;
    version: number;

}


