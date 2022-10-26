package grpc_service

import (
	"context"
	"data-handler/grpc/stub"
	"data-handler/service"
)

type UserGrpcService interface {
	stub.UserServiceServer
}

type userServiceServer struct {
	stub.UserServiceServer
	service service.UserService
}

func (u *userServiceServer) Create(ctx context.Context, request *stub.CreateUserRequest) (*stub.CreateUserResponse, error) {
	users, err := u.service.Create(ctx, request.Users)

	return &stub.CreateUserResponse{
		Users: users,
		Error: toProtoError(err),
	}, nil
}

func (u *userServiceServer) Update(ctx context.Context, request *stub.UpdateUserRequest) (*stub.UpdateUserResponse, error) {
	users, err := u.service.Update(ctx, request.Users)

	return &stub.UpdateUserResponse{
		Users: users,
		Error: toProtoError(err),
	}, err
}

func (u *userServiceServer) Delete(ctx context.Context, request *stub.DeleteUserRequest) (*stub.DeleteUserResponse, error) {
	err := u.service.Delete(ctx, request.Ids)

	return &stub.DeleteUserResponse{
		Error: toProtoError(err),
	}, nil
}

func (u *userServiceServer) Get(ctx context.Context, request *stub.GetUserRequest) (*stub.GetUserResponse, error) {
	user, err := u.service.Get(ctx, request.Id)

	return &stub.GetUserResponse{
		User:  user,
		Error: toProtoError(err),
	}, nil
}

func (u *userServiceServer) List(ctx context.Context, request *stub.ListUserRequest) (*stub.ListUserResponse, error) {
	users, err := u.service.List(ctx, request.Query, request.Limit, request.Offset)

	return &stub.ListUserResponse{
		Content: users,
		Error:   toProtoError(err),
	}, err
}

func NewUserServiceServer(service service.UserService) stub.UserServiceServer {
	return &userServiceServer{service: service}
}
