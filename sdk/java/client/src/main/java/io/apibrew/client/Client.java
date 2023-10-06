package io.apibrew.client;

import io.apibrew.client.impl.ClientImpl;
import io.apibrew.client.model.Extension;
import io.apibrew.client.model.Resource;

import java.util.List;
import java.util.Map;
import java.util.function.Predicate;

public interface Client {

    static Client newClient(String url) {
        return new ClientImpl(url);
    }

    static Client newClientByServerName(String serverName) {
        return ClientImpl.newClientByServerName(serverName);
    }

    static Client newClientByServerConfig(Config.Server serverConfig) {
        return ClientImpl.newClientByServerConfig(serverConfig);
    }

    Resource ApplyResource(Resource resource);

    Resource GetResourceByName(String namespace, String name);

    List<Resource> listResources();

    Resource CreateResource(Resource resource);

    Resource UpdateResource(Resource resource);

    void AuthenticateWithToken(String token);

    void authenticateWithUsernameAndPassword(String username, String password);

    <T extends Entity> Repository<T> repo(Class<T> entityClass);

    <T extends Entity> Repository<T> repository(Class<T> entityClass);

    <T extends Entity> Repository<T> repo(EntityInfo<T> entityInfo);

    <T extends Entity> Repository<T> repository(EntityInfo<T> entityInfo);

    <T extends Entity> Container<T> listRecords(Class<T> entityClass, String namespace, String resource);
    <T extends Entity> Container<T> listRecords(Class<T> entityClass, String namespace, String resource, Extension.BooleanExpression query);

    <T extends Entity> T applyRecord(Class<T> entityClass, String namespace, String resource, T record);

    <T extends Entity> T deleteRecord(Class<T> entityClass, String namespace, String resource, String id);

    <T extends Entity> T updateRecord(Class<T> entityClass, String namespace, String resource, T record);

    <T extends Entity> T getRecord(Class<T> entityClass, String namespace, String resource, String id);

    <T extends Entity> T createRecord(Class<T> entityClass, String namespace, String resource, T record);

    void writeEvent(String channelKey, Extension.Event event);

    void bypassExtensions(boolean bypassExtensions);

    Map<String, String> headers();

    String getUrl();
}
