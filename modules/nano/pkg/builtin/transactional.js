function transactional(fn) {
    return function (...args) {
        try {
            begin()

            const result = fn(...args)

            commit()

            return result
        } catch (e) {
            rollback()

            throw e
        }
    }
}