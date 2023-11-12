package io.apibrew.client.impl;

import com.fasterxml.jackson.databind.JavaType;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import io.apibrew.client.*;
import io.apibrew.client.model.Extension;
import io.apibrew.client.model.Resource;
import io.apibrew.client.model.Token;
import kong.unirest.GenericType;
import kong.unirest.HttpResponse;
import kong.unirest.Unirest;
import lombok.Getter;
import lombok.SneakyThrows;
import lombok.extern.log4j.Log4j2;

import java.io.IOException;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.stream.Collectors;

import static io.apibrew.client.helper.EventHelper.shortInfo;

@Log4j2
public class ClientImpl implements Client {

    private boolean bypassExtensions;

    private final Map<Class<?>, Repository<?>> repositoryClassMap = new HashMap<>();
    private final Map<EntityInfo<?>, Repository<?>> repositoryEntityMap = new HashMap<>();


    public static class Urls {
        static String eventsUrl(String url) {
            return String.format("%s/_events", url);
        }

        static String resourceUrl(String url) {
            return String.format("%s/resources", url);
        }

        static String recordUrl(String url, String restPath) {
            return String.format("%s/%s", url, restPath);
        }

        static String recordSearchUrl(String url, String restPath) {
            return String.format("%s/%s/_search", url, restPath);
        }

        static String recordWatchUrl(String url, String restPath) {
            return String.format("%s/%s/_watch", url, restPath);
        }

        static String resourceByName(String url, String namespace, String name) {
            return String.format("%s/by-name/%s/%s", resourceUrl(url), namespace, name);
        }

        static String resourceById(String url, String id) {
            return String.format("%s/%s", resourceUrl(url), id);
        }

        static String recordByIdUrl(String url, String restPath, String id) {
            return String.format("%s/%s/%s", url, restPath, id);
        }

        static String recordActionByIdUrl(String url, String restPath, String id, String action) {
            return String.format("%s/%s/%s/_%s", url, restPath, id, action);
        }

        static String authenticate(String url) {
            return String.format("%s/authentication/token", url);
        }
    }

    @Getter
    private final String url;
    private String token;

    static final ObjectMapper objectMapper = new ObjectMapper();

    public ClientImpl(String url) {
        this.url = url;
        objectMapper.registerModule(new JavaTimeModule());
    }

    public static Client newClient() {
        Config config = ConfigLoader.load();

        Config.Server serverConfig = config.getServers().stream().filter(item -> item.getName().equals(config.getDefaultServer())).findAny().orElseThrow(() -> new IllegalArgumentException("Default server cannot be located"));

        return newClientByServerConfig(serverConfig);
    }

    public static Client newClientByServerName(String serverName) {
        Config config = ConfigLoader.load();

        Config.Server serverConfig = config.getServers().stream().filter(item -> item.getName().equals(serverName)).findAny().orElseThrow(() -> new IllegalArgumentException("Server not found:" + serverName));

        return newClientByServerConfig(serverConfig);
    }

    public static ClientImpl newClientByServerConfig(Config.Server serverConfig) {
        int httpPort = serverConfig.getHttpPort();

        if (httpPort == 0) {
            httpPort = serverConfig.getPort();
        }

        String addr = serverConfig.getHost() + ":" + httpPort;

        if (!addr.startsWith("http")) {
            if (serverConfig.isInsecure()) {
                addr = "http://" + addr;
            } else {
                addr = "https://" + addr;
            }
        }

        if (addr.endsWith("/")) {
            addr = addr.substring(0, addr.length() - 1);
        }

        ClientImpl client = new ClientImpl(addr);

        if (serverConfig.getAuthentication().getToken() != null && !serverConfig.getAuthentication().getToken().trim().isEmpty()) {
            client.authenticateWithToken(serverConfig.getAuthentication().getToken());
        } else {
            client.authenticateWithUsernameAndPassword(serverConfig.getAuthentication().getUsername(), serverConfig.getAuthentication().getPassword());
        }

        return client;
    }

    public void authenticateWithToken(String token) {
        this.token = token;
    }

