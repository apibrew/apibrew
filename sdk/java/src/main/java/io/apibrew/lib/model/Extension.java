package io.apibrew.common.model;

import java.util.Objects;
import io.apibrew.common.EntityInfo;
import io.apibrew.common.Entity;
import com.fasterxml.jackson.annotation.JsonValue;
import com.fasterxml.jackson.annotation.JsonFormat;

public class Extension extends Entity {
    private java.util.UUID id;
    private int version;
    private String createdBy;
    private String updatedBy;
    private java.time.Instant createdOn;
    private java.time.Instant updatedOn;
    private String name;
    private String description;
    private Extension.EventSelector selector;
    private int order;
    private boolean finalizes;
    private boolean sync;
    private boolean responds;
    private Extension.ExternalCall call;
    private java.util.Map<String, String> annotations;

    public static final EntityInfo<Extension> entityInfo = new EntityInfo<>("system", "Extension", Extension.class, "system-extension");

    public static class BooleanExpression {

    }
    public static class FunctionCall {
        
        private String host;
        
        private String functionName;

        public String getHost() {
            return host;
        }

        public void setHost(String host) {
            this.host = host;
        }
        public String getFunctionName() {
            return functionName;
        }

        public void setFunctionName(String functionName) {
            this.functionName = functionName;
        }
    }
    public static class HttpCall {
        
        private String uri;
        
        private String method;

        public String getUri() {
            return uri;
        }

        public void setUri(String uri) {
            this.uri = uri;
        }
        public String getMethod() {
            return method;
        }

        public void setMethod(String method) {
            this.method = method;
        }
    }
    public static class ExternalCall {
        
        private Extension.FunctionCall functionCall;
        
        private Extension.HttpCall httpCall;

        public Extension.FunctionCall getFunctionCall() {
            return functionCall;
        }

        public void setFunctionCall(Extension.FunctionCall functionCall) {
            this.functionCall = functionCall;
        }
        public Extension.HttpCall getHttpCall() {
            return httpCall;
        }

        public void setHttpCall(Extension.HttpCall httpCall) {
            this.httpCall = httpCall;
        }
    }
    public static class EventSelector {
        
        private java.util.List<Extension.Action> actions;
        
        private Extension.BooleanExpression recordSelector;
        
        private java.util.List<String> namespaces;
        
        private java.util.List<String> resources;
        
        private java.util.List<String> ids;
        
        private java.util.Map<String, String> annotations;

        public java.util.List<Extension.Action> getActions() {
            return actions;
        }

        public void setActions(java.util.List<Extension.Action> actions) {
            this.actions = actions;
        }
        public Extension.BooleanExpression getRecordSelector() {
            return recordSelector;
        }

        public void setRecordSelector(Extension.BooleanExpression recordSelector) {
            this.recordSelector = recordSelector;
        }
        public java.util.List<String> getNamespaces() {
            return namespaces;
        }

        public void setNamespaces(java.util.List<String> namespaces) {
            this.namespaces = namespaces;
        }
        public java.util.List<String> getResources() {
            return resources;
        }

        public void setResources(java.util.List<String> resources) {
            this.resources = resources;
        }
        public java.util.List<String> getIds() {
            return ids;
        }

        public void setIds(java.util.List<String> ids) {
            this.ids = ids;
        }
        public java.util.Map<String, String> getAnnotations() {
            return annotations;
        }

        public void setAnnotations(java.util.Map<String, String> annotations) {
            this.annotations = annotations;
        }
    }
    public static class RecordSearchParams {
        
        private Extension.BooleanExpression query;
        
        private Integer limit;
        
        private Integer offset;
        
        private java.util.List<String> resolveReferences;

        public Extension.BooleanExpression getQuery() {
            return query;
        }

        public void setQuery(Extension.BooleanExpression query) {
            this.query = query;
        }
        public Integer getLimit() {
            return limit;
        }

        public void setLimit(Integer limit) {
            this.limit = limit;
        }
        public Integer getOffset() {
            return offset;
        }

        public void setOffset(Integer offset) {
            this.offset = offset;
        }
        public java.util.List<String> getResolveReferences() {
            return resolveReferences;
        }

        public void setResolveReferences(java.util.List<String> resolveReferences) {
            this.resolveReferences = resolveReferences;
        }
    }
    public static class Event {
        
        private String id;
        
        private Extension.Action action;
        
        private Extension.RecordSearchParams recordSearchParams;
        
        private String actionSummary;
        
        private String actionDescription;
        
        private Resource resource;
        
        private java.util.List<Record> records;
        
        private java.util.List<String> ids;
        
        private Boolean finalizes;
        
        private Boolean sync;
        @JsonFormat(shape = JsonFormat.Shape.STRING, timezone = "UTC")
        private java.time.Instant time;
        
        private java.util.Map<String, String> annotations;

        public String getId() {
            return id;
        }

        public void setId(String id) {
            this.id = id;
        }
        public Extension.Action getAction() {
            return action;
        }

        public void setAction(Extension.Action action) {
            this.action = action;
        }
        public Extension.RecordSearchParams getRecordSearchParams() {
            return recordSearchParams;
        }

        public void setRecordSearchParams(Extension.RecordSearchParams recordSearchParams) {
            this.recordSearchParams = recordSearchParams;
        }
        public String getActionSummary() {
            return actionSummary;
        }

