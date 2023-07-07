export function isObject(item: any) {
    return (item && typeof item === 'object' && !Array.isArray(item));
}

export function isObjectModified(source: any, updated: any): boolean {
    if (JSON.stringify(source) === JSON.stringify(updated)) {
        return false
    }

    if (isObject(source) && isObject(updated)) {
        for (const key of Object.keys(updated)) {
            console.log('looking for diff: ', key)
            if (isObjectModified(source[key], updated[key])) {
                console.log('Diff found: ', source[key], updated[key])
                return true
            }
        }

        return false
    } else if (Array.isArray(source) && Array.isArray(updated)) {
        if (source.length !== updated.length) {
            console.log('Diff found: ', source.length !== updated.length)
            return true
        }

        for (let i = 0; i < source.length; i++) {
            if (isObjectModified(source[i], updated[i])) {
                console.log('Diff found: ', source[i], updated[i])
                return true
            }
        }

        return false
    } else {
        console.log('Diff found{e}: ', source, updated)
        return true
    }
}
