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


        this.initDynamicMethods()
    }

    handle(fn, order, action, sync = true) {
        var orderNumber = 49
        if (order === 'after') {
            orderNumber = 101
        }

        let actions = [action.toUpperCase()]

        if (action === 'read') {
            actions = ['GET', 'LIST']
        }

        const name = `${this.type}[${order}${action.charAt(0).toUpperCase() + action.slice(1)}]`

        console.log('registering handler with name: ' + name, actions, orderNumber, sync, fn)

        handle({
            name: name,
            selector: {
                namespaces: [this.namespace],
                resources: [this.name],
                actions: actions,
            },
            order: orderNumber,
            sync: sync,
            responds: true,
            fn: fn
        })
    }

    initDynamicMethods() {
        const orders = ['before', 'after']
        const actions = ['create', 'update', 'delete', 'get', 'list', 'read']

        for (let order of orders) {
            for (let action of actions) {
                this[order + capitalizeFirstLetter(action)] = (fn) => {
                    this.handle(fn, order, action, true)
                }

                this[order + capitalizeFirstLetter(action) + 'Async'] = (fn) => {
                    this.handle(fn, order, action, false)
                }
            }
        }
    }

    on(fn) {
        console.log('registering on handler', fn)
        this.beforeCreate(fn)
    }

    modifier(fn) {
        this.postModifier(fn)
    }

    preModifier(fn) {
        this.beforeRead((record, event) => {
            const result = fn(record, event)

            if (result) {
                return result
            }

            return record
        })
    }

    postModifier(fn) {
        console.log('this.afterRead', this.afterRead)
        this.afterRead((record, event) => {
            if (!record) {
                return
            }
            const result = fn(record, event)

            if (result) {
                return result
            }

            return record
        })
    }

    list(params) {
        return list({
            type: this.type,
            ...params
        })
    }

    find(params) {
        return this.load(params)
    }

    findBy(property, value) {
        return this.find({[property]: value})
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

    apply(record) {
        return apply({
            type: this.type,
            ...record
        })
    }

    delete(record) {
        return delete_({
            type: this.type,
            ...record
        })
    }

    deleteAll() {
        const records = this.list({
            limit: 10000
        })

        records.content.forEach(record => {
            this.delete(record)
        })
    }

    get(id, params) {
        return this.load({
            id,
            ...params
        })
    }

    count(filters) {
        const result = this.list({
            filters: filters,
            limit: 0
        })

        return result.total
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

    each(fn, params) {
        this.list({limit: 1000, ...params}).content.forEach(fn)
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

        setTimeout(() => {
            const listResponse = this.list({
                limit: 1000
            })

            if (!listResponse) {
                return
            }

            listResponse.content.forEach(register)
        }, 0)
    }
}

function resource(...args) {
    if (args.length === 1) {
        if (typeof args[0] === 'object' && args[0] !== null) {
            const resObj = args[0]
            let type = ''

            if (resObj.namespace && resObj.namespace.name && resObj.namespace.name !== 'default') {
                type = resObj.namespace.name + '/' + resObj.name
            } else {
                type = resObj.name
            }

            apply({
                type: 'system/Resource',
                ...resObj
            })

            return new Resource(type)
        } else {
            return new Resource(args[0])
        }
    } else if (args.length === 2) {
        return new Resource(args[0] + '/' + args[1])
    } else {
        throw new Error('Invalid resource type')
    }
}

function capitalizeFirstLetter(str) {
    if (typeof str !== 'string' || str.length === 0) {
        return '';
    }
    return str.charAt(0).toUpperCase() + str.slice(1);
}
