package usecase

import (
	"context"
	"prisma-go/prisma/generated/prisma-client"
)

type findUserUseCase interface {
	FindUsers(context.Context, *prisma.UsersParams) ([]prisma.User, error)
}

func (u *userUseCaseImpl) FindUsers(ctx context.Context, req *prisma.UsersParams) ([]prisma.User, error) {
	res, err := u.prisma.Users(req).Exec(ctx)

	return res, err
}
