import {Entity} from "./entity";

export interface Container<T extends Entity> {
    content: T[]
    total: number
}