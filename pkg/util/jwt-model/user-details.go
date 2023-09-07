package jwt_model

import "github.com/apibrew/apibrew/pkg/resource_model"

type UserDetails struct {
	UserId              string
	Username            string
	SecurityConstraints []*resource_model.Permission
}
