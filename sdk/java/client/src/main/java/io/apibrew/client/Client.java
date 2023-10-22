package io.apibrew.client;

import io.apibrew.client.impl.ClientImpl;
import io.apibrew.client.model.Extension;
import io.apibrew.client.model.Resource;

import java.util.List;
import java.util.Map;
import java.util.function.Consumer;
import java.util.function.Predicate;

public interface Client {

    static Client newClient() {
        return ClientImpl.newClient();
    }

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

    <T extends Entity> Container<T> listRecords(EntityInfo<T> entityInfo);

    <T extends Entity> Container<T> listRecords(EntityInfo<T> entityInfo, Extension.BooleanExpression query);

    <T extends Entity> T applyRecord(EntityInfo<T> entityInfo, T record);

    <T extends Entity> T deleteRecord(EntityInfo<T> entityInfo, String id);

    <T extends Entity> T updateRecord(EntityInfo<T> entityInfo, T record);

    <T extends Entity> T getRecord(EntityInfo<T> entityInfo, String id);

    <T extends Entity, ActionRequest, ActionResponse> ActionResponse executeRecordAction(Class<ActionResponse> responseClass, EntityInfo<T> entityInfo, String id, String actionName, ActionRequest request);

    <T extends Entity> T createRecord(EntityInfo<T> entityInfo, T record);

    void writeEvent(String channelKey, Extension.Event event);

    void bypassExtensions(boolean bypassExtensions);

    Map<String, String> headers();

    String getUrl();
}
