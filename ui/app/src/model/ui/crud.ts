


export const CrudResource = {
    resource: "Crud",
    namespace: "ui",
};

// Sub Types

export interface FormItem {
     kind?: 'property' | 'input' | 'tab' | 'section' | 'group' | 'custom' | 'container';
     propertyPath?: string;
     title?: string;
     component?: string;
     children?: FormItem[];
     params?: object;

}

export interface FormConfig {
     children?: FormItem[];

}

export interface GridConfig {
     columns?: GridColumnConfig[];
     disableDefaultActions?: boolean;
     disablePagination?: number;
     defaultPageSize?: boolean;
     actions?: GridActionConfig[];
     sizeMode?: 'compact' | 'normal';

}

export interface GridActionConfig {
     name: string;
     title?: string;
     icon?: string;
     component?: string;

}

export interface GridColumnConfig {
     name: string;
     title?: string;
     type?: string;
     width?: number;
     flex?: number;
     sortable?: boolean;
     filterable?: boolean;
     hidden?: boolean;
     disabled?: boolean;
     component?: string;

}

// Resource Type
export interface Crud {
    id: string;
name?: string;
resource: string;
namespace: string;
gridConfig?: GridConfig;
formConfig?: FormConfig;
hideSettings?: boolean;
version: number;

}
// Resource and Property Names
export const CrudName = "Crud";

export const CrudIdName = "Id";

export const CrudNameName = "Name";

export const CrudResourceName = "Resource";

export const CrudNamespaceName = "Namespace";

export const CrudGridConfigName = "GridConfig";

export const CrudFormConfigName = "FormConfig";

export const CrudHideSettingsName = "HideSettings";

export const CrudVersionName = "Version";


