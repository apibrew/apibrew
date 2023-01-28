package grpc

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/data-handler/logging"
	"github.com/tislib/data-handler/server/stub"
	"github.com/tislib/data-handler/server/util"
	"github.com/tislib/data-handler/service"
)

type authenticationServiceServer struct {
	stub.AuthenticationServiceServer
	service service.AuthenticationService
}

type RequestWithToken interface {
	GetToken() string
}

func (s *authenticationServiceServer) Authenticate(ctx context.Context, req *stub.AuthenticationRequest) (*stub.AuthenticationResponse, error) {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Debug("Begin Authenticate")
	logger.WithField("req", req).Trace("Params")

	token, err := s.service.Authenticate(ctx, req.Username, req.Password, req.Term)

	logger.Debug("End Authenticate")

	return &stub.AuthenticationResponse{
		Token: token,
	}, util.ToStatusError(err)
}

func (s *authenticationServiceServer) RenewToken(ctx context.Context, req *stub.RenewTokenRequest) (*stub.RenewTokenResponse, error) {
	token, err := s.service.RenewToken(ctx, req.Token, req.Term)

	return &stub.RenewTokenResponse{
		Token: token,
	}, util.ToStatusError(err)
}

func NewAuthenticationServiceServer(service service.AuthenticationService) stub.AuthenticationServiceServer {
	return &authenticationServiceServer{service: service}
}
