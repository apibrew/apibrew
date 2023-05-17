export type Annotations = {
    [key: string]: string | undefined;
}

export const isAnnotationEnabled = (annotations: Annotations | undefined, annotation: string): boolean => {
    return (annotations?.[annotation]) === 'true';
}