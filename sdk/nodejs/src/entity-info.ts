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
function getRestPath(resource: Resource): string {
    if (resource.annotations && (resource.annotations as any)["OpenApiRestPath"]) {
        return (resource.annotations as any)["OpenApiRestPath"];
    } else if (!resource.namespace.name || resource.namespace.name === "default") {
        return slug(resource.name);
    } else {
        return slug(resource.namespace.name + "/" + resource.name);
    }
}

// slug
function slug(name: string): string {
    return name.toLowerCase().replace(/[^a-z0-9]+/g, "-");
}
