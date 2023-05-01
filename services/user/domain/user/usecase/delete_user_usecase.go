package usecase

import (
	"context"
	prisma "prisma-go/prisma/db"
	userpb "prisma-go/proto/proto/user"
)

type deleteUserUseCase interface {
	DeleteUser(context.Context, *userpb.DeleteUserResponse) (*prisma.UserModel, error)
}

func (u *userUseCaseImpl) DeleteUser(ctx context.Context, req *userpb.DeleteUserResponse) (*prisma.UserModel, error) {
	res, err := u.db.User.
		FindUnique(prisma.User.ID.Equals(req.GetId())).
		Delete().Exec(ctx)

	return res, err
}
