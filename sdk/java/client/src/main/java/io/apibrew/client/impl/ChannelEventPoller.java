package io.apibrew.client.impl;

import io.apibrew.client.ApiException;
import io.apibrew.client.Client;
import io.apibrew.client.model.Extension;
import kong.unirest.HttpResponse;
import kong.unirest.Unirest;
import lombok.Builder;
import lombok.extern.log4j.Log4j2;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.concurrent.atomic.AtomicBoolean;
import java.util.function.Consumer;

import static io.apibrew.client.helper.EventHelper.shortInfo;

@Builder
@Log4j2
public class ChannelEventPoller implements AutoCloseable {

    private final Client client;

    private final String channelKey;
    private final String threadName;

    private final Consumer<Extension.Event> consumer;
    private final AtomicBoolean stillRunning = new AtomicBoolean(false);
    private final AtomicBoolean isStopped = new AtomicBoolean(false);

    private final Thread thread = new Thread(this::run);

    // runs in background
    public void start() {
        ensureCanRun();
        thread.setName(threadName == null ? "ChannelEventPoller" : threadName);
        thread.start();
    }

    private void ensureCanRun() {
        if (isStopped.get()) {
            throw new IllegalStateException("Poller is stopped");
        }

        if (stillRunning.get()) {
            throw new IllegalStateException("Poller is already running");
        }
    }

    // runs in foreground
    public void run() {
        ensureCanRun();
        stillRunning.set(true);

        while (isRunning()) {
            try {
                log.debug("Begin Polling channel: {}", channelKey);

                HttpResponse<?> result = Unirest.get(ClientImpl.Urls.eventsUrl(client.getUrl()) + "?channelKey=" + channelKey)
                        .connectTimeout(10 * 1000)
                        .headers(client.headers())
                        .asObject(response -> {
                            if (!isRunning()) {
                                return null;
                            }

                            if (response.getStatus() != 200) {
                                throw new ApiException(response.getStatusText());
                            }

                            BufferedReader br = new BufferedReader(new InputStreamReader(response.getContent()));

                            try {
                                while (isRunning()) {
                                    String line = br.readLine();

                                    if (line == null) {
                                        break;
                                    }

                                    Extension.Event event = ClientImpl.objectMapper.readValue(line, Extension.Event.class);

                                    if (event.getId().equals("heartbeat-message")) {
                                        log.debug("Received heartbeat message");
                                        continue;
                                    }

                                    log.debug("Received event: {}", shortInfo(event));

                                    consumer.accept(event);

                                    return line;
                                }
                            } catch (IOException e) {
                                throw new RuntimeException(e);
                            } finally {
                                try {
                                    br.close();
                                } catch (IOException e) {
                                    throw new RuntimeException(e);
                                }
                            }

                            return null;
                        });

                ClientImpl.ensureResponseSuccess(result);

                log.debug("Polling channel: {} complete", channelKey);
            } catch (Exception e) {
                log.error("Error polling channel: {}", channelKey, e);
                try {
                    Thread.sleep(1000);
                } catch (InterruptedException ex) {
                    throw new RuntimeException(ex);
                }
            }
        }
    }

    private boolean isRunning() {
        return !isStopped.get() && stillRunning.get();
    }

    @Override
    public void close() {
        isStopped.set(true);
        stillRunning.set(false);

        thread.interrupt();
    }
}
