import {ResourceProperty} from "../model";
import {isAnnotationEnabled} from "./annotation";

export function isSpecialProperty(property: ResourceProperty): boolean {
    return isAnnotationEnabled(property.annotations, 'SpecialProperty');
}

export function isSimpleProperty(property: ResourceProperty): boolean {
    return property.type === 'STRING' || property.type === 'INT32' || property.type === 'INT64' || property.type === 'FLOAT32' || property.type === 'FLOAT64';
}
