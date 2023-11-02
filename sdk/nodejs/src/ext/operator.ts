import {Handler} from "./handler";
import {Entity} from "../entity";
import {beforeCreate, beforeDelete, beforeList, beforeUpdate, afterGet, and} from "./condition";
import {Code, Event} from "../model/extension";
import {ApiException} from "../api-exception";
import {BooleanExpressionBuilder} from "../boolean-expression-builder";

export interface Operator<T extends Entity> {
    operate(handler: Handler<T>): string;
}

export function graceFullDelete<T extends Entity>(property: string, value: any): Operator<T> {
    return {
        operate(handler: Handler<T>): string {
            handler = handler.configure(ei => {
                return {
                    ...ei,
                    responds: true,
                }
            });
            handler.when(beforeDelete()).localOperator(graceFullDeleteOperator(property, value));
            handler.when(afterGet()).localOperator(graceFullDeleteOperator(property, value));
            handler.when(beforeCreate()).localOperator(graceFullDeleteOperator(property, value));
            handler.when(beforeUpdate()).localOperator(graceFullDeleteOperator(property, value));

            return handler.when(beforeList()).localOperator((event, entity) => {
                const query = event.recordSearchParams.query;

                const deletedFilterExp = BooleanExpressionBuilder.not(
                    BooleanExpressionBuilder.eq(property, value)
                );

                if (query == null) {
                    event.recordSearchParams.query = deletedFilterExp
                } else {
                    event.recordSearchParams.query = BooleanExpressionBuilder.and(query, deletedFilterExp)
                }

                return entity;
            });

        }
    };
}

export function graceFullDeleteOperator<T extends Entity>(property: string, value: any) {
    return (event: Event, entity: T) => {
        if ((entity as any)[property] === value) {
            throw new ApiException(Code.RECORD_VALIDATION_ERROR, "Entity is already deleted");
        }

        return entity;
    };
}

export function dataSeparation<T extends Entity>(property: string, ownerFieldGetter: (entity: T) => string): Operator<T> {
    const dataSeparator = (event: Event, entity: T) => {
        if ((entity as any)[property] === event.annotations.user) {
            return entity;
        } else {
            throw new ApiException(Code.RECORD_VALIDATION_ERROR, "Entity is not owned by user");
        }
    };

    return {
        operate(handler: Handler<T>): string {
            handler = handler.configure(ei => {
                return {
                    ...ei,
                    responds: true,
                }
            });
            handler.when(beforeCreate()).localOperator(dataSeparator);
            handler.when(beforeUpdate()).localOperator(dataSeparator);
            handler.when(beforeDelete()).localOperator(dataSeparator); // fix delete
            handler.when(afterGet()).localOperator(dataSeparator);

            return handler.when(beforeList()).localOperator((event, entity) => {
                const query = event.recordSearchParams.query;

                const ownerFilterExp = BooleanExpressionBuilder.eq(property, event.annotations.user);

                if (query == null) {
                    event.recordSearchParams.query = ownerFilterExp
                } else {
                    event.recordSearchParams.query = BooleanExpressionBuilder.and(query, ownerFilterExp)
                }

                return entity;
            });

        }
    };
}


export function execute<T extends Entity>(consumer: (event: Event, entity: T) => T | void): Operator<T> {
    return {
        operate(handler: Handler<T>): string {
            return handler.localOperator((event, entity) => {
                consumer(event, entity);

                return entity;
            });
        }
    };
}

export function reject<T extends Entity>(message?: string): Operator<T> {
    return check((event, entity) => false, message);
}

export function check<T extends Entity>(condition: (event: Event, entity: T) => boolean, message?: string): Operator<T> {
    return {
        operate(handler: Handler<T>): string {
            return handler.localOperator((event, entity) => {
                if (condition(event, entity)) {
                    return entity;
                } else {
                    throw new ApiException(Code.RECORD_VALIDATION_ERROR, message || "Condition not met");
                }
            });
        }
    };
}
