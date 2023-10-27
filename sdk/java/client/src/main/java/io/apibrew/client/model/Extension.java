package io.apibrew.client.model;

import java.util.Objects;
import io.apibrew.client.EntityInfo;
import io.apibrew.client.Entity;
import io.apibrew.client.Client;
import com.fasterxml.jackson.annotation.JsonValue;
import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class Extension extends Entity {
    
    private java.util.UUID id;
    
    private int version;
    
    private Extension.AuditData auditData;
    
    private String name;
    
    private String description;
    
    private Extension.EventSelector selector;
    
    private int order;
    
    private boolean finalizes;
    
    private boolean sync;
    
    private boolean responds;
    
    private Extension.ExternalCall call;
    
    private java.util.Map<String, String> annotations;

    public static final String NAMESPACE = "system";
    public static final String RESOURCE = "Extension";

    @JsonIgnore
    public static final EntityInfo<Extension> entityInfo = new EntityInfo<>("system", "Extension", Extension.class, "system-extension");

    public static class BooleanExpression {
        
        private java.util.List<Extension.BooleanExpression> and;
        
        private java.util.List<Extension.BooleanExpression> or;
        
        private Extension.BooleanExpression not;
        
        private Extension.PairExpression equal;
        
        private Extension.PairExpression lessThan;
        
        private Extension.PairExpression greaterThan;
        
        private Extension.PairExpression lessThanOrEqual;
        
        private Extension.PairExpression greaterThanOrEqual;
        
        private Extension.PairExpression in;
        
        private Extension.Expression isNull;
        
        private Extension.RegexMatchExpression regexMatch;

        public java.util.List<Extension.BooleanExpression> getAnd() {
            return and;
        }

        public void setAnd(java.util.List<Extension.BooleanExpression> and) {
            this.and = and;
        }

        public BooleanExpression withAnd(java.util.List<Extension.BooleanExpression> and) {
            this.and = and;

            return this;
        }
        public java.util.List<Extension.BooleanExpression> getOr() {
            return or;
        }

        public void setOr(java.util.List<Extension.BooleanExpression> or) {
            this.or = or;
        }

        public BooleanExpression withOr(java.util.List<Extension.BooleanExpression> or) {
            this.or = or;

            return this;
        }
        public Extension.BooleanExpression getNot() {
            return not;
        }

        public void setNot(Extension.BooleanExpression not) {
            this.not = not;
        }

        public BooleanExpression withNot(Extension.BooleanExpression not) {
            this.not = not;

            return this;
        }
        public Extension.PairExpression getEqual() {
            return equal;
        }

        public void setEqual(Extension.PairExpression equal) {
            this.equal = equal;
        }

        public BooleanExpression withEqual(Extension.PairExpression equal) {
            this.equal = equal;

            return this;
        }
        public Extension.PairExpression getLessThan() {
            return lessThan;
        }

        public void setLessThan(Extension.PairExpression lessThan) {
            this.lessThan = lessThan;
        }

        public BooleanExpression withLessThan(Extension.PairExpression lessThan) {
            this.lessThan = lessThan;

            return this;
        }
        public Extension.PairExpression getGreaterThan() {
            return greaterThan;
        }

        public void setGreaterThan(Extension.PairExpression greaterThan) {
            this.greaterThan = greaterThan;
        }

        public BooleanExpression withGreaterThan(Extension.PairExpression greaterThan) {
            this.greaterThan = greaterThan;

            return this;
        }
        public Extension.PairExpression getLessThanOrEqual() {
            return lessThanOrEqual;
        }

        public void setLessThanOrEqual(Extension.PairExpression lessThanOrEqual) {
            this.lessThanOrEqual = lessThanOrEqual;
        }

        public BooleanExpression withLessThanOrEqual(Extension.PairExpression lessThanOrEqual) {
            this.lessThanOrEqual = lessThanOrEqual;

            return this;
        }
        public Extension.PairExpression getGreaterThanOrEqual() {
            return greaterThanOrEqual;
        }

        public void setGreaterThanOrEqual(Extension.PairExpression greaterThanOrEqual) {
            this.greaterThanOrEqual = greaterThanOrEqual;
        }

        public BooleanExpression withGreaterThanOrEqual(Extension.PairExpression greaterThanOrEqual) {
            this.greaterThanOrEqual = greaterThanOrEqual;

            return this;
        }
        public Extension.PairExpression getIn() {
            return in;
        }

        public void setIn(Extension.PairExpression in) {
            this.in = in;
        }

        public BooleanExpression withIn(Extension.PairExpression in) {
            this.in = in;

            return this;
        }
        public Extension.Expression getIsNull() {
            return isNull;
        }

        public void setIsNull(Extension.Expression isNull) {
            this.isNull = isNull;
        }

        public BooleanExpression withIsNull(Extension.Expression isNull) {
            this.isNull = isNull;

            return this;
        }
        public Extension.RegexMatchExpression getRegexMatch() {
            return regexMatch;
        }

        public void setRegexMatch(Extension.RegexMatchExpression regexMatch) {
            this.regexMatch = regexMatch;
        }

        public BooleanExpression withRegexMatch(Extension.RegexMatchExpression regexMatch) {
            this.regexMatch = regexMatch;

            return this;
        }

        @Override
        public boolean equals(Object o) {
            if (!(o instanceof BooleanExpression)) {
                return false;
            }

            BooleanExpression obj = (BooleanExpression) o;

            if (!Objects.equals(this.and, obj.and)) {
                return false;
            }
            if (!Objects.equals(this.or, obj.or)) {
                return false;
            }
            if (!Objects.equals(this.not, obj.not)) {
                return false;
            }
            if (!Objects.equals(this.equal, obj.equal)) {
                return false;
            }
            if (!Objects.equals(this.lessThan, obj.lessThan)) {
                return false;
            }
            if (!Objects.equals(this.greaterThan, obj.greaterThan)) {
                return false;
            }
            if (!Objects.equals(this.lessThanOrEqual, obj.lessThanOrEqual)) {
                return false;
            }
            if (!Objects.equals(this.greaterThanOrEqual, obj.greaterThanOrEqual)) {
                return false;
            }
            if (!Objects.equals(this.in, obj.in)) {
                return false;
            }
            if (!Objects.equals(this.isNull, obj.isNull)) {
                return false;
            }
            if (!Objects.equals(this.regexMatch, obj.regexMatch)) {
                return false;
            }

            return true;
        }

        @Override
        public int hashCode() {
           return Objects.hash(and, or, not, equal, lessThan, greaterThan, lessThanOrEqual, greaterThanOrEqual, in, isNull, regexMatch);
        }
    }
    public static class PairExpression {
        
        private Extension.Expression left;
        
        private Extension.Expression right;

        public Extension.Expression getLeft() {
            return left;
        }

        public void setLeft(Extension.Expression left) {
            this.left = left;
        }

        public PairExpression withLeft(Extension.Expression left) {
            this.left = left;

            return this;
        }
        public Extension.Expression getRight() {
            return right;
        }

        public void setRight(Extension.Expression right) {
            this.right = right;
        }

        public PairExpression withRight(Extension.Expression right) {
            this.right = right;

            return this;
        }

        @Override
        public boolean equals(Object o) {
            if (!(o instanceof PairExpression)) {
                return false;
            }

            PairExpression obj = (PairExpression) o;

            if (!Objects.equals(this.left, obj.left)) {
                return false;
            }
            if (!Objects.equals(this.right, obj.right)) {
                return false;
            }

            return true;
        }

        @Override
        public int hashCode() {
           return Objects.hash(left, right);
        }
    }
    public static class RegexMatchExpression {
        
        private String pattern;
        
        private Extension.Expression expression;

        public String getPattern() {
            return pattern;
        }

        public void setPattern(String pattern) {
            this.pattern = pattern;
        }

        public RegexMatchExpression withPattern(String pattern) {
            this.pattern = pattern;

            return this;
        }
        public Extension.Expression getExpression() {
            return expression;
        }

        public void setExpression(Extension.Expression expression) {
            this.expression = expression;
        }

        public RegexMatchExpression withExpression(Extension.Expression expression) {
            this.expression = expression;

            return this;
        }

        @Override
        public boolean equals(Object o) {
            if (!(o instanceof RegexMatchExpression)) {
                return false;
            }

            RegexMatchExpression obj = (RegexMatchExpression) o;

            if (!Objects.equals(this.pattern, obj.pattern)) {
                return false;
            }
            if (!Objects.equals(this.expression, obj.expression)) {
                return false;
            }

            return true;
        }

        @Override
        public int hashCode() {
           return Objects.hash(pattern, expression);
        }
    }
    public static class Expression {
        
        private String property;
        
        private Object value;

        public String getProperty() {
            return property;
        }

        public void setProperty(String property) {
            this.property = property;
        }

        public Expression withProperty(String property) {
            this.property = property;

            return this;
        }
        public Object getValue() {
            return value;
        }

        public void setValue(Object value) {
            this.value = value;
        }

        public Expression withValue(Object value) {
            this.value = value;

            return this;
        }

        @Override
        public boolean equals(Object o) {
            if (!(o instanceof Expression)) {
                return false;
            }

            Expression obj = (Expression) o;

            if (!Objects.equals(this.property, obj.property)) {
                return false;
            }
            if (!Objects.equals(this.value, obj.value)) {
                return false;
            }

            return true;
        }

        @Override
        public int hashCode() {
           return Objects.hash(property, value);
        }
    }
    public static class AuditData {
        
        private String createdBy;
        
        private String updatedBy;
        @JsonFormat(shape = JsonFormat.Shape.STRING, timezone = "UTC")
        private java.time.Instant createdOn;
        @JsonFormat(shape = JsonFormat.Shape.STRING, timezone = "UTC")
        private java.time.Instant updatedOn;

        public String getCreatedBy() {
            return createdBy;
        }

        public void setCreatedBy(String createdBy) {
            this.createdBy = createdBy;
        }

        public AuditData withCreatedBy(String createdBy) {
            this.createdBy = createdBy;

            return this;
        }
        public String getUpdatedBy() {
            return updatedBy;
        }

        public void setUpdatedBy(String updatedBy) {
            this.updatedBy = updatedBy;
        }

        public AuditData withUpdatedBy(String updatedBy) {
            this.updatedBy = updatedBy;

            return this;
        }
        public java.time.Instant getCreatedOn() {
            return createdOn;
        }

        public void setCreatedOn(java.time.Instant createdOn) {
            this.createdOn = createdOn;
        }

        public AuditData withCreatedOn(java.time.Instant createdOn) {
            this.createdOn = createdOn;

            return this;
        }
        public java.time.Instant getUpdatedOn() {
            return updatedOn;
        }

        public void setUpdatedOn(java.time.Instant updatedOn) {
            this.updatedOn = updatedOn;
        }

        public AuditData withUpdatedOn(java.time.Instant updatedOn) {
            this.updatedOn = updatedOn;

            return this;
        }

        @Override
        public boolean equals(Object o) {
            if (!(o instanceof AuditData)) {
                return false;
            }

            AuditData obj = (AuditData) o;

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

            return true;
        }

        @Override
        public int hashCode() {
           return Objects.hash(createdBy, updatedBy, createdOn, updatedOn);
        }
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

        public FunctionCall withHost(String host) {
            this.host = host;

            return this;
        }
        public String getFunctionName() {
            return functionName;
        }

        public void setFunctionName(String functionName) {
            this.functionName = functionName;
        }

        public FunctionCall withFunctionName(String functionName) {
            this.functionName = functionName;

            return this;
        }

        @Override
        public boolean equals(Object o) {
            if (!(o instanceof FunctionCall)) {
                return false;
            }

            FunctionCall obj = (FunctionCall) o;

            if (!Objects.equals(this.host, obj.host)) {
                return false;
            }
            if (!Objects.equals(this.functionName, obj.functionName)) {
                return false;
            }

            return true;
        }

        @Override
        public int hashCode() {
           return Objects.hash(host, functionName);
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

        public HttpCall withUri(String uri) {
            this.uri = uri;

            return this;
        }
        public String getMethod() {
            return method;
        }

        public void setMethod(String method) {
            this.method = method;
        }

        public HttpCall withMethod(String method) {
            this.method = method;

            return this;
        }

        @Override
        public boolean equals(Object o) {
            if (!(o instanceof HttpCall)) {
                return false;
            }

            HttpCall obj = (HttpCall) o;

            if (!Objects.equals(this.uri, obj.uri)) {
                return false;
            }
            if (!Objects.equals(this.method, obj.method)) {
                return false;
            }

            return true;
        }

        @Override
        public int hashCode() {
           return Objects.hash(uri, method);
        }
    }
    public static class ChannelCall {
        
        private String channelKey;

        public String getChannelKey() {
            return channelKey;
        }

        public void setChannelKey(String channelKey) {
            this.channelKey = channelKey;
        }

        public ChannelCall withChannelKey(String channelKey) {
            this.channelKey = channelKey;

            return this;
        }

        @Override
        public boolean equals(Object o) {
            if (!(o instanceof ChannelCall)) {
                return false;
            }

            ChannelCall obj = (ChannelCall) o;

            if (!Objects.equals(this.channelKey, obj.channelKey)) {
                return false;
            }

            return true;
        }

        @Override
        public int hashCode() {
           return Objects.hash(channelKey);
        }
    }
    public static class ExternalCall {
        
        private Extension.FunctionCall functionCall;
        
        private Extension.HttpCall httpCall;
        
        private Extension.ChannelCall channelCall;

        public Extension.FunctionCall getFunctionCall() {
            return functionCall;
        }

        public void setFunctionCall(Extension.FunctionCall functionCall) {
            this.functionCall = functionCall;
        }

        public ExternalCall withFunctionCall(Extension.FunctionCall functionCall) {
            this.functionCall = functionCall;

            return this;
        }
        public Extension.HttpCall getHttpCall() {
            return httpCall;
        }

        public void setHttpCall(Extension.HttpCall httpCall) {
            this.httpCall = httpCall;
        }

        public ExternalCall withHttpCall(Extension.HttpCall httpCall) {
            this.httpCall = httpCall;

            return this;
        }
        public Extension.ChannelCall getChannelCall() {
            return channelCall;
        }

        public void setChannelCall(Extension.ChannelCall channelCall) {
            this.channelCall = channelCall;
        }

        public ExternalCall withChannelCall(Extension.ChannelCall channelCall) {
            this.channelCall = channelCall;

            return this;
        }

        @Override
        public boolean equals(Object o) {
            if (!(o instanceof ExternalCall)) {
                return false;
            }

            ExternalCall obj = (ExternalCall) o;

            if (!Objects.equals(this.functionCall, obj.functionCall)) {
                return false;
            }
            if (!Objects.equals(this.httpCall, obj.httpCall)) {
                return false;
            }
            if (!Objects.equals(this.channelCall, obj.channelCall)) {
                return false;
            }

            return true;
        }

        @Override
        public int hashCode() {
           return Objects.hash(functionCall, httpCall, channelCall);
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

        public EventSelector withActions(java.util.List<Extension.Action> actions) {
            this.actions = actions;

            return this;
        }
        public Extension.BooleanExpression getRecordSelector() {
            return recordSelector;
        }

        public void setRecordSelector(Extension.BooleanExpression recordSelector) {
            this.recordSelector = recordSelector;
        }

        public EventSelector withRecordSelector(Extension.BooleanExpression recordSelector) {
            this.recordSelector = recordSelector;

            return this;
        }
        public java.util.List<String> getNamespaces() {
            return namespaces;
        }

        public void setNamespaces(java.util.List<String> namespaces) {
            this.namespaces = namespaces;
        }

        public EventSelector withNamespaces(java.util.List<String> namespaces) {
            this.namespaces = namespaces;

            return this;
        }
        public java.util.List<String> getResources() {
            return resources;
        }

        public void setResources(java.util.List<String> resources) {
            this.resources = resources;
        }

        public EventSelector withResources(java.util.List<String> resources) {
            this.resources = resources;

            return this;
        }
        public java.util.List<String> getIds() {
            return ids;
        }

        public void setIds(java.util.List<String> ids) {
            this.ids = ids;
        }

        public EventSelector withIds(java.util.List<String> ids) {
            this.ids = ids;

            return this;
        }
        public java.util.Map<String, String> getAnnotations() {
            return annotations;
        }

        public void setAnnotations(java.util.Map<String, String> annotations) {
            this.annotations = annotations;
        }

        public EventSelector withAnnotations(java.util.Map<String, String> annotations) {
            this.annotations = annotations;

            return this;
        }

        @Override
        public boolean equals(Object o) {
            if (!(o instanceof EventSelector)) {
                return false;
            }

            EventSelector obj = (EventSelector) o;

            if (!Objects.equals(this.actions, obj.actions)) {
                return false;
            }
            if (!Objects.equals(this.recordSelector, obj.recordSelector)) {
                return false;
            }
            if (!Objects.equals(this.namespaces, obj.namespaces)) {
                return false;
            }
            if (!Objects.equals(this.resources, obj.resources)) {
                return false;
            }
            if (!Objects.equals(this.ids, obj.ids)) {
                return false;
            }
            if (!Objects.equals(this.annotations, obj.annotations)) {
                return false;
            }

            return true;
        }

        @Override
        public int hashCode() {
           return Objects.hash(actions, recordSelector, namespaces, resources, ids, annotations);
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

        public RecordSearchParams withQuery(Extension.BooleanExpression query) {
            this.query = query;

            return this;
        }
        public Integer getLimit() {
            return limit;
        }

        public void setLimit(Integer limit) {
            this.limit = limit;
        }

        public RecordSearchParams withLimit(Integer limit) {
            this.limit = limit;

            return this;
        }
        public Integer getOffset() {
            return offset;
        }

        public void setOffset(Integer offset) {
            this.offset = offset;
        }

        public RecordSearchParams withOffset(Integer offset) {
            this.offset = offset;

            return this;
        }
        public java.util.List<String> getResolveReferences() {
            return resolveReferences;
        }

        public void setResolveReferences(java.util.List<String> resolveReferences) {
            this.resolveReferences = resolveReferences;
        }

        public RecordSearchParams withResolveReferences(java.util.List<String> resolveReferences) {
            this.resolveReferences = resolveReferences;

            return this;
        }

        @Override
        public boolean equals(Object o) {
            if (!(o instanceof RecordSearchParams)) {
                return false;
            }

            RecordSearchParams obj = (RecordSearchParams) o;

            if (!Objects.equals(this.query, obj.query)) {
                return false;
            }
            if (!Objects.equals(this.limit, obj.limit)) {
                return false;
            }
            if (!Objects.equals(this.offset, obj.offset)) {
                return false;
            }
            if (!Objects.equals(this.resolveReferences, obj.resolveReferences)) {
                return false;
            }

            return true;
        }

        @Override
        public int hashCode() {
           return Objects.hash(query, limit, offset, resolveReferences);
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
        
        private Boolean finalizes;
        
        private Boolean sync;
        @JsonFormat(shape = JsonFormat.Shape.STRING, timezone = "UTC")
        private java.time.Instant time;
        
        private Long total;
        
        private String actionName;
        
        private Object input;
        
        private Object output;
        
        private java.util.Map<String, String> annotations;
        
        private Extension.Error error;

        public String getId() {
            return id;
        }

        public void setId(String id) {
            this.id = id;
        }

        public Event withId(String id) {
            this.id = id;

            return this;
        }
        public Extension.Action getAction() {
            return action;
        }

        public void setAction(Extension.Action action) {
            this.action = action;
        }

        public Event withAction(Extension.Action action) {
            this.action = action;

            return this;
        }
        public Extension.RecordSearchParams getRecordSearchParams() {
            return recordSearchParams;
        }

        public void setRecordSearchParams(Extension.RecordSearchParams recordSearchParams) {
            this.recordSearchParams = recordSearchParams;
        }

        public Event withRecordSearchParams(Extension.RecordSearchParams recordSearchParams) {
            this.recordSearchParams = recordSearchParams;

            return this;
        }
        public String getActionSummary() {
            return actionSummary;
        }

        public void setActionSummary(String actionSummary) {
            this.actionSummary = actionSummary;
        }

        public Event withActionSummary(String actionSummary) {
            this.actionSummary = actionSummary;

            return this;
        }
        public String getActionDescription() {
            return actionDescription;
        }

        public void setActionDescription(String actionDescription) {
            this.actionDescription = actionDescription;
        }

        public Event withActionDescription(String actionDescription) {
            this.actionDescription = actionDescription;

            return this;
        }
        public Resource getResource() {
            return resource;
        }

        public void setResource(Resource resource) {
            this.resource = resource;
        }

        public Event withResource(Resource resource) {
            this.resource = resource;

            return this;
        }
        public java.util.List<Record> getRecords() {
            return records;
        }

        public void setRecords(java.util.List<Record> records) {
            this.records = records;
        }

        public Event withRecords(java.util.List<Record> records) {
            this.records = records;

            return this;
        }
        public Boolean getFinalizes() {
            return finalizes;
        }

        public void setFinalizes(Boolean finalizes) {
            this.finalizes = finalizes;
        }

        public Event withFinalizes(Boolean finalizes) {
            this.finalizes = finalizes;

            return this;
        }
        public Boolean getSync() {
            return sync;
        }

        public void setSync(Boolean sync) {
            this.sync = sync;
        }

        public Event withSync(Boolean sync) {
            this.sync = sync;

            return this;
        }
        public java.time.Instant getTime() {
            return time;
        }

        public void setTime(java.time.Instant time) {
            this.time = time;
        }

        public Event withTime(java.time.Instant time) {
            this.time = time;

            return this;
        }
        public Long getTotal() {
            return total;
        }

        public void setTotal(Long total) {
            this.total = total;
        }

        public Event withTotal(Long total) {
            this.total = total;

            return this;
        }
        public String getActionName() {
            return actionName;
        }

        public void setActionName(String actionName) {
            this.actionName = actionName;
        }

        public Event withActionName(String actionName) {
            this.actionName = actionName;

            return this;
        }
        public Object getInput() {
            return input;
        }

        public void setInput(Object input) {
            this.input = input;
        }

        public Event withInput(Object input) {
            this.input = input;

            return this;
        }
        public Object getOutput() {
            return output;
        }

        public void setOutput(Object output) {
            this.output = output;
        }

        public Event withOutput(Object output) {
            this.output = output;

            return this;
        }
        public java.util.Map<String, String> getAnnotations() {
            return annotations;
        }

        public void setAnnotations(java.util.Map<String, String> annotations) {
            this.annotations = annotations;
        }

        public Event withAnnotations(java.util.Map<String, String> annotations) {
            this.annotations = annotations;

            return this;
        }
        public Extension.Error getError() {
            return error;
        }

        public void setError(Extension.Error error) {
            this.error = error;
        }

        public Event withError(Extension.Error error) {
            this.error = error;

            return this;
        }

        @Override
        public boolean equals(Object o) {
            if (!(o instanceof Event)) {
                return false;
            }

            Event obj = (Event) o;

            if (!Objects.equals(this.id, obj.id)) {
                return false;
            }
            if (!Objects.equals(this.action, obj.action)) {
                return false;
            }
            if (!Objects.equals(this.recordSearchParams, obj.recordSearchParams)) {
                return false;
            }
            if (!Objects.equals(this.actionSummary, obj.actionSummary)) {
                return false;
            }
            if (!Objects.equals(this.actionDescription, obj.actionDescription)) {
                return false;
            }
            if (!Objects.equals(this.resource, obj.resource)) {
                return false;
            }
            if (!Objects.equals(this.records, obj.records)) {
                return false;
            }
            if (!Objects.equals(this.finalizes, obj.finalizes)) {
                return false;
            }
            if (!Objects.equals(this.sync, obj.sync)) {
                return false;
            }
            if (!Objects.equals(this.time, obj.time)) {
                return false;
            }
            if (!Objects.equals(this.total, obj.total)) {
                return false;
            }
            if (!Objects.equals(this.actionName, obj.actionName)) {
                return false;
            }
            if (!Objects.equals(this.input, obj.input)) {
                return false;
            }
            if (!Objects.equals(this.output, obj.output)) {
                return false;
            }
            if (!Objects.equals(this.annotations, obj.annotations)) {
                return false;
            }
            if (!Objects.equals(this.error, obj.error)) {
                return false;
            }

            return true;
        }

        @Override
        public int hashCode() {
           return Objects.hash(id, action, recordSearchParams, actionSummary, actionDescription, resource, records, finalizes, sync, time, total, actionName, input, output, annotations, error);
        }
    }
    public static class ErrorField {
        
        private String recordId;
        
        private String property;
        
        private String message;
        
        private Object value;

        public String getRecordId() {
            return recordId;
        }

        public void setRecordId(String recordId) {
            this.recordId = recordId;
        }

        public ErrorField withRecordId(String recordId) {
            this.recordId = recordId;

            return this;
        }
        public String getProperty() {
            return property;
        }

        public void setProperty(String property) {
            this.property = property;
        }

        public ErrorField withProperty(String property) {
            this.property = property;

            return this;
        }
        public String getMessage() {
            return message;
        }

        public void setMessage(String message) {
            this.message = message;
        }

        public ErrorField withMessage(String message) {
            this.message = message;

            return this;
        }
        public Object getValue() {
            return value;
        }

        public void setValue(Object value) {
            this.value = value;
        }

        public ErrorField withValue(Object value) {
            this.value = value;

            return this;
        }

        @Override
        public boolean equals(Object o) {
            if (!(o instanceof ErrorField)) {
                return false;
            }

            ErrorField obj = (ErrorField) o;

            if (!Objects.equals(this.recordId, obj.recordId)) {
                return false;
            }
            if (!Objects.equals(this.property, obj.property)) {
                return false;
            }
            if (!Objects.equals(this.message, obj.message)) {
                return false;
            }
            if (!Objects.equals(this.value, obj.value)) {
                return false;
            }

            return true;
        }

        @Override
        public int hashCode() {
           return Objects.hash(recordId, property, message, value);
        }
    }
    public static class Error {
        
        private Extension.Code code;
        
        private String message;
        
        private java.util.List<Extension.ErrorField> fields;

        public Extension.Code getCode() {
            return code;
        }

        public void setCode(Extension.Code code) {
            this.code = code;
        }

        public Error withCode(Extension.Code code) {
            this.code = code;

            return this;
        }
        public String getMessage() {
            return message;
        }

        public void setMessage(String message) {
            this.message = message;
        }

        public Error withMessage(String message) {
            this.message = message;

            return this;
        }
        public java.util.List<Extension.ErrorField> getFields() {
            return fields;
        }

        public void setFields(java.util.List<Extension.ErrorField> fields) {
            this.fields = fields;
        }

        public Error withFields(java.util.List<Extension.ErrorField> fields) {
            this.fields = fields;

            return this;
        }

        @Override
        public boolean equals(Object o) {
            if (!(o instanceof Error)) {
                return false;
            }

            Error obj = (Error) o;

            if (!Objects.equals(this.code, obj.code)) {
                return false;
            }
            if (!Objects.equals(this.message, obj.message)) {
                return false;
            }
            if (!Objects.equals(this.fields, obj.fields)) {
                return false;
            }

            return true;
        }

        @Override
        public int hashCode() {
           return Objects.hash(code, message, fields);
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
    public static enum Code {
        UNKNOWN_ERROR("UNKNOWN_ERROR"),
        RECORD_NOT_FOUND("RECORD_NOT_FOUND"),
        UNABLE_TO_LOCATE_PRIMARY_KEY("UNABLE_TO_LOCATE_PRIMARY_KEY"),
        INTERNAL_ERROR("INTERNAL_ERROR"),
        PROPERTY_NOT_FOUND("PROPERTY_NOT_FOUND"),
        RECORD_VALIDATION_ERROR("RECORD_VALIDATION_ERROR"),
        RESOURCE_VALIDATION_ERROR("RESOURCE_VALIDATION_ERROR"),
        AUTHENTICATION_FAILED("AUTHENTICATION_FAILED"),
        ALREADY_EXISTS("ALREADY_EXISTS"),
        ACCESS_DENIED("ACCESS_DENIED"),
        BACKEND_ERROR("BACKEND_ERROR"),
        UNIQUE_VIOLATION("UNIQUE_VIOLATION"),
        REFERENCE_VIOLATION("REFERENCE_VIOLATION"),
        RESOURCE_NOT_FOUND("RESOURCE_NOT_FOUND"),
        UNSUPPORTED_OPERATION("UNSUPPORTED_OPERATION"),
        EXTERNAL_BACKEND_COMMUNICATION_ERROR("EXTERNAL_BACKEND_COMMUNICATION_ERROR"),
        EXTERNAL_BACKEND_ERROR("EXTERNAL_BACKEND_ERROR"),
        RATE_LIMIT_ERROR("RATE_LIMIT_ERROR");

        private final String value;

        Code(String value) {
            this.value = value;
        }

        @JsonValue
        public String getValue() {
            return value;
        }
    }

    

    public Extension() {
    }

    public java.util.UUID getId() {
        return id;
    }

    public void setId(java.util.UUID id) {
        this.id = id;
    }

    public Extension withId(java.util.UUID id) {
        this.id = id;

        return this;
    }
    public int getVersion() {
        return version;
    }

    public void setVersion(int version) {
        this.version = version;
    }

    public Extension withVersion(int version) {
        this.version = version;

        return this;
    }
    public Extension.AuditData getAuditData() {
        return auditData;
    }

    public void setAuditData(Extension.AuditData auditData) {
        this.auditData = auditData;
    }

    public Extension withAuditData(Extension.AuditData auditData) {
        this.auditData = auditData;

        return this;
    }
    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Extension withName(String name) {
        this.name = name;

        return this;
    }
    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public Extension withDescription(String description) {
        this.description = description;

        return this;
    }
    public Extension.EventSelector getSelector() {
        return selector;
    }

    public void setSelector(Extension.EventSelector selector) {
        this.selector = selector;
    }

    public Extension withSelector(Extension.EventSelector selector) {
        this.selector = selector;

        return this;
    }
    public int getOrder() {
        return order;
    }

    public void setOrder(int order) {
        this.order = order;
    }

    public Extension withOrder(int order) {
        this.order = order;

        return this;
    }
    public boolean getFinalizes() {
        return finalizes;
    }

    public void setFinalizes(boolean finalizes) {
        this.finalizes = finalizes;
    }

    public Extension withFinalizes(boolean finalizes) {
        this.finalizes = finalizes;

        return this;
    }
    public boolean getSync() {
        return sync;
    }

    public void setSync(boolean sync) {
        this.sync = sync;
    }

    public Extension withSync(boolean sync) {
        this.sync = sync;

        return this;
    }
    public boolean getResponds() {
        return responds;
    }

    public void setResponds(boolean responds) {
        this.responds = responds;
    }

    public Extension withResponds(boolean responds) {
        this.responds = responds;

        return this;
    }
    public Extension.ExternalCall getCall() {
        return call;
    }

    public void setCall(Extension.ExternalCall call) {
        this.call = call;
    }

    public Extension withCall(Extension.ExternalCall call) {
        this.call = call;

        return this;
    }
    public java.util.Map<String, String> getAnnotations() {
        return annotations;
    }

    public void setAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;
    }

    public Extension withAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;

        return this;
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
        if (!Objects.equals(this.auditData, obj.auditData)) {
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


