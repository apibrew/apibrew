package io.apibrew.common.impl;

import io.apibrew.client.ApiException;
import io.apibrew.client.Client;
import io.apibrew.client.impl.ChannelEventPoller;
import io.apibrew.common.impl.AbstractExtensionServiceImpl;
import io.apibrew.common.ext.ExtensionService;
import io.apibrew.client.model.Extension;
import lombok.extern.log4j.Log4j2;

import java.io.IOException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

@Log4j2
public class PollerExtensionService extends AbstractExtensionServiceImpl implements ExtensionService {

    private final String channelKey;
    private final ChannelEventPoller poller;
    private final ExecutorService executorService = Executors.newFixedThreadPool(200);

    public PollerExtensionService(Client client, String channelKey) {
        super(client);
        this.channelKey = channelKey;
        this.poller = ChannelEventPoller.builder()
                .consumer(this::handleEvent)
                .channelKey(channelKey)
                .client(client)
                .build();
    }

    private void handleEvent(Extension.Event event) {
        executorService.submit(() -> {
            try {
                log.trace("Received event: {}", objectMapper.writeValueAsString(event));
                Extension.Event processedEvent = processEvent(event);

                this.client.writeEvent(channelKey, processedEvent);
            } catch (ApiException e) {
                log.error("Unable to process event[ApiException]", e);
                event.setError(e.getError());

                this.client.writeEvent(channelKey, event);
            } catch (Exception e) {
                log.error("Unable to process event", e);
                event.setError(new Extension.Error().withMessage(e.getMessage()));

                this.client.writeEvent(channelKey, event);
            }
        });
    }

    protected Extension.ExternalCall prepareExternalCall() {
        Extension.ExternalCall externalCall = new Extension.ExternalCall();

        Extension.ChannelCall httpCall = new Extension.ChannelCall();
        httpCall.setChannelKey(channelKey);

        externalCall.setChannelCall(httpCall);
        return externalCall;
    }

    @Override
    public void run() throws IOException {
        registerExtensions();

        this.poller.run();
    }

    @Override
    public void close() throws Exception {
        this.poller.close();
    }
}
