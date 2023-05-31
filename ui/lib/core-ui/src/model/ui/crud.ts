


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
     disablePagination?: number;
     defaultPageSize?: boolean;
     actions?: GridActionConfig[];
     sizeMode?: string;

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
     component?: string;

}

// Resource Type
export interface Crud {
    namespace: string;
gridConfig?: GridConfig;
formConfig?: FormConfig;
version: number;
id: string;
name?: string;
resource: string;

}
// Resource and Property Names
export const CrudName = "Crud";

export const CrudNamespaceName = "Namespace";

export const CrudGridConfigName = "GridConfig";

export const CrudFormConfigName = "FormConfig";

export const CrudVersionName = "Version";

export const CrudIdName = "Id";

export const CrudNameName = "Name";

export const CrudResourceName = "Resource";


