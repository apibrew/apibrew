package io.apibrew.client;

import io.apibrew.client.model.Extension;
import io.apibrew.client.model.Extension.BooleanExpression;
import lombok.experimental.UtilityClass;

import java.util.List;

@UtilityClass
public class BooleanExpressionBuilder {
    public BooleanExpression and(BooleanExpression... expressions) {
        return new BooleanExpression().withAnd(List.of(expressions));
    }
    public BooleanExpression and(List<BooleanExpression> expressions) {
        return new BooleanExpression().withAnd(expressions);
    }

    public BooleanExpression or(BooleanExpression... expressions) {
        return new BooleanExpression().withOr(List.of(expressions));
    }

    public BooleanExpression or(List<BooleanExpression> expressions) {
        return new BooleanExpression().withOr(expressions);
    }

    public BooleanExpression not(BooleanExpression expression) {
        return new BooleanExpression().withNot(expression);
    }

    public BooleanExpression equal(String property, Object value) {
        return new BooleanExpression().withEqual(
                new Extension.PairExpression()
                        .withLeft(new Extension.Expression().withProperty(property))
                        .withRight(new Extension.Expression().withValue(value))
        );
    }

    public BooleanExpression eq(String property, Object value) {
        return new BooleanExpression().withEqual(
                new Extension.PairExpression()
                        .withLeft(new Extension.Expression().withProperty(property))
                        .withRight(new Extension.Expression().withValue(value))
        );
    }

    public BooleanExpression notEqual(String property, Object value) {
        return not(equal(property, value));
    }

    public BooleanExpression greaterThan(String property, Object value) {
        return new BooleanExpression().withGreaterThan(
                new Extension.PairExpression()
                        .withLeft(new Extension.Expression().withProperty(property))
                        .withRight(new Extension.Expression().withValue(value))
        );
    }

    public BooleanExpression gt(String property, Object value) {
        return greaterThan(property, value);
    }

    public BooleanExpression greaterThanOrEqual(String property, Object value) {
        return new BooleanExpression().withGreaterThanOrEqual(
                new Extension.PairExpression()
                        .withLeft(new Extension.Expression().withProperty(property))
                        .withRight(new Extension.Expression().withValue(value))
        );
    }

    public BooleanExpression gte(String property, Object value) {
        return greaterThanOrEqual(property, value);
    }

    public BooleanExpression lessThan(String property, Object value) {
        return new BooleanExpression().withLessThan(
                new Extension.PairExpression()
                        .withLeft(new Extension.Expression().withProperty(property))
                        .withRight(new Extension.Expression().withValue(value))
        );
    }

    public BooleanExpression lt(String property, Object value) {
        return lessThan(property, value);
    }

    public BooleanExpression lessThanOrEqual(String property, Object value) {
        return new BooleanExpression().withLessThanOrEqual(
                new Extension.PairExpression()
                        .withLeft(new Extension.Expression().withProperty(property))
                        .withRight(new Extension.Expression().withValue(value))
        );
    }

    public BooleanExpression lte(String property, Object value) {
        return lessThanOrEqual(property, value);
    }

    public BooleanExpression in(String property, Object... values) {
        return new BooleanExpression().withIn(
                new Extension.PairExpression()
                        .withLeft(new Extension.Expression().withProperty(property))
                        .withRight(new Extension.Expression().withValue(values))
        );
    }

    public BooleanExpression notIn(String property, Object... values) {
        return not(in(property, values));
    }

    public BooleanExpression isNull(String property) {
        return new BooleanExpression().withIsNull(
                new Extension.Expression().withProperty(property)
        );
    }

    public BooleanExpression isNotNull(String property) {
        return not(isNull(property));
    }
}
