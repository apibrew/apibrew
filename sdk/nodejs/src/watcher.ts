import {Entity} from "./entity";

export interface Watcher<T extends Entity> {
    start(): void;
    run(): Promise<void>;
    stop(): void;
}