import { Property } from "@apibrew/client";
import { isAnnotationEnabled } from "./annotation";

export function isSpecialProperty(property: Property): boolean {
    return isAnnotationEnabled(property.annotations as any, 'SpecialProperty');
}

export function isSimpleProperty(property: Property): boolean {
    return property.type as any === 'STRING' || property.type as any === 'INT32' || property.type as any === 'INT64' || property.type as any === 'FLOAT32' || property.type as any === 'FLOAT64';
}
