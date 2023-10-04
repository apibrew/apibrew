package io.apibrew.common.ext;

import io.apibrew.client.Entity;
import io.apibrew.common.ExtensionInfo;
import io.apibrew.client.model.Extension;

import java.util.function.BiPredicate;
import java.util.function.Function;

public class SimpleCondition<T extends Entity> implements Condition<T> {
    private final Function<ExtensionInfo, ExtensionInfo> extensionInfoConfigurer;

    private final BiPredicate<Extension.Event, T> predicate;

    public SimpleCondition(Function<ExtensionInfo, ExtensionInfo> extensionInfoConfigurer, BiPredicate<Extension.Event, T> predicate) {
        this.extensionInfoConfigurer = extensionInfoConfigurer;
        this.predicate = predicate;
    }

    public SimpleCondition(Function<ExtensionInfo, ExtensionInfo> extensionInfoConfigurer) {
        this.extensionInfoConfigurer = extensionInfoConfigurer;
        this.predicate = (e, t) -> true;
    }

    @Override
    public ExtensionInfo configureExtensionInfo(ExtensionInfo extensionInfo) {
        return extensionInfoConfigurer.apply(extensionInfo);
    }

    @Override
    public boolean eventMatches(Extension.Event event, T entity) {
        return predicate.test(event, entity);
    }
}
