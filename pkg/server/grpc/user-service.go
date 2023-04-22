package grpc

import (
	"context"
	"github.com/tislib/apibrew/pkg/abs"
	"github.com/tislib/apibrew/pkg/stub"
	"github.com/tislib/apibrew/pkg/util"
)

type UserGrpcService interface {
	stub.UserServer
}

type userServer struct {
	stub.UserServer
	service abs.UserService
}

func (u *userServer) Create(ctx context.Context, request *stub.CreateUserRequest) (*stub.CreateUserResponse, error) {
	users, err := u.service.Create(ctx, util.ArrayPrepend(request.Users, request.User))

	return &stub.CreateUserResponse{
		User:  util.ArrayFirst(users),
		Users: users,
	}, util.ToStatusError(err)
}

func (u *userServer) Update(ctx context.Context, request *stub.UpdateUserRequest) (*stub.UpdateUserResponse, error) {
	users, err := u.service.Update(ctx, util.ArrayPrepend(request.Users, request.User))

	return &stub.UpdateUserResponse{
		User:  util.ArrayFirst(users),
		Users: users,
	}, util.ToStatusError(err)
}

func (u *userServer) Delete(ctx context.Context, request *stub.DeleteUserRequest) (*stub.DeleteUserResponse, error) {
	err := u.service.Delete(ctx, request.Ids)

	return &stub.DeleteUserResponse{}, util.ToStatusError(err)
}

func (u *userServer) Get(ctx context.Context, request *stub.GetUserRequest) (*stub.GetUserResponse, error) {
	user, err := u.service.Get(ctx, request.Id)

	return &stub.GetUserResponse{
		User: user,
	}, util.ToStatusError(err)
}

func (u *userServer) List(ctx context.Context, request *stub.ListUserRequest) (*stub.ListUserResponse, error) {
	users, err := u.service.List(ctx, nil, request.Limit, request.Offset)

	return &stub.ListUserResponse{
		Content: users,
	}, util.ToStatusError(err)
}

func NewUserServer(service abs.UserService) stub.UserServer {
	return &userServer{service: service}
}