    @Override
    public Resource applyResource(Resource resource) {
        HttpResponse<Resource> resp = Unirest.get(Urls.resourceByName(url, resource.getNamespace().getName(), resource.getName())).headers(headers()).asObject(Resource.class);

        int existsStatus = resp.getStatus();

        if (existsStatus == 200) {
            return updateResource(resource);
        } else if (existsStatus == 404) {
            return createResource(resource);
        } else {
            return ensureResponseSuccess(resp);
        }
    }

    @Override
    public Resource getResourceByName(String namespace, String name) {
        return Unirest.get(Urls.resourceByName(url, namespace, name)).headers(headers()).asObject(Resource.class).getBody();
    }

    @Override
    public List<Resource> listResources() {
        HttpResponse<List<Resource>> result = Unirest.get(Urls.resourceUrl(url)).headers(headers()).asObject(new GenericType<List<Resource>>() {
        });

        ensureResponseSuccess(result);

        return result.getBody();
    }

    @Override
    public Resource createResource(Resource resource) {
        HttpResponse<Resource> result = Unirest.post(Urls.resourceUrl(url)).body(resource).headers(headers()).asObject(Resource.class);

        ensureResponseSuccess(result);

        return result.getBody();
    }

    @Override
    public Resource updateResource(Resource resource) {
        HttpResponse<Resource> result = Unirest.post(Urls.resourceById(url, Objects.toString(resource.getId()))).body(resource).headers(headers()).asObject(Resource.class);

        ensureResponseSuccess(result);

        return result.getBody();
    }

    @Override
    public void deleteResource(Resource resource) {
        HttpResponse<?> result = Unirest.delete(Urls.resourceById(url, Objects.toString(resource.getId()))).headers(headers()).asEmpty();

        ensureResponseSuccess(result);
    }

    @Override
    public void authenticateWithUsernameAndPassword(String username, String password) {
        Map<String, String> body = new HashMap<>();

        body.put("username", username);
        body.put("password", password);
        body.put("term", Token.TokenTerm.VERY_LONG.name());

        HttpResponse<Token.Container> tokenResponse = Unirest.post(Urls.authenticate(url)).body(body).asObject(Token.Container.class);

        if (tokenResponse.getStatus() == 200) {
            this.token = tokenResponse.getBody().getToken().getContent();
        } else {
            ensureResponseSuccess(tokenResponse);
        }
    }

    @Override
    public Client newClientAuthenticateWithToken(String token) {
        ClientImpl client = new ClientImpl(url);
        client.bypassExtensions = bypassExtensions;
        client.authenticateWithToken(token);

        return client;
    }

    @Override
    public Client newClientAuthenticateWithUsernameAndPassword(String username, String password) {
        ClientImpl client = new ClientImpl(url);
        client.bypassExtensions = bypassExtensions;
        client.authenticateWithUsernameAndPassword(username, password);

        return client;
    }

    @Override
    public <T extends Entity> Repository<T> repo(Class<T> entityClass) {
        return repository(entityClass);
    }

    @Override
    @SuppressWarnings("unchecked")
    public synchronized <T extends Entity> Repository<T> repository(Class<T> entityClass) {
        if (!repositoryClassMap.containsKey(entityClass)) {
            repositoryClassMap.put(entityClass, new RepositoryImpl<>(this, entityClass));
        }

        return (Repository<T>) repositoryClassMap.get(entityClass);
    }

    @Override
    public <T extends Entity> Repository<T> repo(EntityInfo<T> entityInfo) {
        return repository(entityInfo);
    }

    @Override
    @SuppressWarnings("unchecked")
    public <T extends Entity> Repository<T> repository(EntityInfo<T> entityInfo) {
        if (!repositoryEntityMap.containsKey(entityInfo)) {
            repositoryEntityMap.put(entityInfo, new RepositoryImpl<>(this, entityInfo));
        }

        return (Repository<T>) repositoryEntityMap.get(entityInfo);
    }

