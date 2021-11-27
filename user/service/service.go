package service

import (
	"context"
	"ponica/user/gen/api"
)

type UserService struct{}

func NewUserService() UserService {
	return UserService{}
}

func (service *UserService) ShowUser(ctx context.Context, req *api.ShowUserRequest) (*api.ShowUserResponse, error) {
	return &api.ShowUserResponse{
		User: &api.User{
			Id:    1,
			Name:  "show user name",
			Email: "show user email",
		},
	}, nil
}

func (service *UserService) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	return &api.CreateUserResponse{
		Id: 2,
	}, nil
}
