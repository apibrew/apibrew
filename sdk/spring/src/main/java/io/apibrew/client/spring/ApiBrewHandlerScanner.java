package io.apibrew.client.spring;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import io.apibrew.client.ApiException;
import io.apibrew.client.Entity;
import io.apibrew.client.EntityInfo;
import io.apibrew.client.ExtensionInfo;
import io.apibrew.client.ext.ExtensionService;
import io.apibrew.client.model.Extension;
import io.apibrew.client.model.Record;
import io.apibrew.client.spring.ApiBrewHandler;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.InitializingBean;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.Lazy;
import org.springframework.stereotype.Component;

import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;
import java.lang.reflect.Modifier;
import java.util.Map;
import java.util.function.BiFunction;

import static io.apibrew.client.spring.ApiBrewHandler.Order.ON;

@RequiredArgsConstructor
@Component
@Lazy(false)
public class ApiBrewHandlerScanner implements InitializingBean {

    private final ApplicationContext applicationContext;
    private final ExtensionService extensionService;
    private final ObjectMapper objectMapper = new ObjectMapper();

    public void init() {
        objectMapper.registerModule(new JavaTimeModule());
        objectMapper.disable(SerializationFeature.WRITE_DATES_AS_TIMESTAMPS);

        Map<String, Object> beansWithAnnotation = applicationContext.getBeansWithAnnotation(ApiBrewHandler.class);

        ExtensionInfo extensionInfo = new ExtensionInfo();

        for (Object bean : beansWithAnnotation.values()) {
            Class<?> clazz = bean.getClass();
            ApiBrewHandler annotation = clazz.getAnnotation(ApiBrewHandler.class);

            extensionInfo = mapAnnotationToExtensionInfo(extensionInfo, annotation);

            Class<? extends Entity> entityClass = annotation.entityClass();

            if (entityClass != null && entityClass.equals(Entity.class)) {
                entityClass = null;
            }

            Method[] methods = clazz.getDeclaredMethods();
            for (Method method : methods) {
                if (method.isAnnotationPresent(ApiBrewHandler.class)) {
                    if (!Modifier.isPublic(method.getModifiers())) {
                        throw new RuntimeException("Handler method " + bean.getClass().getName() + "." + method.getName() + " must be public");
                    }

                    ApiBrewHandler methodAnnotation = method.getAnnotation(ApiBrewHandler.class);

                    if (entityClass == null) {
                        entityClass = methodAnnotation.entityClass();
                    }

                    if (entityClass != null && entityClass.equals(Entity.class)) {
                        entityClass = null;
                    }

                    if (entityClass == null) {
                        throw new RuntimeException("Entity class not specified for handler " + method.getName());
                    }

                    extensionInfo = mapAnnotationToExtensionInfo(extensionInfo, methodAnnotation);

                    registerMethod(bean, extensionInfo, method, entityClass, EntityInfo.fromEntityClass(entityClass));
                }
            }
        }

        extensionService.registerPendingItems();
    }

    private void registerMethod(Object bean, ExtensionInfo extensionInfo, Method method, Class<? extends Entity> entityClass, EntityInfo<? extends Entity> entityInfo) {
        BiFunction<Extension.Event, Record, Record> handler = (event, record) -> handle(bean, extensionInfo, method, event, record, entityClass, entityInfo);

        extensionService.registerExtensionWithOperator(extensionInfo, handler);
    }

    private Record handle(Object bean, ExtensionInfo extensionInfo, Method method, Extension.Event event, Record record, Class<? extends Entity> entityClass, EntityInfo<? extends Entity> entityInfo) {
        Entity entity = recordToEntity(entityInfo, record);

        Entity processedEntity;

        Object[] args = new Object[method.getParameterTypes().length];

        for (int i = 0; i < args.length; i++) {
            Class<?> typ = method.getParameterTypes()[i];

            Object value;
            if (typ.equals(entityClass)) {
                value = entity;
            } else if (typ.equals(Extension.Event.class)) {
                value = event;
            } else if (typ.equals(ExtensionInfo.class)) {
                value = extensionInfo;
            } else if (typ.equals(Record.class)) {
                value = record;
            } else {
                value = null;
            }

            args[i] = value;
        }

        try {
            Object result;
            try {
                result = method.invoke(bean, args);
            } catch (InvocationTargetException t) {
                throw (Exception) t.getTargetException();
            }

            Class<?> rt = method.getReturnType();

            if (rt.equals(void.class)) {
                processedEntity = entity;
            } else {
                processedEntity = (Entity) result;
            }
        } catch (ApiException e) {
            throw e;
        } catch (Exception e) {
            e.printStackTrace();
            throw new ApiException(Extension.Code.INTERNAL_ERROR, e.getMessage());
        }

        return recordFromEntity(processedEntity);
    }

    private Entity recordToEntity(EntityInfo entityInfo, Record entity) {
        if (entity == null) {
            return null;
        }

        return (Entity) objectMapper.convertValue(entity.getProperties(), entityInfo.getEntityClass());
    }

    private Record recordFromEntity(Entity entity) {
        if (entity == null) {
            return null;
        }

        Record record = new Record();
        record.setProperties(objectMapper.convertValue(entity, Object.class));
        record.setId(entity.getId());

        return record;
    }

    private ExtensionInfo mapAnnotationToExtensionInfo(ExtensionInfo extensionInfo, ApiBrewHandler annotation) {
        if (annotation.entityClass() != null && !annotation.entityClass().equals(Entity.class)) {
            EntityInfo<?> entityInfo = EntityInfo.fromEntityClass(annotation.entityClass());

            extensionInfo = extensionInfo.withNamespace(entityInfo.getNamespace());
            extensionInfo = extensionInfo.withResource(entityInfo.getResource());
        }

        extensionInfo = extensionInfo.withSync(!annotation.async());
        extensionInfo = extensionInfo.withFinalizer(annotation.order() == ON);
        extensionInfo = extensionInfo.withResponds(!annotation.async());

        if (annotation.action() != null) {
            extensionInfo = extensionInfo.withAction(annotation.action());
        }

        if (annotation.order() != null) {
            extensionInfo = extensionInfo.withOrder(switch (annotation.order()) {
                case ON -> 100;
                case BEFORE -> 10;
                case AFTER -> 110;
            });
        }

        return extensionInfo;
    }

    @Override
    public void afterPropertiesSet() throws Exception {
        init();
    }
}
