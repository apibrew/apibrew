package io.apibrew.client.nano.impl;

import io.apibrew.client.Client;
import io.apibrew.client.EntityInfo;
import io.apibrew.client.nano.model.Code;
import io.apibrew.client.nano.NanoService;
import lombok.RequiredArgsConstructor;

import java.util.Base64;

@RequiredArgsConstructor
public class NanoServiceImpl implements NanoService {
    private final Client client;

    @Override
    public Code deploy(Code code) {
        return client.repo(Code.class).create(code);
    }

    @Override
    public Code deploy(String name, Code.Language language, String source) {
        return deploy(name, language, source, false);
    }

    @Override
    public Code deploy(String name, Code.Language language, String source, boolean override) {
        return deploy(new Code().withName(name).withLanguage(language).withContent(Base64.getEncoder().encodeToString(source.getBytes())), override);
    }

    @Override
    public Code deploy(Code code, boolean override) {
        if (override) {
            return client.repo(Code.class).apply(code);
        } else {
            return client.repo(Code.class).create(code);
        }
    }

    @Override
    public Code unDeploy(Code code) {
        if (code.getId() != null) {
            return client.repo(Code.class).delete(code.getId().toString());
        } else {
            Code loadedCode = client.loadRecord(EntityInfo.fromEntityClass(Code.class), new Code().withName(code.getName()));

            return client.repo(Code.class).delete(loadedCode.getId().toString());
        }
    }
}
