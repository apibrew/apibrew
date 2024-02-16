const dynamicResource = function (params, handler) {
    const Resource = resource('system/Resource')

    const resourceMap = {}

    const matchResource = res => {
        if (res.annotations && params.annotations) {
            for (let key in params.annotations) {
                if (res.annotations[key] === params.annotations[key]) {
                    return true
                }

                if (res.annotations[key] && params[key] === '*') {
                    return true
                }
            }
        }

        if (params.matches) {
            if (params.matches(res)) {
                return true
            }
        }
        return false
    }

    const registerResourceHandler = res => {
        console.log('registerResourceHandler', res.name)
        const ItemResource = resource(`${res.namespace.name}/${res.name}`)

        resourceMap[res.id] = ItemResource
        handler(ItemResource)
    }

    const unregisterResourceHandler = res => {
        if (!resourceMap[res.id]) {
            return
        }
        console.log('unregisterResourceHandler', res.name)

        resourceMap[res.id].unregisterAllHandlers()
        delete resourceMap[res.id]
    }

    Resource.beforeCreate(res => {
        registerResourceHandler(res)
    })

    Resource.beforeUpdate(res => {
        if (resourceMap[res.id] && !matchResource(res)) {
            unregisterResourceHandler(res)
        }

        if (!resourceMap[res.id] && matchResource(res)) {
            registerResourceHandler(res)
        }
    })

    Resource.beforeDelete(res => {
        unregisterResourceHandler(res)
    })

    const currentResources = Resource.list().content

    for (res of currentResources) {
        if (matchResource(res)) {
            registerResourceHandler(res)
        }
    }
}