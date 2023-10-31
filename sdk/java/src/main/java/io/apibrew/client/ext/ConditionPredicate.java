package io.apibrew.client.ext;


import io.apibrew.client.Entity;
import io.apibrew.client.model.Extension;

import java.util.Objects;
import java.util.function.Function;

public class ConditionPredicate<T extends Entity> {

    public static <T extends Entity, K> java.util.function.BiPredicate<Extension.Event, T> equal(Function<T, K> getter, K value) {
        return (e, t) -> Objects.equals(getter.apply(t), value);
    }

    public static <T extends Entity, K> java.util.function.BiPredicate<Extension.Event, T> equalToUsername(Function<T, K> getter) {
        return (e, t) -> {
            return Objects.equals(getter.apply(t), "$username");
        };
    }

}
