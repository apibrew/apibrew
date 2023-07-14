


export const RouterResource = {
    resource: "Router",
    namespace: "ui",
};

// Sub Types

export interface Route {
     path: string;
     component: string;
     system?: boolean;
     params?: object;
     routes?: Route[];

}

// Resource Type
export interface Router {
    id: string;
name: string;
routes: Route[];
version: number;

}
// Resource and Property Names
export const RouterName = "Router";

export const RouterIdName = "Id";

export const RouterNameName = "Name";

export const RouterRoutesName = "Routes";

export const RouterVersionName = "Version";


