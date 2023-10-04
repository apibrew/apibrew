package io.apibrew.client.helper;

import io.apibrew.client.model.Extension;

public class EventHelper {

    public static String shortInfo(Extension.Event event) {
        return String.format("[%s]%s/%s/%s", event.getAction(), event.getResource().getNamespace().getName(), event.getResource().getName(), event.getId());
    }

}
