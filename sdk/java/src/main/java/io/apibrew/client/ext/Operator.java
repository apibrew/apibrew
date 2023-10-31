package io.apibrew.client.ext;


import io.apibrew.client.ApiException;
import io.apibrew.client.Entity;
import io.apibrew.client.model.Extension;

import java.util.List;
import java.util.Objects;
import java.util.function.*;

import static io.apibrew.client.ext.Condition.*;

public interface Operator<T extends Entity> {

    String operate(Handler<T> handler);

    static <T extends Entity> Operator<T> graceFullDelete(String property, Object value) {

        return handler -> {
            handler = handler.configure(ei -> ei.withResponds(true));
            handler.when(beforeDelete()).operate(graceFullDeleteOperator(property, value));
            handler.when(afterGet()).operate(graceFullDeleteOperator(property, value));
            handler.when(beforeCreate()).operate(graceFullDeleteOperator(property, value));
            handler.when(beforeUpdate()).operate(graceFullDeleteOperator(property, value));

            return handler.when(beforeList()).operate((event, entity) -> {
                Extension.BooleanExpression query = event.getRecordSearchParams().getQuery();

                Extension.BooleanExpression deletedFilterExp = new Extension.BooleanExpression().withNot(
                        new Extension.BooleanExpression().withEqual(
                                new Extension.PairExpression()
                                        .withLeft(new Extension.Expression().withProperty(property))
                                        .withRight(new Extension.Expression().withValue(value))
                        )
                );

                if (query == null) {
                    event.getRecordSearchParams().setQuery(deletedFilterExp);
                } else {
                    event.getRecordSearchParams().setQuery(new Extension.BooleanExpression().withAnd(List.of(query, deletedFilterExp)));
                }

                return entity;
            });
        };
    }

    private static <T extends Entity> BiFunction<Extension.Event, T, T> graceFullDeleteOperator(String property, Object value) {
        return (event, entity) -> {
            if (entity == null) {
                return null;
            }
            if (Objects.equals(entity.getProperty(property), value)) {
                throw new ApiException(Extension.Code.RECORD_VALIDATION_ERROR, "Entity is already deleted");
            } else {
                return entity;
            }
        };
    }

    static <T extends Entity> Operator<T> dataSeparation(String property, Function<T, String> ownerFieldGetter) {
        BiFunction<Extension.Event, T, T> dataSeparator = (event, entity) -> {
            if (entity == null) {
                return null;
            }
            if (Objects.equals(ownerFieldGetter.apply(entity), event.getAnnotations().get("user")))
                return entity;
            else
                throw new ApiException(Extension.Code.RECORD_VALIDATION_ERROR, "Entity is not owned by user");
        };

        return handler -> {
            handler = handler.configure(ei -> ei.withResponds(true));
            handler.when(beforeCreate()).operate(dataSeparator);
            handler.when(beforeUpdate()).operate(dataSeparator);
            handler.when(beforeDelete()).operate(dataSeparator); // fix delete
            handler.when(afterGet()).operate(dataSeparator);

            return handler.when(beforeList()).operate((event, entity) -> {
                Extension.BooleanExpression query = event.getRecordSearchParams().getQuery();

                Extension.BooleanExpression ownerFilterExp = new Extension.BooleanExpression().withEqual(
                        new Extension.PairExpression()
                                .withLeft(new Extension.Expression().withProperty(property))
                                .withRight(new Extension.Expression().withValue(event.getAnnotations().get("user")))
                );

                if (query == null) {
                    event.getRecordSearchParams().setQuery(ownerFilterExp);
                } else {
                    event.getRecordSearchParams().setQuery(new Extension.BooleanExpression().withAnd(List.of(query, ownerFilterExp)));
                }

                return entity;
            });
        };
    }

    static <T extends Entity> Operator<T> execute(BiConsumer<Extension.Event, T> consumer) {
        return handler -> handler.operate((event, entity) -> {
            consumer.accept(event, entity);

            return entity;
        });
    }

    static <T extends Entity> Operator<T> execute(BiFunction<Extension.Event, T, T> consumer) {
        return handler -> handler.operate(consumer);
    }

    static <T extends Entity> Operator<T> reject(String message) {
        return check((event, entity) -> false, message);
    }

    static <T extends Entity> Operator<T> reject() {
        return reject("Operation rejected");
    }

    static <T extends Entity> Operator<T> check(BiPredicate<Extension.Event, T> condition) {
        return check(condition, "Condition not met");
    }

    static <T extends Entity> Operator<T> check(Predicate<T> condition) {
        return check((t, e) -> condition.test(e), "Condition not met");
    }

    static <T extends Entity> Operator<T> check(Predicate<T> condition, String failMessage) {
        return check((t, e) -> condition.test(e), failMessage);
    }

    static <T extends Entity> Operator<T> check(BiPredicate<Extension.Event, T> condition, String failMessage) {
        return handler -> handler.operate((event, entity) -> {
            if (condition.test(event, entity)) {
                return entity;
            } else {
                throw new ApiException(Extension.Code.RECORD_VALIDATION_ERROR, failMessage);
            }
        });
    }
}
