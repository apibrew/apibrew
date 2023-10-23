package io.apibrew.common.impl;

import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpServer;
import io.apibrew.client.ApiException;
import io.apibrew.client.Client;
import io.apibrew.common.ext.ExtensionService;
import io.apibrew.client.model.Extension;
import lombok.extern.log4j.Log4j2;

import java.io.IOException;
import java.net.InetSocketAddress;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

@Log4j2
public class HostedExtensionServiceImpl extends AbstractExtensionServiceImpl implements ExtensionService {

    private final String host;
    private final int port;
    private final String remoteUrl;
    private HttpServer server;

    public HostedExtensionServiceImpl(Client client, String host, int port, String remoteUrl) {
        super(client);
        this.host = host;
        this.port = port;
        this.remoteUrl = remoteUrl;
    }

    ExecutorService executorService = Executors.newFixedThreadPool(200);

    protected Extension.ExternalCall prepareExternalCall() {
        Extension.ExternalCall externalCall = new Extension.ExternalCall();

        Extension.HttpCall httpCall = new Extension.HttpCall();
        httpCall.setUri(remoteUrl);
        httpCall.setMethod("POST");

        externalCall.setHttpCall(httpCall);
        return externalCall;
    }

    @Override
    public void run() throws IOException {
        registerExtensions();
        this.server = HttpServer.create(new InetSocketAddress(host, port), 0);

        server.createContext("/").setHandler(e -> {
            executorService.execute(() -> handle(e));
        });

        server.start();
    }

    private void handle(HttpExchange httpExchange) {
        try {
            String msg = new String(httpExchange.getRequestBody().readAllBytes());

            log.debug("Received message: {}", msg);

            Extension.Event event = objectMapper.readValue(msg, Extension.Event.class);

            event = processEvent(event);

            event.setResource(null);

            String message = objectMapper.writeValueAsString(event);

            log.debug("Sending message: {}", message);

            httpExchange.sendResponseHeaders(200, message.getBytes().length);
            httpExchange.getResponseBody().write(message.getBytes());
        } catch (ApiException e) {
            try {
                log.error(e);

                String message = objectMapper.writeValueAsString(e.getError());

                httpExchange.sendResponseHeaders(400, message.length());
                httpExchange.getResponseBody().write(message.getBytes());
            } catch (Exception e2) {
                log.error(e2);
            }
        } catch (Exception e) {
            log.error(e);
            try {
                Extension.Error error = new Extension.Error();

                error.setCode(Extension.Code.RECORD_VALIDATION_ERROR);
                error.setMessage(e.getMessage());

                String message = objectMapper.writeValueAsString(error);

                httpExchange.sendResponseHeaders(400, message.length());
                httpExchange.getResponseBody().write(message.getBytes());
            } catch (IOException ex) {
                System.out.println(ex);
            }
        } finally {
            try {
                httpExchange.getResponseBody().close();
            } catch (IOException e) {
                System.out.println(e);
            }
        }
    }

    @Override
    public void close() throws Exception {
        executorService.shutdown();
        this.server.stop(0);
    }
}
