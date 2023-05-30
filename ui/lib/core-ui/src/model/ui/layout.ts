


// Sub Types

export interface LayoutComponent {
     component: string;
     syntax?: object;
     params?: object;
     children?: LayoutComponent[];

}

export interface InputParameter {
     name: string;
     type: string;

}

// Resource Type
export interface Layout {
    id: string;
name: string;
input?: InputParameter[];
layoutComponent: LayoutComponent;
version: number;

}
// Resource and Property Names
export const LayoutName = "Layout";

export const LayoutIdName = "Id";

export const LayoutNameName = "Name";

export const LayoutInputName = "Input";

export const LayoutLayoutComponentName = "LayoutComponent";

export const LayoutVersionName = "Version";


