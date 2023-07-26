


export const MenuResource = {
    resource: "Menu",
    namespace: "ui",
};

// Sub Types

export interface RequireSecurityConstraintParams {
     resource?: string;
     namespace?: string;
     operation?: 'READ' | 'WRITE' | 'CREATE' | 'DELETE';

}

export interface MenuItem {
     title: string;
     system?: boolean;
     link?: string;
     icon?: string;
     target?: 'internal' | 'external' | 'externalRedirect';
     securityConstraint?: RequireSecurityConstraintParams;
     children?: MenuItem[];
     component?: string;
     params?: object;

}

// Resource Type
export interface Menu {
    id: string;
name: string;
children: MenuItem[];
version: number;

}
// Resource and Property Names
export const MenuName = "Menu";

export const MenuIdName = "Id";

export const MenuNameName = "Name";

export const MenuChildrenName = "Children";

export const MenuVersionName = "Version";


