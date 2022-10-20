package service

import (
	"context"
	"data-handler/stub"
)

type ResourceIdentifier interface {
}

type CheckParams struct {
	Ctx       context.Context
	Token     string
	Service   string
	Method    string
	Resources any
}

type AuthenticationService interface {
	stub.AuthenticationServiceServer
	validateToken(token string) error
	Check(params CheckParams) error
}

type authenticationService struct {
	stub.AuthenticationServiceServer
}

func (s authenticationService) Check(params CheckParams) error {
	return nil
}

func (s authenticationService) validateToken(token string) error {
	return nil
}

func NewAuthenticationService() AuthenticationService {
	return &authenticationService{}
}
