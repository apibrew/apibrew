import { Property } from "@apibrew/client";
import React from "react";

export const ResourcePropertyContext = React.createContext<Property | undefined>(undefined)

export function useResourceProperty(required: boolean) {
    const Property = React.useContext(ResourcePropertyContext)

    if (required && !Property) {
        throw new Error("useResourceProperty must be used within a ResourcePropertyProvider")
    }
    return Property
}