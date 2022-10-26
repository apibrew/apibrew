package grpc_service

import (
	"context"
	"data-handler/grpc/stub"
	"data-handler/model"
	"data-handler/service"
	"data-handler/service/errors"
)

type authenticationServiceServer struct {
	stub.AuthenticationServiceServer
	service service.AuthenticationService
}

type RequestWithToken interface {
	GetToken() string
}

func (s *authenticationServiceServer) Authenticate(ctx context.Context, req *stub.AuthenticationRequest) (*stub.AuthenticationResponse, error) {
	token, err := s.service.Authenticate(ctx, req.Username, req.Password, req.Term)

	return &stub.AuthenticationResponse{
		Token: token,
		Error: toProtoError(err),
	}, nil
}

func toProtoError(err errors.ServiceError) *model.Error {
	if err == nil {
		return nil
	}

	return err.ProtoError()
}

func (s *authenticationServiceServer) RenewToken(ctx context.Context, req *stub.RenewTokenRequest) (*stub.RenewTokenResponse, error) {
	token, err := s.service.RenewToken(ctx, req.Token, req.Term)

	return &stub.RenewTokenResponse{
		Token: token,
		Error: toProtoError(err),
	}, nil
}

func NewAuthenticationServiceServer(service service.AuthenticationService) stub.AuthenticationServiceServer {
	return &authenticationServiceServer{service: service}
}
