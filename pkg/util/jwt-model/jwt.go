package jwt_model

import (
	"crypto/rsa"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

type JwtUserDetailsSignParams struct {
	Key         rsa.PrivateKey
	UserDetails UserDetails
	ExpiresAt   time.Time
	Issuer      string
}

func JwtUserDetailsSign(params JwtUserDetailsSignParams) (string, errors.ServiceError) {
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
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedToken, err := token.SignedString(&params.Key)

	if err != nil {
		log.Error(err)

		return "", errors.InternalError.WithDetails(err.Error())
	}

	return signedToken, nil
}

func JwtVerifyAndUnpackUserDetails(key rsa.PublicKey, tokenContent string) (*UserDetails, errors.ServiceError) {
	claims := new(JwtUserClaims)

	_, err := jwt.ParseWithClaims(tokenContent, claims, func(token *jwt.Token) (interface{}, error) {
		return &key, nil
	})

	if err != nil {
		log.Error(err)

		return nil, errors.InternalError.WithDetails(err.Error())
	}

	return &UserDetails{
		UserId:      claims.UserId,
		Username:    claims.Username,
		Roles:       claims.Roles,
		Permissions: claims.Permissions,
	}, nil
}
