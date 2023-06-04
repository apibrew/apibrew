


// Sub Types

export interface MenuItem {
     title: string;
     system?: boolean;
     link?: string;
     icon?: string;
     target?: 'internal' | 'external' | 'external-redirect';
     children?: MenuItem[];

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


