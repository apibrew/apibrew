package io.apibrew.client.model;

import java.util.Objects;
import io.apibrew.client.EntityInfo;
import io.apibrew.client.Entity;
import com.fasterxml.jackson.annotation.JsonValue;
import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class User extends Entity {
    
    private java.util.UUID id;
    
    private int version;
    
    private String createdBy;
    
    private String updatedBy;
    @JsonFormat(shape = JsonFormat.Shape.STRING, timezone = "UTC")
    private java.time.Instant createdOn;
    @JsonFormat(shape = JsonFormat.Shape.STRING, timezone = "UTC")
    private java.time.Instant updatedOn;
    
    private String username;
    
    private String password;
    
    private java.util.List<Role> roles;
    
    private java.util.List<Permission> permissions;
    
    private Object details;

    @JsonIgnore
    public static final EntityInfo<User> entityInfo = new EntityInfo<>("system", "User", User.class, "system-user");



    public User() {
    }

    public java.util.UUID getId() {
        return id;
    }

    public void setId(java.util.UUID id) {
        this.id = id;
    }

    public User withId(java.util.UUID id) {
        this.id = id;

        return this;
    }
    public int getVersion() {
        return version;
    }

    public void setVersion(int version) {
        this.version = version;
    }

    public User withVersion(int version) {
        this.version = version;

        return this;
    }
    public String getCreatedBy() {
        return createdBy;
    }

    public void setCreatedBy(String createdBy) {
        this.createdBy = createdBy;
    }

    public User withCreatedBy(String createdBy) {
        this.createdBy = createdBy;

        return this;
    }
    public String getUpdatedBy() {
        return updatedBy;
    }

    public void setUpdatedBy(String updatedBy) {
        this.updatedBy = updatedBy;
    }

    public User withUpdatedBy(String updatedBy) {
        this.updatedBy = updatedBy;

        return this;
    }
    public java.time.Instant getCreatedOn() {
        return createdOn;
    }

    public void setCreatedOn(java.time.Instant createdOn) {
        this.createdOn = createdOn;
    }

    public User withCreatedOn(java.time.Instant createdOn) {
        this.createdOn = createdOn;

        return this;
    }
    public java.time.Instant getUpdatedOn() {
        return updatedOn;
    }

    public void setUpdatedOn(java.time.Instant updatedOn) {
        this.updatedOn = updatedOn;
    }

    public User withUpdatedOn(java.time.Instant updatedOn) {
        this.updatedOn = updatedOn;

        return this;
    }
    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public User withUsername(String username) {
        this.username = username;

        return this;
    }
    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public User withPassword(String password) {
        this.password = password;

        return this;
    }
    public java.util.List<Role> getRoles() {
        return roles;
    }

    public void setRoles(java.util.List<Role> roles) {
        this.roles = roles;
    }

    public User withRoles(java.util.List<Role> roles) {
        this.roles = roles;

        return this;
    }
    public java.util.List<Permission> getPermissions() {
        return permissions;
    }

    public void setPermissions(java.util.List<Permission> permissions) {
        this.permissions = permissions;
    }

    public User withPermissions(java.util.List<Permission> permissions) {
        this.permissions = permissions;

        return this;
    }
    public Object getDetails() {
        return details;
    }

    public void setDetails(Object details) {
        this.details = details;
    }

    public User withDetails(Object details) {
        this.details = details;

        return this;
    }

    @Override
    public boolean equals(Object o) {
        if (!(o instanceof User)) {
            return false;
        }

        User obj = (User) o;

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
        if (!Objects.equals(this.username, obj.username)) {
            return false;
        }
        if (!Objects.equals(this.password, obj.password)) {
            return false;
        }
        if (!Objects.equals(this.roles, obj.roles)) {
            return false;
        }
        if (!Objects.equals(this.permissions, obj.permissions)) {
            return false;
        }
        if (!Objects.equals(this.details, obj.details)) {
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


