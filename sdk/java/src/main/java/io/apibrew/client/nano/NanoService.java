package io.apibrew.client.nano;

import io.apibrew.client.nano.model.Code;

public interface NanoService {
    public Code deploy(Code code);
    public Code deploy(String name, Code.Language language, String source);
    public Code deploy(String name, Code.Language language, String source, boolean override);

    public Code deploy(Code code, boolean override);

    public Code unDeploy(Code code);
}
