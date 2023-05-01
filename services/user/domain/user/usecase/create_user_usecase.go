package usecase

import (
	"context"
	"prisma-go/internal/security"
	prisma "prisma-go/prisma/db"
	userpb "prisma-go/proto/proto/user"
	"time"
)

type createUserUseCase interface {
	CreateUser(context.Context, *userpb.CreateUserRequest) (*prisma.UserModel, error)
}

func (u *userUseCaseImpl) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*prisma.UserModel, error) {
	password, _ := security.HashAndSalt([]byte(req.Password))
	role := prisma.RoleUSER

	res, err := u.db.User.CreateOne(
		prisma.User.Name.Set(req.Name),
		prisma.User.Email.Set(req.Name),
		prisma.User.Password.Set(password),
		prisma.User.UpdatedAt.Set(time.Now()),
		prisma.User.Role.Set(role),
	).Exec(ctx)

	return res, err
}
