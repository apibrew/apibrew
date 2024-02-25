import {getAnnotation, isAnnotationEnabled} from "./annotation";
import {Property} from "../model";
import {Type} from "../model/resource";

export function isSpecialProperty(property: Property): boolean {
    return isAnnotationEnabled(property.annotations as any, 'SpecialProperty');
}

export function getPropertyOrder(name: string, property: Property): number {
    const order = parseInt(getAnnotation(property.annotations as any, 'Order', '0'));

    if (order !== 0) {
        return order
    }

    if (name === 'id') {
        return -2
    }

    if (isSpecialProperty(property)) {
        return 1000
    }

    if (hasComplexStructure(property)) {
        return 100
    }

    return 0;
}

export function hasComplexStructure(property: Property): boolean {
    return property.type === Type.LIST || property.type === Type.MAP || property.type === Type.STRUCT;
}

export function sortedProperties(properties: { [key: string]: Property }): string[] {
    const propertyNames = Object.keys(properties)

    return propertyNames.sort((a, b) => {
        const aProperty = properties[a]
        const bProperty = properties[b]

        if (!aProperty || !bProperty) {
            return 0
        }

        const aOrder = getPropertyOrder(a, aProperty)
        const bOrder = getPropertyOrder(b, bProperty)

        return aOrder - bOrder
    })
}

export function isComparableProperty(property: Property): boolean {
    if (property.type === Type.INT32 || property.type === Type.INT64 || property.type === Type.FLOAT32 || property.type === Type.FLOAT64) {
        return true
    }

    if (property.type === Type.DATE || property.type === Type.TIME || property.type === Type.TIMESTAMP) {
        return true
    }

    return false
}

export function withPropertyOrder(property: Property, order: number): Property {
    if (!property.annotations) {
        property.annotations = {}
    }

    property.annotations['Order'] = order.toString()

    return property
}

export function isSimpleProperty(property: Property): boolean {
    return property.type === 'BOOL' || property.type === 'STRING' || property.type === 'INT32' || property.type === 'INT64' || property.type === 'FLOAT32' || property.type === 'FLOAT64';
}

export function makeProperties(properties: { [key: string]: Property }) {
    let orderStart = -2;

    if (properties['id']) {
        properties['id'].annotations['Order'] = orderStart.toString()
        orderStart++;
    }

    if (properties['version']) {
        properties['version'].annotations['Order'] = orderStart.toString()
        orderStart++;
    }

    if (properties['auditData']) {
        properties['auditData'].annotations['Order'] = orderStart.toString()
        orderStart++;
    }

    return Object.entries(properties)
        .map(item => {
            return {name: item[0], property: item[1]}
        })
        .sort((a, b) => getPropertyOrder(a.name, a.property) - getPropertyOrder(a.name, b.property))
        .map((item, index) => {
            if (!item.property.annotations) {
                item.property.annotations = {}
            }

            item.property.annotations['Order'] = (index + orderStart).toString()

            return item
        })
        .sort((a, b) => getPropertyOrder(a.name, a.property) - getPropertyOrder(a.name, b.property))
}

