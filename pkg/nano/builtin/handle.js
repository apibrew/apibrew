const selectorRe = /^((([$a-zA-Z][$a-zA-Z0-9_-]*)\/)?([*$a-zA-Z][$a-zA-Z0-9_-]*))?(\[([$a-zA-Z][$a-zA-Z0-9_-]*)(=([$a-zA-Z][$a-zA-Z0-9_-]*))?])?$/
const orderRe = /^(before|after|on\((\d+)\))?(create|update|delete|get|list)?$/

function parseSelector(str) {
    if (!selectorRe.test(str)) {
        throw new Error("Selector does not match regex: " + selectorRe + " => " + str)
    }

    const match = selectorRe.exec(str)

    console.log('parseSelector match', match)

    const selector = {}

    const namespace = match[3]
    const resource = match[4]
    const annotationKey =  match[6]
    const annotationValue =  match[8]

    if (resource) {
        selector.resources = [resource]
    }

    if (namespace && namespace !== '*') {
        selector.namespaces = [namespace]
    }

    if (annotationKey) {
        if (annotationValue) {
            selector.annotations = {
                [annotationKey]: annotationValue
            }
        } else {
            selector.annotations = {
                [annotationKey]: "*"
            }
        }
    }

    return selector
}

function parseOrder(handleConfig, str) {
    if (!orderRe.test(str.toLowerCase())) {
        throw new Error("Order does not match regex: " + orderRe + " => " + str)
    }

    const match = orderRe.exec(str.toLowerCase())

    console.log('parseOrder match', match)

    const order = match[1]
    const on = match[2]
    const action = match[3]

    if (order) {
        if (order === 'before') {
            handleConfig.order = 99
        } else if (order === 'after') {
            handleConfig.order = 101
        } else {
            handleConfig.order = parseInt(on)
        }
    }

    if (action) {
        handleConfig.selector.actions = [action.toUpperCase()]
    }
}

function on(exp, fn) {
    if (!exp) {
        throw new Error("Expression must be provided")
    }
    if (!fn) {
        throw new Error("Function")
    }
    const parts = exp.split(":")

    if (parts.length < 2 || parts.length > 3) {
        throw new Error("Expression must have 2 or 3 part separated with: (example: system/User:beforeCreate)")
    }

    const handleConfig = {}

    // parsing selector
    handleConfig.selector = parseSelector(parts[0])

    parseOrder(handleConfig, parts[1])

    console.log("Registering Handler: ", handleConfig)

    handle({
        name: `on[${exp}]`,
        ...handleConfig,
        sync: true,
        responds: true,
        fn: fn
    })
}