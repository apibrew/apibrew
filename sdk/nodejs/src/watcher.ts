import {Entity} from "./entity";

export interface Watcher<T extends Entity> {
    start(): void;
    stop(): void;
}