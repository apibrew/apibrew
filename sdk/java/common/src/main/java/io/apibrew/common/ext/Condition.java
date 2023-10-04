package io.apibrew.common.ext;

import io.apibrew.client.Entity;
import io.apibrew.client.EntityInfo;
import io.apibrew.client.model.Extension;
import io.apibrew.common.ExtensionInfo;

import java.util.Arrays;
import java.util.Collections;
import java.util.Objects;

import static io.apibrew.client.model.Extension.Action.*;

public interface Condition<T extends Entity> {

    ExtensionInfo configureExtensionInfo(ExtensionInfo extensionInfo);

    boolean eventMatches(Extension.Event event, T entity);

    static <T extends Entity> Condition<T> before() {
        return new SimpleCondition<>(ei -> ei.withOrder(10));
    }

    static <T extends Entity> Condition<T> after() {
        return new SimpleCondition<>(ei -> ei.withOrder(110));
    }

    static <T extends Entity> Condition<T> create() {
        return new SimpleCondition<>(ei -> ei.withAction(CREATE));
    }

    static <T extends Entity> Condition<T> update() {
        return new SimpleCondition<>(ei -> ei.withAction(UPDATE));
    }

    static <T extends Entity> Condition<T> delete() {
        return new SimpleCondition<>(ei -> ei.withAction(DELETE));
    }

    static <T extends Entity> Condition<T> get() {
        return new SimpleCondition<>(ei -> ei.withAction(GET));
    }

    static <T extends Entity> Condition<T> list() {
        return new SimpleCondition<>(ei -> ei.withAction(LIST));
    }

    static <T extends Entity> Condition<T> beforeCreate() {
        return and(before(), create());
    }

    static <T extends Entity> Condition<T> beforeUpdate() {
        return and(before(), update());
    }

    static <T extends Entity> Condition<T> beforeDelete() {
        return and(before(), delete());
    }

    static <T extends Entity> Condition<T> beforeGet() {
        return and(before(), get());
    }

    static <T extends Entity> Condition<T> beforeList() {
        return and(before(), list());
    }

    static <T extends Entity> Condition<T> afterCreate() {
        return and(after(), create());
    }

    static <T extends Entity> Condition<T> afterUpdate() {
        return and(after(), update());
    }

    static <T extends Entity> Condition<T> afterList() {
        return and(after(), list());
    }

    static <T extends Entity> Condition<T> afterDelete() {
        return and(after(), delete());
    }

    static <T extends Entity> Condition<T> afterGet() {
        return and(after(), get());
    }

    static <T extends Entity> Condition<T> async() {
        return new SimpleCondition<>(ei -> ei.withSync(false));
    }

    static <T extends Entity> Condition<T> entityExists() {
        return new SimpleCondition<>(ei -> ei, (e, t) -> t != null);
    }

    static <T extends Entity> Condition<T> user(String... expectedUsers) {
        return new SimpleCondition<>(ei -> ei, (e, t) -> {
            String user = e.getAnnotations().get("user");
            if (user == null) {
                return false;
            }

            for (String eu : expectedUsers) {
                if (Objects.equals(user, eu)) {
                    return true;
                }
            }

            return false;

        });
    }

    static <T extends Entity> Condition<T> resource(String namespace, String... resource) {
        return new SimpleCondition<>(ei -> {
            return ei.withNamespace(namespace).withResources(Arrays.asList(resource));
        }, (e, t) -> {
            return true;
        });
    }

    static <T extends Entity> Condition<T> resource(Class<T> entityClass) {
        return resource(EntityInfo.fromEntityClass(entityClass));
    }

    static <T extends Entity> Condition<T> resource(EntityInfo<T> entityInfo) {
        return new SimpleCondition<>(ei -> {
            return ei.withNamespace(entityInfo.getNamespace()).withResources(Collections.singletonList(entityInfo.getResource()));
        }, (e, t) -> {
            return true;
        });
    }

    static <T extends Entity> Condition<T> group(String... expectedGroups) {
        return new SimpleCondition<>(ei -> ei, (e, r) -> {
            String groups = e.getAnnotations().get("groups");
            if (groups == null) {
                return false;
            }

            for (String g : groups.split(",")) {
                for (String eg : expectedGroups) {
                    if (Objects.equals(g, eg)) {
                        return true;
                    }
                }
            }

            return false;
        });
    }

    static <T extends Entity> Condition<T> and(Condition<T> condition1, Condition<T> condition2) {
        return new SimpleCondition<>(ei -> condition2.configureExtensionInfo(condition1.configureExtensionInfo(ei)),
                (e, t) -> condition1.eventMatches(e, t) && condition2.eventMatches(e, t));
    }
}
