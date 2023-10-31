package io.apibrew.client;

import lombok.Builder;
import lombok.Getter;
import lombok.Singular;

import java.util.List;

@Builder
@Getter
public class GetRecordParams {

    @Singular
    private final List<String> resolveReferences;

    private final String id;

}
