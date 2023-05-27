


// Sub Types

export interface FormItem {
     kind?: string;
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
     actions?: GridActionConfig[];

}

export interface GridActionConfig {
     name: string;
     title?: string;
     icon?: string;
     component?: Component;

}

export interface GridColumnConfig {
     name: string;
     title?: string;
     type?: string;
     propertyPath?: string;
     width?: number;
     sortable?: boolean;
     filterable?: boolean;
     hidden?: boolean;
     component?: Component;

}

export interface Component {
     package?: string;
     name?: string;
     componentName?: string;
     params?: object;

}

// Resource Type
export interface Crud {
    id: string;
name?: string;
resource: string;
namespace: string;
gridConfig?: GridConfig;
formConfig?: FormConfig;
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

export const CrudVersionName = "Version";


