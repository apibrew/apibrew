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
public class Permission extends Entity {
    
    private java.util.UUID id;
    
    private int version;
    
    private Permission.AuditData auditData;
    
    private String namespace;
    
    private String resource;
    
    private Permission.BooleanExpression recordSelector;
    
    private Permission.Operation operation;
    @JsonFormat(shape = JsonFormat.Shape.STRING, timezone = "UTC")
    private java.time.Instant before;
    @JsonFormat(shape = JsonFormat.Shape.STRING, timezone = "UTC")
    private java.time.Instant after;
    
    private User user;
    
    private Role role;
    
    private Permission.Permit permit;
    
    private Object localFlags;

    public static final String NAMESPACE = "system";
    public static final String RESOURCE = "Permission";

    @JsonIgnore
    public static final EntityInfo<Permission> entityInfo = new EntityInfo<>("system", "Permission", Permission.class, "system-permission");

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
    public static class BooleanExpression {
        
        private java.util.List<Permission.BooleanExpression> and;
        
        private java.util.List<Permission.BooleanExpression> or;
        
        private Permission.BooleanExpression not;
        
        private Permission.PairExpression equal;
        
        private Permission.PairExpression lessThan;
        
        private Permission.PairExpression greaterThan;
        
        private Permission.PairExpression lessThanOrEqual;
        
        private Permission.PairExpression greaterThanOrEqual;
        
        private Permission.PairExpression in;
        
        private Permission.Expression isNull;
        
        private Permission.RegexMatchExpression regexMatch;

        public java.util.List<Permission.BooleanExpression> getAnd() {
            return and;
        }

        public void setAnd(java.util.List<Permission.BooleanExpression> and) {
            this.and = and;
        }

        public BooleanExpression withAnd(java.util.List<Permission.BooleanExpression> and) {
            this.and = and;

            return this;
        }
        public java.util.List<Permission.BooleanExpression> getOr() {
            return or;
        }

        public void setOr(java.util.List<Permission.BooleanExpression> or) {
            this.or = or;
        }

        public BooleanExpression withOr(java.util.List<Permission.BooleanExpression> or) {
            this.or = or;

            return this;
        }
        public Permission.BooleanExpression getNot() {
            return not;
        }

        public void setNot(Permission.BooleanExpression not) {
            this.not = not;
        }

        public BooleanExpression withNot(Permission.BooleanExpression not) {
            this.not = not;

            return this;
        }
        public Permission.PairExpression getEqual() {
            return equal;
        }

        public void setEqual(Permission.PairExpression equal) {
            this.equal = equal;
        }

        public BooleanExpression withEqual(Permission.PairExpression equal) {
            this.equal = equal;

            return this;
        }
        public Permission.PairExpression getLessThan() {
            return lessThan;
        }

        public void setLessThan(Permission.PairExpression lessThan) {
            this.lessThan = lessThan;
        }

        public BooleanExpression withLessThan(Permission.PairExpression lessThan) {
            this.lessThan = lessThan;

            return this;
        }
        public Permission.PairExpression getGreaterThan() {
            return greaterThan;
        }

        public void setGreaterThan(Permission.PairExpression greaterThan) {
            this.greaterThan = greaterThan;
        }

        public BooleanExpression withGreaterThan(Permission.PairExpression greaterThan) {
            this.greaterThan = greaterThan;

            return this;
        }
        public Permission.PairExpression getLessThanOrEqual() {
            return lessThanOrEqual;
        }

        public void setLessThanOrEqual(Permission.PairExpression lessThanOrEqual) {
            this.lessThanOrEqual = lessThanOrEqual;
        }

        public BooleanExpression withLessThanOrEqual(Permission.PairExpression lessThanOrEqual) {
            this.lessThanOrEqual = lessThanOrEqual;

            return this;
        }
        public Permission.PairExpression getGreaterThanOrEqual() {
            return greaterThanOrEqual;
        }

        public void setGreaterThanOrEqual(Permission.PairExpression greaterThanOrEqual) {
            this.greaterThanOrEqual = greaterThanOrEqual;
        }

        public BooleanExpression withGreaterThanOrEqual(Permission.PairExpression greaterThanOrEqual) {
            this.greaterThanOrEqual = greaterThanOrEqual;

            return this;
        }
        public Permission.PairExpression getIn() {
            return in;
        }

        public void setIn(Permission.PairExpression in) {
            this.in = in;
        }

        public BooleanExpression withIn(Permission.PairExpression in) {
            this.in = in;

            return this;
        }
        public Permission.Expression getIsNull() {
            return isNull;
        }

        public void setIsNull(Permission.Expression isNull) {
            this.isNull = isNull;
        }

        public BooleanExpression withIsNull(Permission.Expression isNull) {
            this.isNull = isNull;

            return this;
        }
        public Permission.RegexMatchExpression getRegexMatch() {
            return regexMatch;
        }

        public void setRegexMatch(Permission.RegexMatchExpression regexMatch) {
            this.regexMatch = regexMatch;
        }

