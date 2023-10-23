package io.apibrew.common;

import io.apibrew.client.model.Extension;

import java.util.List;
import java.util.Objects;

public class ExtensionInfo {
    private final List<String> namespaces;
    private final List<String> resources;
    private final Extension.Action action;

    private final boolean sync;
    private final boolean responds;
    private final boolean finalizer;

    private final boolean sealResource;

    private final int order;

    public ExtensionInfo(List<String> namespaces, List<String> resources, Extension.Action action, boolean sync, boolean responds, boolean finalizer, int order, boolean sealResource) {
        this.namespaces = namespaces;
        this.resources = resources;
        this.action = action;
        this.sync = sync;
        this.responds = responds;
        this.finalizer = finalizer;
        this.order = order;
        this.sealResource = sealResource;
    }

    public ExtensionInfo() {
        this.namespaces = null;
        this.resources = null;
        this.action = null;
        this.order = 0;
        this.sync = true;
        this.responds = true;
        this.finalizer = false;
        this.sealResource = false;
    }

    public ExtensionInfo withResources(List<String> list) {
        return new ExtensionInfo(this.namespaces, list, action, sync, responds, finalizer, this.order, this.sealResource);
    }

    public ExtensionInfo withNamespace(String namespace) {
        if (this.sealResource) {
            throw new IllegalStateException("Cannot change namespace after resource is set");
        }
        return new ExtensionInfo(List.of(namespace), this.resources, this.action, sync, responds, finalizer, this.order, this.sealResource);
    }

    public ExtensionInfo withResource(String resource) {
        if (this.sealResource) {
            throw new IllegalStateException("Cannot change namespace after resource is set");
        }
        return new ExtensionInfo(this.namespaces, List.of(resource), this.action, sync, responds, finalizer, this.order, this.sealResource);
    }

    public ExtensionInfo withAction(Extension.Action action) {
        return new ExtensionInfo(this.namespaces, this.resources, action, sync, responds, finalizer, this.order, this.sealResource);
    }

    public ExtensionInfo withOrder(int order) {
        return new ExtensionInfo(this.namespaces, this.resources, this.action, sync, responds, finalizer, order, this.sealResource);
    }

    public ExtensionInfo withSync(boolean sync) {
        return new ExtensionInfo(this.namespaces, this.resources, this.action, sync, responds, finalizer, this.order, this.sealResource);
    }

    public ExtensionInfo withResponds(boolean responds) {
        return new ExtensionInfo(this.namespaces, this.resources, this.action, sync, responds, finalizer, this.order, this.sealResource);
    }

    public ExtensionInfo withFinalizer(boolean finalizer) {
        return new ExtensionInfo(this.namespaces, this.resources, this.action, sync, responds, this.finalizer, this.order, this.sealResource);
    }

    public ExtensionInfo withSealResource(boolean sealResource) {
        return new ExtensionInfo(this.namespaces, this.resources, this.action, sync, responds, this.finalizer, this.order, sealResource);
    }

    public boolean isSync() {
        return sync;
    }

    public boolean isResponds() {
        return responds;
    }

    public boolean isFinalizer() {
        return finalizer;
    }

    public Extension.Action getAction() {
        return action;
    }

    public int getOrder() {
        return order;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        ExtensionInfo that = (ExtensionInfo) o;
        return sync == that.sync && responds == that.responds && finalizer == that.finalizer && order == that.order && Objects.equals(namespaces, that.namespaces) && Objects.equals(resources, that.resources) && action == that.action;
    }

    @Override
    public int hashCode() {
        return Objects.hash(namespaces, resources, action, sync, responds, finalizer, order);
    }

    public Extension toExtension() {
        Extension extension = new Extension();

        extension.setName(namespaces + "/" + resources + "/" + action + "/" + order + "/" + sync + "/" + responds + "/" + finalizer);

        extension.setFinalizes(finalizer);
        extension.setSync(sync);
        extension.setResponds(responds);
        extension.setOrder(order);
        extension.setSelector(prepareEventSelector());


        return extension;
    }

    private Extension.EventSelector prepareEventSelector() {
        Extension.EventSelector eventSelector = new Extension.EventSelector();

        if (namespaces != null) {
            eventSelector.setNamespaces(namespaces);
        }

        if (resources != null) {
            eventSelector.setResources(resources);
        }

        if (action != null) {
            eventSelector.setActions(List.of(action));
        }

        return eventSelector;
    }

    @Override
    public String toString() {
        return "ExtensionInfo{" +
                "namespaces=" + namespaces +
                ", resources=" + resources +
                ", action=" + action +
                ", sync=" + sync +
                ", responds=" + responds +
                ", finalizer=" + finalizer +
                ", order=" + order +
                '}';
    }
}