    public <T extends Entity> Container<T> listRecordsBasic(EntityInfo<T> entityInfo, ListRecordParams params) {
        JavaType type = objectMapper.getTypeFactory().constructCollectionType(List.class, entityInfo.getEntityClass());

        StringBuilder finalUrl = new StringBuilder(Urls.recordUrl(url, entityInfo.getRestPath()));

        // fixme build url properly
        finalUrl.append("?");

        if (params.getLimit() > 0) {
            finalUrl.append("limit=").append(params.getLimit()).append("&");
        }

        if (params.getOffset() > 0) {
            finalUrl.append("offset=").append(params.getOffset()).append("&");
        }

        if (params.getFilters() != null) {
            for (Map.Entry<String, String> entry : params.getFilters().entrySet()) {
                finalUrl.append(entry.getKey()).append("=").append(entry.getValue()).append("&");
            }
        }

        if (params.getResolveReferences() != null && !params.getResolveReferences().isEmpty()) {
            finalUrl.append("resolve-references=").append(String.join(",", params.getResolveReferences())).append("&");
        }

        HttpResponse<Container<T>> result = Unirest.get(finalUrl.toString()).headers(headers()).asObject(resp -> {
            try {
                JsonNode json = objectMapper.readTree(resp.getContent());

                List<T> list = objectMapper.readValue(json.get("content").toString(), type);

                Container<T> container = new Container<>();

                container.setContent(list);

                container.setTotal(json.get("total").asInt());

                return container;
            } catch (IOException e) {
                throw new RuntimeException(e);
            }
        });

        ensureResponseSuccess(result);

        return result.getBody();
    }

    @Override
    public <T extends Entity> Container<T> listRecords(EntityInfo<T> entityInfo, ListRecordParams params) {
        if (params == null) {
            params = ListRecordParams.builder().build();
        }
        if (params.getQuery() == null) {
            return listRecordsBasic(entityInfo, params);
        }
        JavaType type = objectMapper.getTypeFactory().constructCollectionType(List.class, entityInfo.getEntityClass());

        HttpResponse<Container<T>> result = Unirest.post(Urls.recordSearchUrl(url, entityInfo.getRestPath()))
                .body(params)
                .headers(headers())
                .asObject(resp -> {
                    if (resp == null) {
                        return null;
                    }

                    try {
                        JsonNode json = objectMapper.readTree(resp.getContent());

                        List<T> list = objectMapper.readValue(json.get("content").toString(), type);

                        Container<T> container = new Container<>();

                        container.setContent(list);

                        container.setTotal(json.get("total").asInt());

                        return container;
                    } catch (IOException e) {
                        throw new RuntimeException(e);
                    }
                });

        ensureResponseSuccess(result);

        return result.getBody();
    }

    @Override
    public <T extends Entity> T applyRecord(EntityInfo<T> entityInfo, T record) {
        HttpResponse<T> result = Unirest.patch(Urls.recordUrl(url, entityInfo.getRestPath())).body(record).headers(headers()).asObject(entityInfo.getEntityClass());

        ensureResponseSuccess(result);

        return result.getBody();
    }

    @Override
    public <T extends Entity> T deleteRecord(EntityInfo<T> entityInfo, String id) {
        HttpResponse<T> result = Unirest.delete(Urls.recordByIdUrl(url, entityInfo.getRestPath(), id)).headers(headers()).asObject(entityInfo.getEntityClass());

        ensureResponseSuccess(result);

        return result.getBody();
    }

    @Override
    public <T extends Entity> T updateRecord(EntityInfo<T> entityInfo, T record) {
        HttpResponse<T> result = Unirest.put(Urls.recordByIdUrl(url, entityInfo.getRestPath(), Objects.toString(record.getId()))).body(record).headers(headers()).asObject(entityInfo.getEntityClass());

        ensureResponseSuccess(result);

        return result.getBody();
    }

    @Override
    public <T extends Entity> T getRecord(EntityInfo<T> entityInfo, GetRecordParams params) {
        String finalUrl = Urls.recordByIdUrl(url, entityInfo.getRestPath(), params.getId());
        if (params.getResolveReferences() != null && !params.getResolveReferences().isEmpty()) {
            finalUrl += "?resolve-references=" + String.join(",", params.getResolveReferences());
        }

        HttpResponse<T> result = Unirest.get(finalUrl).headers(headers()).asObject(entityInfo.getEntityClass());

        ensureResponseSuccess(result);

        return result.getBody();
    }

