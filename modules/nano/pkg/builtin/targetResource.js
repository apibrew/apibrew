class TargetResource {
    constructor(targetConfig, type) {
        this.targetConfig = targetConfig

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

        let path = resourcePath(this.namespace, this.name)

        if (this.targetConfig.resourcePath) {
            path = this.targetConfig.resourcePath
        }

        if (!targetConfig.baseUrl.endsWith('/')) {
            targetConfig.baseUrl += '/'
        }

        this.resourceBaseUrl = targetConfig.baseUrl + path
        this.resourcePath = path
    }

    httpOptions() {
        return {
            headers: {
                'Authorization': 'Bearer ' + this.targetConfig.token
            }
        }
    }

    body(httpResponse) {
        if (httpResponse.statusCode !== 200) {
            throw new Error("Remote error[" + httpResponse.statusCode + "]" + httpResponse.body.text())
        } else {
            return httpResponse.body.json()
        }
    }

    list(params) {
        if (this.resourcePath === 'resources') {
            return this.body(http.get(this.resourceBaseUrl, this.httpOptions()))
        }

        const url = this.resourceBaseUrl + '/_search'

        return this.body(http.post(url, {
            ...params,
        }, this.httpOptions()))
    }

    load(record) {
        const url = this.resourceBaseUrl + '/_load'
        return this.body(http.post(url, record, this.httpOptions()))
    }

    create(record) {
        const url = this.resourceBaseUrl
        return this.body(http.post(url, record, this.httpOptions()))
    }

    update(record) {
        const url = this.resourceBaseUrl + '/' + record.id
        return this.body(http.put(url, record, this.httpOptions()))
    }

    apply(record) {
        const url = this.resourceBaseUrl
        return this.body(http.put(url, record, this.httpOptions()))
    }

    delete(record) {
        const url = this.resourceBaseUrl + '/' + record.id
        return this.body(http.delete(url, this.httpOptions()))
    }

    get(id, params) {
        let url = this.resourceBaseUrl + '/' + id

        if (params && params.resolveReferences) {
            url += '?resolve-references=' + params.resolveReferences.join(',')
        }

        return this.body(http.get(url, this.httpOptions()))
    }

    count(filters) {
        const result = this.list({
            filters: filters,
            limit: 1
        })

        return result.total
    }
}

function targetResource(targetConfig, ...args) {
    if (args.length === 1) {
        return new TargetResource(targetConfig, args[0])
    } else if (args.length === 2) {
        return new TargetResource(targetConfig, args[0] + '/' + args[1])
    } else {
        throw new Error('Invalid resource type')
    }
}
