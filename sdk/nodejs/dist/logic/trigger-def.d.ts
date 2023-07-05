import { FunctionTrigger, Record } from "../model";
export declare function defineTrigger<T extends Record<unknown>>(functionTrigger: FunctionTrigger, fn: (entity: T) => T): void;
