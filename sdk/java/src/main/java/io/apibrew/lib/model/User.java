package io.apibrew.lib.model;

import java.util.Objects;
import io.apibrew.lib.EntityInfo;
import io.apibrew.lib.Entity;

public class User extends Entity {
    private java.util.UUID id;
    private int version;
    private String createdBy;
    private String updatedBy;
    private java.time.LocalDateTime createdOn;
    private java.time.LocalDateTime updatedOn;
    private String username;
    private String password;
    private java.util.List<Role> roles;
    private java.util.List<Permission> permissions;
    private Object details;

    public static final EntityInfo<User> entityInfo = new EntityInfo<>("system", "User", User.class);



    public User() {
    }

    public EntityInfo<User> getEntityInfo() {
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
    public java.time.LocalDateTime getCreatedOn() {
        return createdOn;
    }

    public void setCreatedOn(java.time.LocalDateTime createdOn) {
        this.createdOn = createdOn;
    }
    public java.time.LocalDateTime getUpdatedOn() {
        return updatedOn;
    }

    public void setUpdatedOn(java.time.LocalDateTime updatedOn) {
        this.updatedOn = updatedOn;
    }
    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }
    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }
    public java.util.List<Role> getRoles() {
        return roles;
    }

    public void setRoles(java.util.List<Role> roles) {
        this.roles = roles;
    }
    public java.util.List<Permission> getPermissions() {
        return permissions;
    }

    public void setPermissions(java.util.List<Permission> permissions) {
        this.permissions = permissions;
    }
    public Object getDetails() {
        return details;
    }

    public void setDetails(Object details) {
        this.details = details;
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


