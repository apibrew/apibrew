import React from "react";
import {Resource} from "../model";

export const ResourceContext = React.createContext<Resource | undefined>(undefined)

export function useResource() {
    const resource = React.useContext(ResourceContext)
    if (!resource) {
        throw new Error("useResource must be used within a ResourceProvider")
    }
    return resource
}