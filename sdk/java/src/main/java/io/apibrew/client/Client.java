package io.apibrew.client;

import io.apibrew.client.impl.ClientImpl;
import io.apibrew.client.model.Extension;
import io.apibrew.client.model.Resource;
import io.apibrew.client.model.Token;

import java.util.List;
import java.util.Map;

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

    Resource applyResource(Resource resource);

    Resource getResourceByName(String namespace, String name);

    List<Resource> listResources();

    Resource createResource(Resource resource);

    Resource updateResource(Resource resource);

    void deleteResource(Resource resource);

    void authenticateWithToken(String token);

    void authenticateWithUsernameAndPassword(String username, String password);
    String authenticateWithUsernameAndPassword(String username, String password, Token.TokenTerm term);

    Client newClientAuthenticateWithToken(String token);

    Client newClientAuthenticateWithUsernameAndPassword(String username, String password);

    <T extends Entity> Repository<T> repo(Class<T> entityClass);

    <T extends Entity> Repository<T> repository(Class<T> entityClass);

    <T extends Entity> Repository<T> repo(EntityInfo<T> entityInfo);

    <T extends Entity> Repository<T> repository(EntityInfo<T> entityInfo);

    <T extends Entity> Container<T> listRecords(EntityInfo<T> entityInfo, ListRecordParams params);

    <T extends Entity> T applyRecord(EntityInfo<T> entityInfo, T record);

    <T extends Entity> T deleteRecord(EntityInfo<T> entityInfo, String id);

    <T extends Entity> T updateRecord(EntityInfo<T> entityInfo, T record);

    <T extends Entity> T getRecord(EntityInfo<T> entityInfo, GetRecordParams params);

    <T extends Entity, ActionRequest, ActionResponse> ActionResponse executeRecordAction(Class<ActionResponse> responseClass, EntityInfo<T> entityInfo, String id, String actionName, ActionRequest request);

    <T extends Entity> T createRecord(EntityInfo<T> entityInfo, T record);

    void writeEvent(String channelKey, Extension.Event event);

    void bypassExtensions(boolean bypassExtensions);

    Map<String, String> headers();

    String getUrl();

    <T extends Entity> T loadRecord(EntityInfo<T> entityInfo, T record);

}
