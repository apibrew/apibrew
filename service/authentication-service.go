package service

import (
	"data-handler/stub"
)

type AuthenticationService interface {
	stub.AuthenticationServiceServer
	validateToken(token string) error
}

type authenticationService struct {
	stub.AuthenticationServiceServer
}

func (s authenticationService) validateToken(token string) error {
	return nil
}

func NewAuthenticationService() AuthenticationService {
	return &authenticationService{}
}
