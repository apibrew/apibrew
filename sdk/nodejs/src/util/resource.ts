import {getAnnotation} from "./annotation";
import {Schema} from "../model";

export function ensureGivenPropertiesOrder(schema: Schema, properties: string[]): boolean {
    let updated = false

    for (let i = 0; i < properties.length; i++) {
        const property = schema.properties[properties[i]]

        const order = parseInt(getAnnotation(property.annotations as any, 'Order', '-100'));

        if (order !== i) {
            updated = true

            console.log(`Updating order of ${properties[i]} from ${order} to ${i}`)

            property.annotations = {
                ...property.annotations || {},
                Order: i.toString()
            }
        }
    }

    return updated
}