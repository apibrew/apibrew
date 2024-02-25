export type Annotations = {
    [key: string]: string;
}

export const isAnnotationEnabled = (annotations: Annotations | undefined, annotation: string): boolean => {
    return (annotations?.[annotation]) === 'true';
}

export const getAnnotation = (annotations: Annotations | undefined, annotation: string, defaultValue?: string): string => {
    return annotations?.[annotation] ?? defaultValue ?? '';
}

export const withBooleanAnnotation = (annotations: Annotations | undefined, annotation: string, value: boolean): Annotations => {
    const newAnnotations: Annotations = annotations ? {...annotations} : {};

    if (value) {
        newAnnotations[annotation] = "true";
    } else {
        delete (newAnnotations[annotation])
    }

    return newAnnotations;
}
