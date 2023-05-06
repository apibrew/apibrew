
type Annotations = {
    [key: string]: string | undefined;
} | undefined

export function isAnnotationEnabled(annotations: Annotations, annotationName: string): boolean {
    return Boolean(annotations && annotations[annotationName] === 'true')
}

export const SpecialProperty = 'SpecialProperty'