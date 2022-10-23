package security

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

type JwtUserDetailsSignParams struct {
	Key         rsa.PrivateKey
	UserDetails UserDetails
	ExpiresAt   time.Time
	Issuer      string
}

type UserDetails struct {
	Username string
	Scopes   []string
}

func JwtUserDetailsSign(params JwtUserDetailsSignParams) (string, error) {
	jit, err := uuid.NewRandom()

	if err != nil {
		return "", err
	}

	claims := &JwtUserClaims{
		Issuer:    params.Issuer,
		Subject:   params.UserDetails.Username,
		Audience:  []string{params.Issuer},
		ExpiresAt: jwt.NewNumericDate(params.ExpiresAt),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        jit.String(),
		Scopes:    params.UserDetails.Scopes,
		Username:  params.UserDetails.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(&params.Key)
}

func JwtVerifyAndUnpackUserDetails(key rsa.PublicKey, tokenContent string) (*UserDetails, error) {
	claims := new(JwtUserClaims)

	_, err := jwt.ParseWithClaims(tokenContent, claims, func(token *jwt.Token) (interface{}, error) {
		return &key, nil
	})

	if err != nil {
		return nil, err
	}

	return &UserDetails{
		Username: claims.Username,
		Scopes:   claims.Scopes,
	}, nil
}
