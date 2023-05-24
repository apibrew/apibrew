export function not<T>(fn: (value: T) => boolean): (value: T) => boolean {
    return (value: T) => !fn(value);
}