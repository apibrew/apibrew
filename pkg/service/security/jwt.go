package security

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/pkg/abs"
	"github.com/tislib/data-handler/pkg/errors"
	"time"
)

type JwtUserDetailsSignParams struct {
	Key         rsa.PrivateKey
	UserDetails abs.UserDetails
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
		Issuer:          params.Issuer,
		Subject:         params.UserDetails.Username,
		Audience:        []string{params.Issuer},
		ExpiresAt:       jwt.NewNumericDate(params.ExpiresAt),
		NotBefore:       jwt.NewNumericDate(time.Now()),
		IssuedAt:        jwt.NewNumericDate(time.Now()),
		ID:              jit.String(),
		SecurityContext: params.UserDetails.SecurityContext,
		Username:        params.UserDetails.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedToken, err := token.SignedString(&params.Key)

	if err != nil {
		log.Error(err)

		return "", errors.InternalError.WithDetails(err.Error())
	}

	return signedToken, nil
}

func JwtVerifyAndUnpackUserDetails(key rsa.PublicKey, tokenContent string) (*abs.UserDetails, errors.ServiceError) {
	claims := new(JwtUserClaims)

	_, err := jwt.ParseWithClaims(tokenContent, claims, func(token *jwt.Token) (interface{}, error) {
		return &key, nil
	})

	if err != nil {
		log.Error(err)

		return nil, errors.InternalError.WithDetails(err.Error())
	}

	return &abs.UserDetails{
		Username:        claims.Username,
		SecurityContext: claims.SecurityContext,
	}, nil
}
