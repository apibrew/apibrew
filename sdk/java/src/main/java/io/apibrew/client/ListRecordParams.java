package io.apibrew.client;

import io.apibrew.client.model.Extension;
import lombok.Builder;
import lombok.Getter;
import lombok.Singular;

import java.util.List;
import java.util.Map;

@Getter
@Builder
public class ListRecordParams {

    // you can either use filters or query, not both
    private final Extension.BooleanExpression query;

    // you can either use filters or query, not both
    private final Map<String, String> filters;

    @Singular
    private final List<String> resolveReferences;

    private final long offset;
    private final int limit;
    private final boolean useHistory;
}
