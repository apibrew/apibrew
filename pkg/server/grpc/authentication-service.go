package grpc

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tislib/apibrew/pkg/abs"
	"github.com/tislib/apibrew/pkg/logging"
	"github.com/tislib/apibrew/pkg/stub"
	"github.com/tislib/apibrew/pkg/util"
)

type authenticationServer struct {
	stub.AuthenticationServer
	service abs.AuthenticationService
}

type RequestWithToken interface {
	GetToken() string
}

func (s *authenticationServer) Authenticate(ctx context.Context, req *stub.AuthenticationRequest) (*stub.AuthenticationResponse, error) {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Debug("Begin Authenticate")
	logger.WithField("req", req).Trace("Params")

	token, err := s.service.Authenticate(ctx, req.Username, req.Password, req.Term)

	logger.Debug("End Authenticate")

	return &stub.AuthenticationResponse{
		Token: token,
	}, util.ToStatusError(err)
}

func (s *authenticationServer) RenewToken(ctx context.Context, req *stub.RenewTokenRequest) (*stub.RenewTokenResponse, error) {
	token, err := s.service.RenewToken(ctx, req.Token, req.Term)

	return &stub.RenewTokenResponse{
		Token: token,
	}, util.ToStatusError(err)
}

func NewAuthenticationServer(service abs.AuthenticationService) stub.AuthenticationServer {
	return &authenticationServer{service: service}
}
