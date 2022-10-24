package service

import (
	"context"
	"crypto/rsa"
	"data-handler/service/errors"
	"data-handler/service/mapping"
	"data-handler/service/security"
	"data-handler/service/system"
	"data-handler/stub"
	"data-handler/stub/model"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"os"
	"time"
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

type AuthenticationServiceInternal interface {
	validateToken(token string) error
}

type AuthenticationService interface {
	stub.AuthenticationServiceServer
	validateToken(token string) error
	GrpcIntercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error)
	InjectRecordService(service RecordService)
	Init(data *model.InitData)
}

type authenticationService struct {
	stub.AuthenticationServiceServer
	recordService         RecordServiceInternal
	privateKey            *rsa.PrivateKey
	publicKey             *rsa.PublicKey
	DisableAuthentication bool
}

type RequestWithToken interface {
	GetToken() string
}

func (s *authenticationService) InjectRecordService(service RecordService) {
	s.recordService = service
}

func (s *authenticationService) GrpcIntercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if rtw, ok := req.(RequestWithToken); !s.DisableAuthentication && ok {
		token := rtw.GetToken()

		userDetails, err := security.JwtVerifyAndUnpackUserDetails(*s.publicKey, token)

		if err != nil {
			return nil, errors.AuthenticationFailedError
		}

		userCtx := security.WithUserDetails(ctx, *userDetails)

		return handler(userCtx, req)
	}
	return handler(ctx, req)
}

func (s *authenticationService) Check(params CheckParams) error {
	return nil
}

func (s *authenticationService) validateToken(token string) error {
	return nil
}

func (s *authenticationService) Authenticate(ctx context.Context, req *stub.AuthenticationRequest) (*stub.AuthenticationResponse, error) {
	// locate user
	user, err := s.LocateUser(ctx, req.Username, req.Password)

	if err != nil {
		return &stub.AuthenticationResponse{
			Error: toProtoError(err),
		}, nil
	}

	// Prepare token
	expiration := s.ExpirationFromTerm(req.Term)
	token, err := security.JwtUserDetailsSign(security.JwtUserDetailsSignParams{
		Key: *s.privateKey,
		UserDetails: security.UserDetails{
			Username: user.Username,
			Scopes:   user.Scopes,
		},
		ExpiresAt: expiration,
		Issuer:    "data-handler",
	})

	if err != nil {
		return nil, err
	}

	return &stub.AuthenticationResponse{
		Token: &model.Token{
			Term:       req.Term,
			Content:    token,
			Expiration: timestamppb.New(expiration),
		},
		Error: nil,
	}, nil
}

func (s *authenticationService) RenewToken(ctx context.Context, req *stub.RenewTokenRequest) (*stub.RenewTokenResponse, error) {
	userDetails, err := security.JwtVerifyAndUnpackUserDetails(*s.publicKey, req.Token)

	if err != nil {
		return &stub.RenewTokenResponse{
			Error: toProtoError(err),
		}, nil
	}

	user, err := s.FindUser(ctx, userDetails.Username)

	// Prepare token
	expiration := s.ExpirationFromTerm(req.Term)
	token, err := security.JwtUserDetailsSign(security.JwtUserDetailsSignParams{
		Key: *s.privateKey,
		UserDetails: security.UserDetails{
			Username: user.Username,
			Scopes:   user.Scopes,
		},
		ExpiresAt: expiration,
		Issuer:    "data-handler",
	})

	if err != nil {
		return nil, err
	}

	return &stub.RenewTokenResponse{
		Token: &model.Token{
			Term:       req.Term,
			Content:    token,
			Expiration: timestamppb.New(expiration),
		},
		Error: nil,
	}, nil
}

func (s *authenticationService) LocateUser(ctx context.Context, username, password string) (*model.User, error) {
	user, err := s.FindUser(ctx, username)
	if err != nil {
		return nil, err
	}

	if security.VerifyKey(user.Password, password) != nil {
		return nil, errors.AuthenticationFailedError
	}

	return user, nil
}

func (s *authenticationService) FindUser(ctx context.Context, username string) (*model.User, error) {
	res, err := s.recordService.FindBy(ctx, system.UserResource.Workspace, system.UserResource.Name, "username", username)

	if err != nil {
		return nil, err
	}

	return mapping.UserFromRecord(res), nil
}

func (s *authenticationService) Init(data *model.InitData) {
	s.DisableAuthentication = data.Config.GetDisableAuthentication()

	if data.Config.DisableAuthentication {
		return
	}

	privateKeyContent, err := os.ReadFile(data.Config.JwtPrivateKey)
	if err != nil {
		panic(err)
	}
	publicKeyContent, err := os.ReadFile(data.Config.JwtPublicKey)
	if err != nil {
		panic(err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyContent)
	if err != nil {
		panic(err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyContent)
	if err != nil {
		panic(err)
	}

	s.privateKey = privateKey
	s.publicKey = publicKey
}

func (s *authenticationService) ExpirationFromTerm(term model.TokenTerm) time.Time {
	switch term {
	case model.TokenTerm_SHORT:
		return time.Now().Add(time.Minute)
	case model.TokenTerm_MIDDLE:
		return time.Now().Add(2 * time.Hour)
	case model.TokenTerm_LONG:
		return time.Now().Add(2 * 24 * time.Hour)
	case model.TokenTerm_VERY_LONG:
		return time.Now().Add(2 * 365 * 24 * time.Hour)
	default:
		panic("unknown token term:" + term.String())
	}
}

func NewAuthenticationService() AuthenticationService {
	return &authenticationService{}
}
