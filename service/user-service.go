package service

import (
	"data-handler/stub"
)

type UserService interface {
	stub.UserServiceServer
}

type userService struct {
	stub.UserServiceServer
}

func NewUserService() UserService {
	return &userService{}
}
