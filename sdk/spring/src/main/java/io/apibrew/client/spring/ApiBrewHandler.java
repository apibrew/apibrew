package io.apibrew.client.spring;

import io.apibrew.client.Entity;
import io.apibrew.client.model.Extension;

import java.lang.annotation.*;

@Target({ElementType.TYPE, ElementType.METHOD})
@Retention(RetentionPolicy.RUNTIME)
@Documented
public @interface ApiBrewHandler {
    Class<? extends Entity> entityClass() default Entity.class;

    Extension.Action action() default Extension.Action.CREATE;

    public boolean async() default false;

    Order order() default Order.ON;

    enum Order {
        BEFORE,
        ON,
        AFTER
    }
}
