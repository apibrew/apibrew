
export function uniqueOnly<T>(value: T, index: number, array: T[]): boolean {
    return array.indexOf(value) === index;
}


export function uniqueBy<T, K>(keyFunc: (elem: T) => K) {
    return (value: T, index: number, array: T[]) => {
        for (let i = 0; i < index; i++) {
            if (keyFunc(value) == keyFunc(array[i])) {
                return false
            }
        }

        return true;
    }

}