import { ResourceProperty } from "../model";
import { isAnnotationEnabled } from "./annotation";

export function isSpecialProperty(property: ResourceProperty): boolean {
    return isAnnotationEnabled(property.annotations, 'SpecialProperty');
}