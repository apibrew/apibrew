import { ResourceProperty } from "@apibrew/client";
import {isAnnotationEnabled} from "./annotation";

export function isSpecialProperty(property: ResourceProperty): boolean {
    return isAnnotationEnabled(property.annotations as any, 'SpecialProperty');
}

export function isSimpleProperty(property: ResourceProperty): boolean {
    return property.type === 'STRING' || property.type === 'INT32' || property.type === 'INT64' || property.type === 'FLOAT32' || property.type === 'FLOAT64';
}
