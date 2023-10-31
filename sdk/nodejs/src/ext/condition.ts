import {ExtensionInfo} from "../extension-info";
import {Entity} from "../entity";
import {Action, Event} from "../model/extension";
import {EntityInfo} from "../entity-info";

export interface Condition<T extends Entity> {
    configureExtensionInfo(extensionInfo: ExtensionInfo): ExtensionInfo;

    eventMatches(event: Event, entity: T): boolean;
}

export class SimpleCondition<T extends Entity> implements Condition<T> {
    constructor(private readonly configureExtensionInfoFn: (extensionInfo: ExtensionInfo) => ExtensionInfo,
                private readonly eventMatchesFn: (event: Event, entity: T) => boolean) {
    }

    configureExtensionInfo(extensionInfo: ExtensionInfo): ExtensionInfo {
        return this.configureExtensionInfoFn(extensionInfo);
    }

    eventMatches(event: Event, entity: T): boolean {
        return this.eventMatchesFn(event, entity);
    }
}

export function and<T extends Entity>(condition1: Condition<T>, condition2: Condition<T>): Condition<T> {
    return new SimpleCondition<T>(ei => condition2.configureExtensionInfo(condition1.configureExtensionInfo(ei)),
        (e, t) => condition1.eventMatches(e, t) && condition2.eventMatches(e, t));
}

export function or<T extends Entity>(condition1: Condition<T>, condition2: Condition<T>): Condition<T> {
    return new SimpleCondition<T>(ei => condition2.configureExtensionInfo(condition1.configureExtensionInfo(ei)),
        (e, t) => condition1.eventMatches(e, t) || condition2.eventMatches(e, t));
}

export function before<T extends Entity>(): Condition<T> {
    return new SimpleCondition<T>(ei => {
        return {...ei, order: 10}
    }, (e, t) => true);
}

export function after<T extends Entity>(): Condition<T> {
    return new SimpleCondition<T>(ei => {
        return {...ei, order: 110}
    }, (e, t) => true);
}

export function on<T extends Entity>(customActionName: string): Condition<T> {
    return new SimpleCondition<T>(ei => {
        return {...ei, action: Action.OPERATE}
    }, (e, t) => e.actionName === customActionName);
}

export function create<T extends Entity>(): Condition<T> {
    return new SimpleCondition<T>(ei => {
        return {...ei, action: Action.CREATE}
    }, (e, t) => true);
}

export function update<T extends Entity>(): Condition<T> {
    return new SimpleCondition<T>(ei => {
        return {...ei, action: Action.UPDATE}
    }, (e, t) => true);
}

export function delete$<T extends Entity>(): Condition<T> {
    return new SimpleCondition<T>(ei => {
        return {...ei, action: Action.DELETE}
    }, (e, t) => true);
}

export function get<T extends Entity>(): Condition<T> {
    return new SimpleCondition<T>(ei => {
        return {...ei, action: Action.GET}
    }, (e, t) => true);
}

export function list<T extends Entity>(): Condition<T> {
    return new SimpleCondition<T>(ei => {
        return {...ei, action: Action.LIST}
    }, (e, t) => true);
}

export function beforeCreate<T extends Entity>(): Condition<T> {
    return and(before(), create());
}

export function beforeUpdate<T extends Entity>(): Condition<T> {
    return and(before(), update());
}

export function beforeDelete<T extends Entity>(): Condition<T> {
    return and(before(), delete$());
}

export function beforeGet<T extends Entity>(): Condition<T> {
    return and(before(), get());
}

export function beforeList<T extends Entity>(): Condition<T> {
    return and(before(), list());
}

export function afterCreate<T extends Entity>(): Condition<T> {
    return and(after(), create());
}

export function afterUpdate<T extends Entity>(): Condition<T> {
    return and(after(), update());
}

export function afterList<T extends Entity>(): Condition<T> {
    return and(after(), list());
}

export function afterDelete<T extends Entity>(): Condition<T> {
    return and(after(), delete$());
}

export function afterGet<T extends Entity>(): Condition<T> {
    return and(after(), get());
}

export function onAction<T extends Entity>(customActionName: string): Condition<T> {
    return and(before(), on(customActionName));
}

export function async<T extends Entity>(): Condition<T> {
    return new SimpleCondition<T>(ei => {
        return {...ei, sync: false}
    }, (e, t) => true);
}

export function entityExists<T extends Entity>(): Condition<T> {
    return new SimpleCondition<T>(ei => ei, (e, t) => t != null);
}

export function user<T extends Entity>(...expectedUsers: string[]): Condition<T> {
    return new SimpleCondition<T>(ei => ei, (e, t) => {
        const user = e.annotations["user"];
        if (user == null) {
            return false;
        }

        for (const eu of expectedUsers) {
            if (user === eu) {
                return true;
            }
        }

        return false;

    });
}

export function resource<T extends Entity>(namespace: string, ...resource: string[]): Condition<T> {
    return new SimpleCondition<T>(ei => {
        return {...ei, namespace: namespace, resources: resource}
    }, (e, t) => true);
}

export function resourceFromEntityInfo<T extends Entity>(entityInfo: EntityInfo): Condition<T> {
    return new SimpleCondition<T>(ei => {
        return {...ei, namespace: entityInfo.namespace, resources: [entityInfo.resource]}
    }, (e, t) => true);
}

export function group<T extends Entity>(...expectedGroups: string[]): Condition<T> {
    return new SimpleCondition<T>(ei => ei, (e, r) => {
        const groups = e.annotations["groups"];
        if (groups == null) {
            return false;
        }

        for (const g of groups.split(",")) {
            for (const eg of expectedGroups) {
                if (g === eg) {
                    return true;
                }
            }
        }

        return false;
    });
}

export const Condition = {
    and,
    or,
    before,
    after,
    on,
    create,
    update,
    delete: delete$,
    get,
    list,
    beforeCreate,
    beforeUpdate,
    beforeDelete,
    beforeGet,
    beforeList,
    afterCreate,
    afterUpdate,
    afterList,
    afterDelete,
    afterGet,
    onAction,
    async,
    entityExists,
    user,
    resource,
    resourceFromEntityInfo,
    group,
}
