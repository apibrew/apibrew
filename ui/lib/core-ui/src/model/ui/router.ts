


// Sub Types

export interface Route {
     path: string;
     component: string;
     params?: object[];
     routes?: Route[];

}

// Resource Type
export interface Router {
    version: number;
id: string;
name: string;
routes: Route[];

}
// Resource and Property Names
export const RouterName = "Router";

export const RouterVersionName = "Version";

export const RouterIdName = "Id";

export const RouterNameName = "Name";

export const RouterRoutesName = "Routes";


