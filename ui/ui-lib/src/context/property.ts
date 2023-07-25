import React from "react";
import { Resource } from "../model/system/resource";
import { ResourceProperty } from "../model/system/resource-property";

export const ResourcePropertyContext = React.createContext<ResourceProperty | undefined>(undefined)

export function useResourceProperty(required: boolean) {
    const ResourceProperty = React.useContext(ResourcePropertyContext)

    if (required && !ResourceProperty) {
        throw new Error("useResourceProperty must be used within a ResourcePropertyProvider")
    }
    return ResourceProperty
}