    @Override
    public <T extends Entity> T loadRecord(EntityInfo<T> entityInfo, T record) {
        if (record.getId() != null) {
            return this.getRecord(entityInfo, GetRecordParams.builder().id(record.getId().toString()).build());
        }

        Extension.BooleanExpression expression = BooleanExpressionBuilder.and(record.getPropertyMap()
                .entrySet()
                .stream()
                .map(entry -> BooleanExpressionBuilder.eq(entry.getKey(), entry.getValue()))
                .collect(Collectors.toList()));

        Container<T> result = this.listRecords(entityInfo, ListRecordParams.builder().query(expression).build());

        if (result.getTotal() == 0) {
            throw new ApiException(Extension.Code.RECORD_NOT_FOUND, "No record found for " + record);
        } else if (result.getTotal() > 1) {
            throw new ApiException(Extension.Code.RECORD_VALIDATION_ERROR, "Multiple records found for " + record);
        }

        return result.getContent().get(0);
    }

    @Override
    @SneakyThrows
    public <T extends Entity, ActionRequest, ActionResponse> ActionResponse executeRecordAction(Class<ActionResponse> responseClass, EntityInfo<T> entityInfo, String id, String actionName, ActionRequest request) {
        byte[] body = objectMapper.writeValueAsBytes(request);

        log.debug("Executing record action: {} / {} / {}", entityInfo.getEntityClass(), id, actionName);
        HttpResponse<ActionResponse> result = Unirest.post(Urls.recordActionByIdUrl(url, entityInfo.getRestPath(), id, actionName))
                .body(body)
                .headers(headers())
                .asObject(responseClass);

        ensureResponseSuccess(result);

        log.debug("Executed record action: {} / {} / {}", entityInfo.getEntityClass(), id, actionName);

        return result.getBody();
    }

    @Override
    public <T extends Entity> T createRecord(EntityInfo<T> entityInfo, T record) {
        log.debug("Creating record: {}", record);
        HttpResponse<T> result = Unirest.post(Urls.recordUrl(url, entityInfo.getRestPath())).body(record).headers(headers()).asObject(entityInfo.getEntityClass());

        ensureResponseSuccess(result);

        log.debug("Created record: {}", result.getBody());

        return result.getBody();
    }

    @SneakyThrows
    @Override
    public void writeEvent(String channelKey, Extension.Event event) {
        log.debug("Sending event: {}", shortInfo(event));
        log.trace("Sending event[body]: {}", objectMapper.writeValueAsString(event));

        HttpResponse<?> result = Unirest.post(Urls.eventsUrl(url) + "?channelKey=" + channelKey).body(event).headers(headers()).asEmpty();

        ensureResponseSuccess(result);

        log.debug("Sent event: {}", shortInfo(event));
    }

    @Override
    public void bypassExtensions(boolean bypassExtensions) {
        this.bypassExtensions = bypassExtensions;
    }

    public static <T, R> R ensureResponseSuccess(HttpResponse<T> result) {
        if (result.getStatus() != 200) {
            ApiException ex;
            try {
                String errorString = result.mapError(String.class);
                if (errorString == null) {
                    ex = new ApiException(errorString + " " + result.getStatus());
                } else {
                    Extension.Error error = objectMapper.readValue(errorString, Extension.Error.class);
                    ex = new ApiException(error);
                }
            } catch (Exception ignored) {
                log.error("Error parsing error response: {}", ignored.getMessage(), ignored);
                ex = new ApiException(result.getStatusText() + ":" + result.mapError(String.class));
            }

            throw ex;
        }

        if (result.getParsingError().isPresent()) {
            throw new ApiException(result.getParsingError().get().getMessage());
        }

        return null;
    }

    public Map<String, String> headers() {
        HashMap<String, String> headers = new HashMap<>();

        headers.put("Authorization", "Bearer " + token);
        headers.put("Content-Type", "application/json");

        if (bypassExtensions) {
            headers.put("BypassExtensions", "true");
        }

        return headers;
    }
}
