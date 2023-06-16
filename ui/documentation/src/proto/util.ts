import {GetProtoImage} from "./proto.ts";
import {EnumType, Field, File, MessageType, ProtoElement, ProtoElementKind} from "./image";

export type ProtoFileKind = 'stub' | 'model' | 'ext'

export function getProtoDescriptor(kind: ProtoFileKind, name: string): File {
    return getProtoDescriptorByPath(`${kind}/${name}.proto`)
}

export function getProtoDescriptorByPath(path: string): File {
    const descriptor = GetProtoImage().file.find(item => {
        return item.name == path
    })

    if (!descriptor) {
        throw new Error(`Unable to find descriptor for ${path}`)
    }

    return descriptor
}

export function getMessageTypeDescriptor(file: File, messageName: string): MessageType {
    const messageTypes = file.messageType || []

    const descriptor = messageTypes.find(item => {
        return item.name == messageName
    })

    if (!descriptor) {
        throw new Error(`Unable to find descriptor for message ${messageName}`)
    }

    return descriptor
}

export function locateElementByName(kind: ProtoElementKind, file: File, name: string): ProtoElement | undefined {
    switch (kind) {
        case 'message':
            const messageTypes = file.messageType || []

            const descriptor = messageTypes.find(item => {
                return item.name == name
            })

            if (descriptor) {
                return descriptor
            }
            break
        case 'enum':
            const enumTypes = file.enumType || []

            const enumDescriptor = enumTypes.find(item => {
                return item.name == name
            })

            if (enumDescriptor) {
                return enumDescriptor
            }
            break
    }
}

export function locateSubElementByName(kind: ProtoElementKind, parentElement: MessageType, name: string): ProtoElement | undefined {
    switch (kind) {
        case 'message':
            const messageTypes = parentElement.nestedType || []

            const descriptor = messageTypes.find(item => {
                return item.name == name
            })

            if (descriptor) {
                return descriptor
            }
            break
        case 'enum':
            const enumTypes = parentElement.enumType || []

            const enumDescriptor = enumTypes.find(item => {
                return item.name == name
            })

            if (enumDescriptor) {
                return enumDescriptor
            }
            break
    }
}

export function locateByTypeName(kind: ProtoElementKind, files: File[], typeName: string): {
    element: ProtoElement
    file: File,
} {
    console.log('locateByTypeName', kind, typeName)
    console.log(files)
    const typeNameParts = typeName.substring(1).split('.')

    mainLoop:
        for (const file of files) {
            if (typeNameParts[0] != file.package) {
                console.log('step1')
                continue
            }

            const localTypeNameParts = typeNameParts.slice(1)

            if (localTypeNameParts.length == 1) {
                let parentElement = locateElementByName(kind, file, localTypeNameParts[0])

                if (!parentElement) {
                    console.log('step2', localTypeNameParts[0])
                    continue
                }

                console.log('returning parentElement', parentElement)
                return {
                    element: parentElement,
                    file: file,
                }
            } else {
                let parentElement = locateElementByName('message', file, localTypeNameParts[0])

                if (!parentElement) {
                    console.log('step2', localTypeNameParts[0])
                    continue
                }

                for (let i = 1; i < localTypeNameParts.length; i++) {
                    const isLast = i == localTypeNameParts.length - 1

                    console.log(i, parentElement, localTypeNameParts[i])
                    parentElement = locateSubElementByName(isLast ? kind : 'message', parentElement as MessageType, localTypeNameParts[i])

                    console.log(i, parentElement)

                    if (!parentElement) {
                        console.log('step3')
                        continue mainLoop
                    }
                }

                console.log('returning parentElement', parentElement)
                return {
                    element: parentElement,
                    file: file,
                }
            }
        }

    throw new Error(`Unable to find descriptor for ${kind} ${typeName}`)
}

export function resolveByTypeName(kind: ProtoElementKind, file: File, typeName: string): {
    element: ProtoElement
    file: File,
    special?: boolean,
} {
    // special types
    if (typeName.startsWith('.google') || typeName.startsWith('.gnostic')) {
        return {
            element: {
                name: typeName,
            },
            special: true,
            file: file,
        }
    }


    // locate dependencies
    const dependencies = (file.dependency || []).filter(item => {
        if (item.startsWith('google/')) {
            return false
        }

        if (item.startsWith('gnostic/')) {
            return false
        }

        return true
    }).map(item => getProtoDescriptorByPath(item))

    return locateByTypeName(kind, [file, ...dependencies], typeName)
}

export function translateType(protoFile: File, field: Field): string {
    let type = field.type as string

    if (field.typeName) {
        if (field.type == 'TYPE_MESSAGE') {
            const resolvedType = resolveByTypeName('message', protoFile, field.typeName)
            const descriptor = resolvedType.element
            type = descriptor.name

            if (resolvedType.special) {
                switch (type) {
                    case '.google.protobuf.Timestamp':
                        type = 'Timestamp'
                        break
                    case '.google.protobuf.Duration':
                        type = 'Duration'
                        break
                    case '.google.protobuf.Value':
                        type = 'Value'
                        break
                    case '.google.protobuf.Struct':
                        type = 'Struct'
                        break
                    case '.google.protobuf.Any':
                        type = 'Any'
                        break
                    default:
                        throw new Error(`Unable to resolve special type ${type}`)
                }
            }

        } else if (field.type == 'TYPE_ENUM') {
            const descriptor = resolveByTypeName('enum', protoFile, field.typeName).element as EnumType
            type = descriptor.name

            type = `${type} [${descriptor.value.map(item => item.name).join(', ')}]`
        } else {
            throw new Error(`Unable to resolve type ${field.typeName}`)
        }
    } else {
        type = type.toLowerCase().substring(5)
    }

    if (field.label == 'optional') {
        return `${type}?`
    }

    if (field.label == 'LABEL_REPEATED') {
        return `${type}[]`
    }

    return type
}

export function locateDescription(protoFile: File, path: number[]) {
    console.log(path)
    for (const location of protoFile.sourceCodeInfo.location || []) {
        if (location.path && location.path.join(',') == path.join(',')) {
            let description = ''

            if (location.leadingComments) {
                description += location.leadingComments
            }

            if (location.trailingComments) {
                description += '\n' + location.trailingComments
            }

            return description
        }
    }

    return ''
}
