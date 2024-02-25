import {Property} from "./resource";

export interface Schema {
    name: string;
    properties: {
        [key: string]: Property
    }
}
