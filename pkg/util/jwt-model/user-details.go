package jwt_model

import "github.com/apibrew/apibrew/pkg/resource_model"

type UserDetails struct {
	UserId      string                       `json:"userId"`
	Username    string                       `json:"username"`
	Permissions []*resource_model.Permission `json:"permissions"`
	Roles       []string                     `json:"roles"`
}
