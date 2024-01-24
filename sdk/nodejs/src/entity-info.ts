import {Resource} from "./model/resource";

export interface EntityInfo {
    namespace: string;
    resource: string;
    restPath: string;
}

// fromResource
export function fromResource(resource: Resource): EntityInfo {
    return {
        namespace: resource.namespace.name,
        resource: resource.name,
        restPath: getRestPath(resource)
    };
}

// getRestPath
export function getRestPath(resource: Resource): string {
    if (resource.annotations && (resource.annotations as any)["OpenApiRestPath"]) {
        return (resource.annotations as any)["OpenApiRestPath"];
    } else if (!resource.namespace.name || resource.namespace.name === "default") {
        return slug(resource.name);
    } else {
        return slug(resource.namespace.name) + "-" + slug(resource.name);
    }
}

// slug
function slug(s: string): string {
    let result = '';
    for (let i = 0; i < s.length; i++) {
        const char = s[i];
        if (i > 0 && char !== char.toLowerCase()) {
            result += '-';
        }
        result += char.toLowerCase();
    }
    return result;
}
