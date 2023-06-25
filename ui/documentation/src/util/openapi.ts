import { OpenAPIV3_1 } from "openapi-types";

function walkSearchProperty(obj: any, prop: string): any[] {
    if (typeof obj !== 'object') {
        return []
    }

    let result: any[] = []
    for (const key in obj) {
        if (key == prop) {
            console.log(obj, key)
            result.push(obj[key])
        } else if (typeof obj[key] === 'object') {
            result = [...result, ...walkSearchProperty(obj[key], prop)]
        }
    }

    return result
}

export function relatedSchemas(doc: OpenAPIV3_1.Document, schemaNames: string[]): string[] {
    const schemas: string[] = []

    let searchRemeaning = [...schemaNames]

    while (searchRemeaning.length > 0) {
        const newItems: string[] = []
        for (let name of [...searchRemeaning]) {
            const walkResult = walkSearchProperty(doc.components?.schemas![name]!, "$ref")
            for (const item of walkResult) {
                const [name, _] = resolve(doc, item.$ref, item)

                if (newItems.indexOf(name) === -1 && schemas.indexOf(name) === -1) {
                    newItems.push(name)
                    schemas.push(name)
                }
            }
        }

        searchRemeaning = newItems
    }

    return schemas;
}

export function resolve(doc: OpenAPIV3_1.Document, schema: OpenAPIV3_1.SchemaObject | OpenAPIV3_1.ReferenceObject, ref: string): [string, OpenAPIV3_1.SchemaObject] {
    if (ref.startsWith('#')) {
        let parent: any = doc
        let name = 'root'

        ref.substring(2).split('/').forEach(item => {
            if (parent) {
                parent = parent[item]
                name = item
            } else {
                throw new Error(`Ref part not found: ${ref} => ${item}`)
            }
        })

        if (!parent) {
            throw new Error('Unresolvable ref: ' + ref)
        }

        return [name, parent]
    } else {
        throw new Error('Unresolvable ref: ' + ref)
    }
}

export function resolveSchema(doc: OpenAPIV3_1.Document, schema: OpenAPIV3_1.SchemaObject | OpenAPIV3_1.ReferenceObject): OpenAPIV3_1.SchemaObject {
    const ref = (schema as OpenAPIV3_1.ReferenceObject).$ref

    if (ref) {
        const [_, resolved] = resolve(doc, schema, ref)
        console.log('resolved', resolved)
        return resolved
    } else {
        return schema as OpenAPIV3_1.SchemaObject
    }
}

export function preparePropertyList(doc: OpenAPIV3_1.Document, schema: OpenAPIV3_1.SchemaObject | OpenAPIV3_1.ReferenceObject): { name: string, schema: OpenAPIV3_1.SchemaObject | OpenAPIV3_1.ReferenceObject }[] {
    const resolvedSchema = resolveSchema(doc, schema)

    if (resolvedSchema.properties) {
        return Object.keys(resolvedSchema.properties).map(item => {
            return {
                name: item,
                schema: resolvedSchema.properties![item]
            }
        })
    } else {
        return []
    }
}