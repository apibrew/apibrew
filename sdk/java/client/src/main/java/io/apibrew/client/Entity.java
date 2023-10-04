package io.apibrew.client;

import lombok.SneakyThrows;

import java.beans.IntrospectionException;
import java.beans.Introspector;
import java.util.*;

public abstract class Entity {

    public abstract UUID getId();

    public abstract void setId(UUID id);

    public static Map<String, Object> beanProperties(Object bean) {
        try {
            Map<String, Object> map = new HashMap<>();
            Arrays.stream(Introspector.getBeanInfo(bean.getClass(), Object.class)
                            .getPropertyDescriptors())
                    // filter out properties with setters only
                    .filter(pd -> Objects.nonNull(pd.getReadMethod()))
                    .forEach(pd -> { // invoke method to get value
                        try {
                            Object value = pd.getReadMethod().invoke(bean);
                            if (value != null) {
                                map.put(pd.getName(), value);
                            }
                        } catch (Exception e) {
                            // add proper error handling here
                        }
                    });
            return map;
        } catch (IntrospectionException e) {
            // and here, too
            return Collections.emptyMap();
        }
    }

    @SneakyThrows
    public Object getProperty(String property) {
        Map<String, Object> bp = beanProperties(this);
        return bp.get(property);
    }

    public void setProperty(String property, Object value) {
    }
}
