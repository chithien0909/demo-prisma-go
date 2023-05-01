package usecase

import (
	"context"
	"prisma-go/internal/security"
	prisma "prisma-go/prisma/db"
	userpb "prisma-go/proto/proto/user"
	"time"
)

type updateUserUseCase interface {
	UpdateUser(context.Context, *userpb.UpdateUserRequest) (*prisma.UserModel, error)
}

func (u *userUseCaseImpl) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*prisma.UserModel, error) {

	name := req.GetName()
	email := req.GetEmail()
	password := req.GetPassword()

	if req.GetPassword() != "" {
		password, _ = security.HashAndSalt([]byte(req.GetPassword()))
	}

	res, err := u.db.User.UpsertOne(
		prisma.User.ID.Equals(req.GetId()),
	).Update(
		prisma.User.Name.SetIfPresent(&name),
		prisma.User.Email.SetIfPresent(&email),
		prisma.User.Password.SetIfPresent(&password),
		prisma.User.UpdatedAt.Set(time.Now()),
	).Exec(ctx)

	return res, err
}