        public void setActionSummary(String actionSummary) {
            this.actionSummary = actionSummary;
        }
        public String getActionDescription() {
            return actionDescription;
        }

        public void setActionDescription(String actionDescription) {
            this.actionDescription = actionDescription;
        }
        public Resource getResource() {
            return resource;
        }

        public void setResource(Resource resource) {
            this.resource = resource;
        }
        public java.util.List<Record> getRecords() {
            return records;
        }

        public void setRecords(java.util.List<Record> records) {
            this.records = records;
        }
        public java.util.List<String> getIds() {
            return ids;
        }

        public void setIds(java.util.List<String> ids) {
            this.ids = ids;
        }
        public Boolean getFinalizes() {
            return finalizes;
        }

        public void setFinalizes(Boolean finalizes) {
            this.finalizes = finalizes;
        }
        public Boolean getSync() {
            return sync;
        }

        public void setSync(Boolean sync) {
            this.sync = sync;
        }
        public java.time.Instant getTime() {
            return time;
        }

        public void setTime(java.time.Instant time) {
            this.time = time;
        }
        public java.util.Map<String, String> getAnnotations() {
            return annotations;
        }

        public void setAnnotations(java.util.Map<String, String> annotations) {
            this.annotations = annotations;
        }
    }

    public static enum Action {
        CREATE("CREATE"),
        UPDATE("UPDATE"),
        DELETE("DELETE"),
        GET("GET"),
        LIST("LIST"),
        OPERATE("OPERATE");

        private final String value;

        Action(String value) {
            this.value = value;
        }

        @JsonValue
        public String getValue() {
            return value;
        }
    }

    public Extension() {
    }

    public EntityInfo<Extension> getEntityInfo() {
        return entityInfo;
    }

    public java.util.UUID getId() {
        return id;
    }

    public void setId(java.util.UUID id) {
        this.id = id;
    }
    public int getVersion() {
        return version;
    }

    public void setVersion(int version) {
        this.version = version;
    }
    public String getCreatedBy() {
        return createdBy;
    }

    public void setCreatedBy(String createdBy) {
        this.createdBy = createdBy;
    }
    public String getUpdatedBy() {
        return updatedBy;
    }

    public void setUpdatedBy(String updatedBy) {
        this.updatedBy = updatedBy;
    }
    public java.time.Instant getCreatedOn() {
        return createdOn;
    }

    public void setCreatedOn(java.time.Instant createdOn) {
        this.createdOn = createdOn;
    }
    public java.time.Instant getUpdatedOn() {
        return updatedOn;
    }

    public void setUpdatedOn(java.time.Instant updatedOn) {
        this.updatedOn = updatedOn;
    }
    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }
    public Extension.EventSelector getSelector() {
        return selector;
    }

    public void setSelector(Extension.EventSelector selector) {
        this.selector = selector;
    }
    public int getOrder() {
        return order;
    }

    public void setOrder(int order) {
        this.order = order;
    }
    public boolean getFinalizes() {
        return finalizes;
    }

    public void setFinalizes(boolean finalizes) {
        this.finalizes = finalizes;
    }
    public boolean getSync() {
        return sync;
    }

    public void setSync(boolean sync) {
        this.sync = sync;
    }
    public boolean getResponds() {
        return responds;
    }

    public void setResponds(boolean responds) {
        this.responds = responds;
    }
    public Extension.ExternalCall getCall() {
        return call;
    }

    public void setCall(Extension.ExternalCall call) {
        this.call = call;
    }
    public java.util.Map<String, String> getAnnotations() {
        return annotations;
    }

    public void setAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;
    }

    @Override
    public boolean equals(Object o) {
        if (!(o instanceof Extension)) {
            return false;
        }

        Extension obj = (Extension) o;

        if (!Objects.equals(this.id, obj.id)) {
            return false;
        }
        if (!Objects.equals(this.version, obj.version)) {
            return false;
        }
        if (!Objects.equals(this.createdBy, obj.createdBy)) {
            return false;
        }
        if (!Objects.equals(this.updatedBy, obj.updatedBy)) {
            return false;
        }
        if (!Objects.equals(this.createdOn, obj.createdOn)) {
            return false;
        }
        if (!Objects.equals(this.updatedOn, obj.updatedOn)) {
            return false;
        }
        if (!Objects.equals(this.name, obj.name)) {
            return false;
        }
        if (!Objects.equals(this.description, obj.description)) {
            return false;
        }
        if (!Objects.equals(this.selector, obj.selector)) {
            return false;
        }
        if (!Objects.equals(this.order, obj.order)) {
            return false;
        }
        if (!Objects.equals(this.finalizes, obj.finalizes)) {
            return false;
        }
        if (!Objects.equals(this.sync, obj.sync)) {
            return false;
        }
        if (!Objects.equals(this.responds, obj.responds)) {
            return false;
        }
        if (!Objects.equals(this.call, obj.call)) {
            return false;
        }
        if (!Objects.equals(this.annotations, obj.annotations)) {
            return false;
        }

        return true;
    }

    @Override
    public int hashCode() {
        if (id == null) {
            return super.hashCode();
        }

        return id.hashCode();
    }
}


