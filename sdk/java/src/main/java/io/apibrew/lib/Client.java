package io.apibrew.common;

import io.apibrew.common.ext.Extension;
import io.apibrew.common.impl.ClientImpl;
import io.apibrew.common.model.Resource;

import java.util.List;

public interface Client {

    static Client newClient(String url) {
        return new ClientImpl(url);
    }

    static Client newClientByServerName(String serverName) {
        return ClientImpl.newClientByServerName(serverName);
    }

    Resource ApplyResource(Resource resource);

    Resource GetResourceByName(String name);

    List<Resource> ListResources();

    Resource CreateResource(Resource resource);

    Resource UpdateResource(Resource resource);

    void AuthenticateWithToken(String token);

    void AuthenticateWithUsernameAndPassword(String username, String password);

    <T extends Entity> Repository<T> repository(Class<T> entityClass);

    <T extends Entity> Repository<T> repository(EntityInfo<T> entityInfo);

    <T extends Entity> Repository<T> repository(String namespace, String resource, Class<T> entityClass);

    Extension extension();
}
