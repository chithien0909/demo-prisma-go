package handler

import (
	"context"
	userpb "prisma-go/proto/proto/user"
	"prisma-go/services/user/domain/user/usecase"
)

type UserHandler interface {
	CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (user *userpb.User, err error)
	UpdateUser(context.Context, *userpb.UpdateUserRequest) (*userpb.User, error)
	DeleteUser(context.Context, *userpb.UserReq) (*userpb.DeleteUserResponse, error)
	FindUser(context.Context, *userpb.UserReq) (*userpb.User, error)
	FindUsers(context.Context, *userpb.FindUsersRequest) (*userpb.FindUsersResponse, error)
	FindManyUserCount(context.Context, *userpb.FindUsersRequest) (*userpb.FindManyUserCountResponse, error)
}

type userHandlerImpl struct {
	UserUseCase usecase.UserUseCase
}

func (u userHandlerImpl) UpdateUser(ctx context.Context, request *userpb.UpdateUserRequest) (*userpb.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userHandlerImpl) DeleteUser(ctx context.Context, req *userpb.UserReq) (*userpb.DeleteUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u userHandlerImpl) FindUser(ctx context.Context, req *userpb.UserReq) (*userpb.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userHandlerImpl) FindUsers(ctx context.Context, request *userpb.FindUsersRequest) (*userpb.FindUsersResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u userHandlerImpl) FindManyUserCount(ctx context.Context, request *userpb.FindUsersRequest) (*userpb.FindManyUserCountResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserHandler(
	userUseCase usecase.UserUseCase) UserHandler {
	return &userHandlerImpl{
		UserUseCase: userUseCase,
	}
}

func (u userHandlerImpl) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (user *userpb.User, err error) {

	return nil, nil
}
