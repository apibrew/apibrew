package io.apibrew.client.storage.model;

import java.util.Objects;
import io.apibrew.client.EntityInfo;
import io.apibrew.client.Entity;
import io.apibrew.client.Client;
import com.fasterxml.jackson.annotation.JsonValue;
import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class Signature extends Entity {
    
    private java.util.UUID id;
    
    private StorageObject object;
    
    private java.util.List<Signature.Permission> permissions;
    @JsonFormat(shape = JsonFormat.Shape.STRING, timezone = "UTC")
    private java.time.Instant expiration;
    
    private String signature;
    
    private java.util.Map<String, String> annotations;
    
    private int version;

    public static final String NAMESPACE = "storage";
    public static final String RESOURCE = "Signature";

    @JsonIgnore
    public static final EntityInfo<Signature> entityInfo = new EntityInfo<>("storage", "Signature", Signature.class, "storage-signature");


    public static enum Permission {
        DOWNLOAD("DOWNLOAD"),
        UPLOAD("UPLOAD");

        private final String value;

        Permission(String value) {
            this.value = value;
        }

        @JsonValue
        public String getValue() {
            return value;
        }
    }

    

    public Signature() {
    }

    public java.util.UUID getId() {
        return id;
    }

    public void setId(java.util.UUID id) {
        this.id = id;
    }

    public Signature withId(java.util.UUID id) {
        this.id = id;

        return this;
    }
    public StorageObject getObject() {
        return object;
    }

    public void setObject(StorageObject object) {
        this.object = object;
    }

    public Signature withObject(StorageObject object) {
        this.object = object;

        return this;
    }
    public java.util.List<Signature.Permission> getPermissions() {
        return permissions;
    }

    public void setPermissions(java.util.List<Signature.Permission> permissions) {
        this.permissions = permissions;
    }

    public Signature withPermissions(java.util.List<Signature.Permission> permissions) {
        this.permissions = permissions;

        return this;
    }
    public java.time.Instant getExpiration() {
        return expiration;
    }

    public void setExpiration(java.time.Instant expiration) {
        this.expiration = expiration;
    }

    public Signature withExpiration(java.time.Instant expiration) {
        this.expiration = expiration;

        return this;
    }
    public String getSignature() {
        return signature;
    }

    public void setSignature(String signature) {
        this.signature = signature;
    }

    public Signature withSignature(String signature) {
        this.signature = signature;

        return this;
    }
    public java.util.Map<String, String> getAnnotations() {
        return annotations;
    }

    public void setAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;
    }

    public Signature withAnnotations(java.util.Map<String, String> annotations) {
        this.annotations = annotations;

        return this;
    }
    public int getVersion() {
        return version;
    }

    public void setVersion(int version) {
        this.version = version;
    }

    public Signature withVersion(int version) {
        this.version = version;

        return this;
    }

    @Override
    public boolean equals(Object o) {
        if (!(o instanceof Signature)) {
            return false;
        }

        Signature obj = (Signature) o;

        if (!Objects.equals(this.id, obj.id)) {
            return false;
        }
        if (!Objects.equals(this.object, obj.object)) {
            return false;
        }
        if (!Objects.equals(this.permissions, obj.permissions)) {
            return false;
        }
        if (!Objects.equals(this.expiration, obj.expiration)) {
            return false;
        }
        if (!Objects.equals(this.signature, obj.signature)) {
            return false;
        }
        if (!Objects.equals(this.annotations, obj.annotations)) {
            return false;
        }
        if (!Objects.equals(this.version, obj.version)) {
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


