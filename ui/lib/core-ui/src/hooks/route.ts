import {useLocation} from "react-router-dom";
import {useContext} from "react";
import {RouteContext} from "../context/route-context.ts";

export function useRoute() {
    const location = useLocation();
    let routePath = useContext(RouteContext)

    if (!routePath) {
        throw new Error("RouteContext is not defined. Please check if you are using the RouteContext.Provider");
    }

    if (routePath.endsWith('/*')) {
        routePath = routePath.substring(0, routePath.length - 1)
    }

    const subPath = location.pathname.substring(routePath.length);

    return {
        fullPath: location.pathname,
        path: routePath,
        subPath: subPath,
    }
}