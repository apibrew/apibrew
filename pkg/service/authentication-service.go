package service

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/resources/mapping"
	"github.com/apibrew/apibrew/pkg/service/security"
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	"os"
	"time"
)

type authenticationService struct {
	recordService         abs.RecordService
	privateKey            *rsa.PrivateKey
	publicKey             *rsa.PublicKey
	DisableAuthentication bool
}

func (s *authenticationService) AuthenticationDisabled() bool {
	return s.DisableAuthentication
}

func (s *authenticationService) Authenticate(ctx context.Context, username string, password string, term model.TokenTerm) (*model.Token, errors.ServiceError) {
	if s.DisableAuthentication {
		return &model.Token{
			Term:       term,
			Content:    "",
			Expiration: timestamppb.New(time.Now().Add(time.Minute)),
		}, nil
	}
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Debug("Begin Authenticate")

	defer logger.Debug("End Authenticate")

	// locate user
	user, err := s.LocateUser(security.WithSystemContext(ctx), username, password)

	if err != nil {
		return nil, errors.AuthenticationFailedError
	}

	// Prepare token
	expiration := s.ExpirationFromTerm(term)
	token, err := security.JwtUserDetailsSign(security.JwtUserDetailsSignParams{
		Key: *s.privateKey,
		UserDetails: abs.UserDetails{
			Username: user.Username,
			//SecurityContext: user.SecurityContext,
		},
		ExpiresAt: expiration,
		Issuer:    "github.com/apibrew/apibrew",
	})

	logger.Tracef("Token prepared: %s", token)

	if err != nil {
		logger.Warning("Token preparation error", err)
		return nil, err
	}

	return &model.Token{
		Term:       term,
		Content:    token,
		Expiration: timestamppb.New(expiration),
	}, nil
}

func (s *authenticationService) RenewToken(ctx context.Context, oldToken string, term model.TokenTerm) (*model.Token, errors.ServiceError) {
	userDetails, err := security.JwtVerifyAndUnpackUserDetails(*s.publicKey, oldToken)

	if err != nil {
		return nil, err
	}

	user, err := s.FindUser(ctx, userDetails.Username)

	if err != nil {
		return nil, err
	}

	// Prepare token
	expiration := s.ExpirationFromTerm(term)

	newToken, err := security.JwtUserDetailsSign(security.JwtUserDetailsSignParams{
		Key: *s.privateKey,
		UserDetails: abs.UserDetails{
			Username: user.Username,
			//SecurityContext: user.SecurityContext,
		},
		ExpiresAt: expiration,
		Issuer:    "github.com/apibrew/apibrew",
	})

	if err != nil {
		return nil, err
	}

	return &model.Token{
		Term:       term,
		Content:    newToken,
		Expiration: timestamppb.New(expiration),
	}, nil
}

func (s *authenticationService) ParseAndVerifyToken(token string) (*abs.UserDetails, errors.ServiceError) {
	return security.JwtVerifyAndUnpackUserDetails(*s.publicKey, token)
}

type RequestWithToken interface {
	GetToken() string
}

func (s *authenticationService) LocateUser(ctx context.Context, username, password string) (*model.User, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Debugf("Locating user: %s", username)

	user, err := s.FindUser(ctx, username)
	if err != nil {
		logger.Debugf("Could not find user: %s", username)
		return nil, err
	}

	logger.Debugf("Checking password: %s", username)
	if security.VerifyKey(user.Password, password) != nil {
		logger.Debugf("Password is wrong: %s", username)
		return nil, errors.AuthenticationFailedError
	}

	return user, nil
}

func (s *authenticationService) FindUser(ctx context.Context, username string) (*model.User, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	logger.Debug("FindUser with username: ", username)

	res, err := s.recordService.FindBy(ctx, resources.UserResource.Namespace, resources.UserResource.Name, "username", username)

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

	if data.Config.JwtPrivateKey == "" {
		priv, err := rsa.GenerateKey(rand.Reader, 2048)

		if err != nil {
			panic(err)
		}

		s.privateKey = priv
		s.publicKey = &priv.PublicKey
	} else {
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

func NewAuthenticationService(recordService abs.RecordService) abs.AuthenticationService {
	return &authenticationService{
		recordService: recordService,
	}
}
