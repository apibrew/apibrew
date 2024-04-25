package jwt_model

import (
	"crypto/rsa"
	"encoding/json"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type JwtUserDetailsSignParams struct {
	Key         rsa.PrivateKey
	UserDetails UserDetails
	ExpiresAt   time.Time
	Issuer      string
}

func JwtUserDetailsSign(params JwtUserDetailsSignParams, minimizeToken bool) (string, error) {
	jit, err := uuid.NewRandom()

	if err != nil {
		log.Error(err)

		return "", errors.InternalError.WithDetails(err.Error())
	}

	claims := &JwtUserClaims{
		Issuer:      params.Issuer,
		Subject:     params.UserDetails.Username,
		Audience:    []string{params.Issuer},
		ExpiresAt:   jwt.NewNumericDate(params.ExpiresAt),
		NotBefore:   jwt.NewNumericDate(time.Now()),
		IssuedAt:    jwt.NewNumericDate(time.Now()),
		ID:          jit.String(),
		Username:    params.UserDetails.Username,
		Roles:       params.UserDetails.Roles,
		Permissions: params.UserDetails.Permissions,
		UserId:      params.UserDetails.UserId,
		Scopes:      prepareScopes(params.UserDetails.Permissions),
	}

	if minimizeToken {
		claims.Permissions = nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedToken, err := token.SignedString(&params.Key)

	if err != nil {
		log.Error(err)

		return "", errors.InternalError.WithDetails(err.Error())
	}

	return signedToken, nil
}

func prepareScopes(permissions []*resource_model.Permission) []string {
	var scopes []string

	var scopeMap = make(map[string]bool)

	for _, permission := range permissions {
		var scope = permissionToScope(permission)

		if scope == "" {
			continue
		}

		if _, ok := scopeMap[scope]; ok {
			continue
		}

		scopes = append(scopes, scope)

		scopeMap[scope] = true
	}

	return scopes
}

func preparePermissions(scopes []string) []*resource_model.Permission {
	var permissions []*resource_model.Permission

	for _, scope := range scopes {
		permissions = append(permissions, scopeToPermission(scope))
	}

	return permissions
}

func permissionToScope(permission *resource_model.Permission) string {
	var parts = []string{}

	parts = append(parts, util.DePointer(permission.Namespace, ""))
	parts = append(parts, util.DePointer(permission.Resource, ""))
	parts = append(parts, string(permission.Operation))

	if permission.RecordSelector != nil {
		serializedRecordSelector, err := json.Marshal(permission.RecordSelector)

		if err != nil {
			log.Error(err)
		} else {
			parts = append(parts, string(serializedRecordSelector))
		}
	} else {
		parts = append(parts, "")
	}

	if permission.Before != nil {
		parts = append(parts, permission.Before.Format(time.RFC3339))
	} else {
		parts = append(parts, "")
	}

	if permission.After != nil {
		parts = append(parts, permission.After.Format(time.RFC3339))
	} else {
		parts = append(parts, "")
	}

	if permission.User != nil {
		parts = append(parts, permission.User.Username)
	} else {
		parts = append(parts, "")
	}

	if permission.Role != nil {
		parts = append(parts, permission.Role.Name)
	} else {
		parts = append(parts, "")
	}

	parts = append(parts, string(permission.Permit))

	return strings.Join(parts, ":")
}

func scopeToPermission(scope string) *resource_model.Permission {
	parts := strings.Split(scope, ":")

	if len(parts) != 9 {
		return nil
	}

	var result = &resource_model.Permission{}

	if parts[0] != "" {
		result.Namespace = &parts[0]
	}

	if parts[1] != "" {
		result.Resource = &parts[1]
	}

	if parts[2] != "" {
		result.Operation = resource_model.PermissionOperation(parts[2])
	}

	if parts[3] != "" {
		var recordSelector = new(resource_model.BooleanExpression)

		err := json.Unmarshal([]byte(parts[3]), recordSelector)

		if err != nil {
			log.Error(err)
		} else {
			result.RecordSelector = recordSelector
		}
	}

	if parts[4] != "" {
		result.Before = new(time.Time)
		*result.Before, _ = time.Parse(time.RFC3339, parts[4])
	}

	if parts[5] != "" {
		result.After = new(time.Time)
		*result.After, _ = time.Parse(time.RFC3339, parts[5])
	}

	if parts[6] != "" {
		result.User = &resource_model.User{Username: parts[6]}
	}

	if parts[7] != "" {
		result.Role = &resource_model.Role{Name: parts[7]}
	}

	if parts[8] != "" {
		result.Permit = resource_model.PermissionPermit(parts[8])
	}

	return result
}

func JwtVerifyAndUnpackUserDetails(key rsa.PublicKey, tokenContent string) (*UserDetails, error) {
	claims := new(JwtUserClaims)

	_, err := jwt.ParseWithClaims(tokenContent, claims, func(token *jwt.Token) (interface{}, error) {
		return &key, nil
	})

	if err != nil {
		log.Error(err)

		return nil, errors.InternalError.WithDetails(err.Error())
	}

	var permissions = claims.Permissions

	if claims.Permissions == nil {
		permissions = preparePermissions(claims.Scopes)
	}

	return &UserDetails{
		UserId:      claims.UserId,
		Username:    claims.Username,
		Roles:       claims.Roles,
		Permissions: permissions,
	}, nil
}
