class Resource {
    constructor(type) {
        const parts = type.split('/')

        if (!type) {
            throw new Error('Invalid resource type')
        }

        if (parts.length > 2) {
            throw new Error('Invalid resource type: type must have 2 parts separated by a / or just resource name', type)
        }

        this.type = type

        if (parts.length === 1) {
            this.namespace = 'default'
            this.name = parts[0]
        } else {
            this.namespace = parts[0]
            this.name = parts[1]
        }
    }

    handle(fn, order, action) {
        var orderNumber = 49
        if (order === 'after') {
            orderNumber = 101
        }
        handle({
            name: `${this.type}[${order}${action.charAt(0).toUpperCase() + action.slice(1)}]`,
            selector: {
                namespaces: [this.namespace],
                resources: [this.name],
                actions: [action.toUpperCase()],
            },
            order: orderNumber,
            sync: true,
            responds: true,
            fn: fn
        })
    }

    beforeCreate(fn) {
        this.handle(fn, 'before', 'create')
    }

    beforeUpdate(fn) {
        this.handle(fn, 'before', 'update')
    }

    beforeDelete(fn) {
        this.handle(fn, 'before', 'delete')
    }

    beforeGet(fn) {
        this.handle(fn, 'before', 'get')
    }

    beforeList(fn) {
        this.handle(fn, 'before', 'list')
    }

    afterCreate(fn) {
        this.handle(fn, 'after', 'create')
    }

    afterUpdate(fn) {
        this.handle(fn, 'after', 'update')
    }

    afterDelete(fn) {
        this.handle(fn, 'after', 'delete')
    }

    afterGet(fn) {
        this.handle(fn, 'after', 'get')
    }

    afterList(fn) {
        this.handle(fn, 'after', 'list')
    }

    list(params) {
        return list({
            type: this.type,
            ...params
        })
    }

    load(params) {
        return load({
            type: this.type,
            ...params
        })
    }

    create(record) {
        return create({
            type: this.type,
            ...record
        })
    }

    update(record) {
        return update({
            type: this.type,
            ...record
        })
    }

    delete(record) {
        return delete ({
            type: this.type,
            ...record
        })
    }

    get(id, params) {
        return this.load({
            id,
            ...params
        })
    }

    bind(toResource, action, mapFrom, mapTo) {
        handle({
            name: `${this.type}[bind${action.charAt(0).toUpperCase() + action.slice(1)}]`,
            selector: {
                namespaces: [this.namespace],
                resources: [this.name],
                actions: [action.toUpperCase()],
            },
            order: 90,
            sync: true,
            responds: true,
            finalized: true,
            replaceEvent: true,
            fn: (record, event) => {
                switch (action) {
                    case 'create':
                        return mapFrom(toResource.create(mapTo(record)))
                    case 'update':
                        return mapFrom(toResource.update(mapTo(record)))
                    case 'delete':
                        toResource.delete(mapTo(record))
                        return
                    case 'get':
                        return mapFrom(toResource.get(record.id))
                    case 'list':
                        const result = toResource.list({
                            resolveReferences: event.recordSearchParams.resolveReferences,
                            query: event.recordSearchParams.query
                        })

                        return {
                            ...event,
                            total: result.total,
                            content: result.content
                                .map(mapFrom)
                                .map(item => {
                                    return {
                                        properties: item
                                    }
                                })
                        }
                    default:
                        throw new Error('Invalid action: ', action)
                }
            }
        })
    }

    bindCreate(toResource, mapFrom, mapTo) {
        this.bind(toResource, 'create', mapFrom, mapTo)
    }

    bindUpdate(toResource, mapFrom, mapTo) {
        this.bind(toResource, 'update', mapFrom, mapTo)
    }

    bindDelete(toResource, mapFrom, mapTo) {
        this.bind(toResource, 'delete', mapFrom, mapTo)
    }

    bindGet(toResource, mapFrom, mapTo) {
        this.bind(toResource, 'get', mapFrom, mapTo)
    }

    bindList(toResource, mapFrom, mapTo) {
        this.bind(toResource, 'list', mapFrom, mapTo)
    }

    handleEach(fn, sync) {
        const registry = {}

        const register = (record) => {
            const call = fn(record)

            if (call) {
                registry[record.id] = call
            }
        }

        const unRegister = (record) => {
            if (registry[record.id]) {
                registry[record.id]()
                delete registry[record.id]
            }
        }

        handle({
            name: `${this.type}[each]`,
            selector: {
                namespaces: [this.namespace],
                resources: [this.name],
                actions: ['CREATE', 'UPDATE', 'DELETE'],
            },
            order: 101,
            sync: Boolean(sync),
            fn: (record, event) => {
                switch (event.action + '') {
                    case 'CREATE':
                        register(record)
                        break
                    case 'UPDATE':
                        unRegister(record)
                        register(record)
                        break
                    case 'DELETE':
                        unRegister(record)
                        break
                    default:
                        throw new Error(`Invalid action:xx${event.action}xx`)
                }
            }
        })

    }
}

function resource(...args) {
    if (args.length === 1) {
        return new Resource(args[0])
    } else if (args.length === 2) {
        return new Resource(args[0] + '/' + args[1])
    } else {
        throw new Error('Invalid resource type')
    }
}