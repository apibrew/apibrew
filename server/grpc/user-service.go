package grpc

import (
	"context"
	"data-handler/server/stub"
	util2 "data-handler/server/util"
	"data-handler/service"
	"data-handler/util"
)

type UserGrpcService interface {
	stub.UserServiceServer
}

type userServiceServer struct {
	stub.UserServiceServer
	service service.UserService
}

func (u *userServiceServer) Create(ctx context.Context, request *stub.CreateUserRequest) (*stub.CreateUserResponse, error) {
	users, err := u.service.Create(ctx, util.ArrayPrepend(request.Users, request.User))

	return &stub.CreateUserResponse{
		User:  util.ArrayFirst(users),
		Users: users,
	}, util2.ToStatusError(err)
}

func (u *userServiceServer) Update(ctx context.Context, request *stub.UpdateUserRequest) (*stub.UpdateUserResponse, error) {
	users, err := u.service.Update(ctx, util.ArrayPrepend(request.Users, request.User))

	return &stub.UpdateUserResponse{
		User:  util.ArrayFirst(users),
		Users: users,
	}, util2.ToStatusError(err)
}

func (u *userServiceServer) Delete(ctx context.Context, request *stub.DeleteUserRequest) (*stub.DeleteUserResponse, error) {
	err := u.service.Delete(ctx, request.Ids)

	return &stub.DeleteUserResponse{}, util2.ToStatusError(err)
}

func (u *userServiceServer) Get(ctx context.Context, request *stub.GetUserRequest) (*stub.GetUserResponse, error) {
	user, err := u.service.Get(ctx, request.Id)

	return &stub.GetUserResponse{
		User: user,
	}, util2.ToStatusError(err)
}

func (u *userServiceServer) List(ctx context.Context, request *stub.ListUserRequest) (*stub.ListUserResponse, error) {
	users, err := u.service.List(ctx, nil, request.Limit, request.Offset)

	return &stub.ListUserResponse{
		Content: users,
	}, util2.ToStatusError(err)
}

func NewUserServiceServer(service service.UserService) stub.UserServiceServer {
	return &userServiceServer{service: service}
}
