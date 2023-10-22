package io.apibrew.client.impl;

import io.apibrew.client.*;
import io.apibrew.client.model.Extension;
import kong.unirest.HttpResponse;
import kong.unirest.Unirest;
import lombok.Builder;
import lombok.RequiredArgsConstructor;
import lombok.extern.log4j.Log4j2;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.concurrent.atomic.AtomicBoolean;
import java.util.function.Consumer;

import static io.apibrew.client.helper.EventHelper.shortInfo;

@Log4j2
@Builder
public class WatcherImpl<T extends Entity> implements Watcher<T> {

    private final AtomicBoolean running = new AtomicBoolean(true);

    private final Thread thread = new Thread(this::run);
    private final AtomicBoolean stillRunning = new AtomicBoolean(false);
    private final AtomicBoolean isStopped = new AtomicBoolean(false);

    private final Client client;
    private final EntityInfo<T> entityInfo;
    private final Consumer<Extension.Event> consumer;

    // runs in background
    public void start() {
        ensureCanRun();
        thread.setName("Watch:[" + entityInfo + "]");
        thread.start();
    }

    @Override
    public void run() {
        ensureCanRun();
        stillRunning.set(true);

        while (isRunning()) {
            try {
                log.debug("Begin watching: {}", entityInfo);

                HttpResponse<?> result = Unirest.get(ClientImpl.Urls.recordWatchUrl(client.getUrl(), entityInfo.getRestPath()))
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

                log.debug("watch completed: {}", entityInfo);
            } catch (Exception e) {
                log.error("Error watching: " + entityInfo, e);
                try {
                    Thread.sleep(1000);
                } catch (InterruptedException ex) {
                    throw new RuntimeException(ex);
                }
            }
        }
    }

    private void ensureCanRun() {
        if (isStopped.get()) {
            throw new IllegalStateException("Poller is stopped");
        }

        if (stillRunning.get()) {
            throw new IllegalStateException("Poller is already running");
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
