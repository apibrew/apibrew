package io.apibrew.common.impl;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import io.apibrew.client.Client;
import io.apibrew.client.Entity;
import io.apibrew.client.Repository;
import io.apibrew.common.ExtensionInfo;
import io.apibrew.common.ext.ExtensionService;
import io.apibrew.common.ext.GenericHandler;
import io.apibrew.common.ext.Handler;
import io.apibrew.client.model.Extension;
import io.apibrew.client.model.Record;
import lombok.extern.log4j.Log4j2;

import java.util.*;
import java.util.function.BiFunction;

import static io.apibrew.client.helper.EventHelper.shortInfo;

@Log4j2
public abstract class AbstractExtensionServiceImpl implements ExtensionService {
    protected final Client client;
    private final Set<ExtensionInfo> extensionInfoSet = new HashSet<>();
    private final Set<ExtensionInfo> registeredExtensionInfoSet = new HashSet<>();
    private final Map<String, ExtensionInfo> extensionInfoIdMap = new HashMap<>();
    private final Map<ExtensionInfo, List<BiFunction<Extension.Event, Record, Record>>> extensionHandlerMap = new HashMap<>();
    private final Repository<Extension> extensionRepo;

    protected final ObjectMapper objectMapper = new ObjectMapper();


    protected AbstractExtensionServiceImpl(Client client) {
        this.client = client;
        this.extensionRepo = client.repo(Extension.class);

        objectMapper.registerModule(new JavaTimeModule());
    }

    @Override
    public <T extends Entity> Handler<T> handler(Class<T> entityClass) {
        return new HandlerImpl<T>(client, this, entityClass);
    }

    @Override
    public GenericHandler genericHandler() {
        return new GenericHandlerImpl(client, this);
    }

    protected synchronized void registerExtensions() {
        extensionInfoSet.forEach(extensionInfo -> {
            if (!registeredExtensionInfoSet.contains(extensionInfo)) {
                registeredExtensionInfoSet.add(extensionInfo);

                registerExtension(extensionInfo);
            }
        });
    }

    private void registerExtension(ExtensionInfo extensionInfo) {
        Extension extension = extensionInfo.toExtension();

        extension.setCall(prepareExternalCall());

        extension = extensionRepo.apply(extension);

        extensionInfoIdMap.put(extension.getId().toString(), extensionInfo);
    }

    protected abstract Extension.ExternalCall prepareExternalCall();

    protected Extension.Event processEvent(Extension.Event event) {
        log.debug("Begin processing event: {}", shortInfo(event));
        // normalize event

        if (event.getAnnotations() == null) {
            event.setAnnotations(new HashMap<>());
        }

        String extensionId = event.getAnnotations().get("ExtensionId");
        ExtensionInfo extensionInfo = extensionInfoIdMap.get(extensionId);

        log.trace("Event ID: {} => Extension ID: {}", event.getId(), extensionId);
        log.trace("ExtensionInfo: {}", extensionInfo);

        if (extensionInfo == null) {
            log.warn("ExtensionInfo not found for event: {}", shortInfo(event));
        }

        Extension.Event eventChain = processEvent(extensionInfo, event);

        log.debug("End processing event: {}", shortInfo(event));

        return eventChain;
    }

    private Extension.Event processEvent(ExtensionInfo extensionInfo, Extension.Event eventChain) {
        List<BiFunction<Extension.Event, Record, Record>> handlers = extensionHandlerMap.get(extensionInfo);

        if (handlers != null) {
            for (BiFunction<Extension.Event, Record, Record> handler : handlers) {
                List<Record> records = eventChain.getRecords();
                boolean handlerHandled = false;
                if (records != null) {
                    List<Record> eventChainRecords = eventChain.getRecords();
                    List<Record> processedRecords = new ArrayList<>();

                    for (Record record : eventChainRecords) {
                        log.debug("Processing record: {}", record.getId());
                        Record processedRecord = handler.apply(eventChain, record);
                        if (processedRecord != null) {
                            processedRecords.add(processedRecord);
                        }
                        handlerHandled = true;
                    }

                    eventChain.setRecords(processedRecords);
                }

                if (!handlerHandled) {
                    handler.apply(eventChain, null);
                }


            }
        }

        return eventChain;
    }

    @Override
    public void registerExtensionWithOperator(ExtensionInfo extensionInfo, BiFunction<Extension.Event, Record, Record> operator) {
        extensionInfoSet.add(extensionInfo);

        extensionHandlerMap.putIfAbsent(extensionInfo, new ArrayList<>());

        extensionHandlerMap.get(extensionInfo).add(operator);
    }
}