        public BooleanExpression withRegexMatch(Permission.RegexMatchExpression regexMatch) {
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
        
        private Permission.Expression left;
        
        private Permission.Expression right;

        public Permission.Expression getLeft() {
            return left;
        }

        public void setLeft(Permission.Expression left) {
            this.left = left;
        }

        public PairExpression withLeft(Permission.Expression left) {
            this.left = left;

            return this;
        }
        public Permission.Expression getRight() {
            return right;
        }

        public void setRight(Permission.Expression right) {
            this.right = right;
        }

        public PairExpression withRight(Permission.Expression right) {
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
        
        private Permission.Expression expression;

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
        public Permission.Expression getExpression() {
            return expression;
        }

        public void setExpression(Permission.Expression expression) {
            this.expression = expression;
        }

        public RegexMatchExpression withExpression(Permission.Expression expression) {
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

    public static enum Operation {
        READ("READ"),
        CREATE("CREATE"),
        UPDATE("UPDATE"),
        DELETE("DELETE"),
        FULL("FULL");

        private final String value;

        Operation(String value) {
            this.value = value;
        }

        @JsonValue
        public String getValue() {
            return value;
        }
    }
    public static enum Permit {
        ALLOW("ALLOW"),
        REJECT("REJECT");

        private final String value;

        Permit(String value) {
            this.value = value;
        }

        @JsonValue
        public String getValue() {
            return value;
        }
    }

    

    public Permission() {
    }

    public java.util.UUID getId() {
        return id;
    }

    public void setId(java.util.UUID id) {
        this.id = id;
    }

    public Permission withId(java.util.UUID id) {
        this.id = id;

        return this;
    }
    public int getVersion() {
        return version;
    }

    public void setVersion(int version) {
        this.version = version;
    }

    public Permission withVersion(int version) {
        this.version = version;

        return this;
    }
    public Permission.AuditData getAuditData() {
        return auditData;
    }

    public void setAuditData(Permission.AuditData auditData) {
        this.auditData = auditData;
    }

    public Permission withAuditData(Permission.AuditData auditData) {
        this.auditData = auditData;

        return this;
    }
    public String getNamespace() {
        return namespace;
    }

    public void setNamespace(String namespace) {
        this.namespace = namespace;
    }

    public Permission withNamespace(String namespace) {
        this.namespace = namespace;

        return this;
    }
    public String getResource() {
        return resource;
    }

    public void setResource(String resource) {
        this.resource = resource;
    }

    public Permission withResource(String resource) {
        this.resource = resource;

        return this;
    }
    public Permission.BooleanExpression getRecordSelector() {
        return recordSelector;
    }

    public void setRecordSelector(Permission.BooleanExpression recordSelector) {
        this.recordSelector = recordSelector;
    }

    public Permission withRecordSelector(Permission.BooleanExpression recordSelector) {
        this.recordSelector = recordSelector;

        return this;
    }
    public Permission.Operation getOperation() {
        return operation;
    }

    public void setOperation(Permission.Operation operation) {
        this.operation = operation;
    }

    public Permission withOperation(Permission.Operation operation) {
        this.operation = operation;

        return this;
    }
    public java.time.Instant getBefore() {
        return before;
    }

    public void setBefore(java.time.Instant before) {
        this.before = before;
    }

    public Permission withBefore(java.time.Instant before) {
        this.before = before;

        return this;
    }
    public java.time.Instant getAfter() {
        return after;
    }

    public void setAfter(java.time.Instant after) {
        this.after = after;
    }

    public Permission withAfter(java.time.Instant after) {
        this.after = after;

        return this;
    }
    public User getUser() {
        return user;
    }

    public void setUser(User user) {
        this.user = user;
    }

    public Permission withUser(User user) {
        this.user = user;

        return this;
    }
    public Role getRole() {
        return role;
    }

    public void setRole(Role role) {
        this.role = role;
    }

    public Permission withRole(Role role) {
        this.role = role;

        return this;
    }
    public Permission.Permit getPermit() {
        return permit;
    }

    public void setPermit(Permission.Permit permit) {
        this.permit = permit;
    }

    public Permission withPermit(Permission.Permit permit) {
        this.permit = permit;

        return this;
    }
    public Object getLocalFlags() {
        return localFlags;
    }

    public void setLocalFlags(Object localFlags) {
        this.localFlags = localFlags;
    }

    public Permission withLocalFlags(Object localFlags) {
        this.localFlags = localFlags;

        return this;
    }

    @Override
    public boolean equals(Object o) {
        if (!(o instanceof Permission)) {
            return false;
        }

        Permission obj = (Permission) o;

        if (!Objects.equals(this.id, obj.id)) {
            return false;
        }
        if (!Objects.equals(this.version, obj.version)) {
            return false;
        }
        if (!Objects.equals(this.auditData, obj.auditData)) {
            return false;
        }
        if (!Objects.equals(this.namespace, obj.namespace)) {
            return false;
        }
        if (!Objects.equals(this.resource, obj.resource)) {
            return false;
        }
        if (!Objects.equals(this.recordSelector, obj.recordSelector)) {
            return false;
        }
        if (!Objects.equals(this.operation, obj.operation)) {
            return false;
        }
        if (!Objects.equals(this.before, obj.before)) {
            return false;
        }
        if (!Objects.equals(this.after, obj.after)) {
            return false;
        }
        if (!Objects.equals(this.user, obj.user)) {
            return false;
        }
        if (!Objects.equals(this.role, obj.role)) {
            return false;
        }
        if (!Objects.equals(this.permit, obj.permit)) {
            return false;
        }
        if (!Objects.equals(this.localFlags, obj.localFlags)) {
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


