package usecase

import (
	"context"
	prisma "prisma-go/prisma/db"
	userpb "prisma-go/proto/proto/user"
	"prisma-go/utils"
)

type findUserUseCase interface {
	FindUsers(ctx context.Context, req *userpb.FindUsersRequest) ([]prisma.UserModel, error)
	FindUser(context.Context, *userpb.UserReq) (*prisma.UserModel, error)
}

func (u *userUseCaseImpl) FindUsers(ctx context.Context, req *userpb.FindUsersRequest) ([]prisma.UserModel, error) {

	var params []prisma.UserWhereParam

	if req.GetQ() != "" {
		params = append(params, prisma.User.Or(
			prisma.User.Email.Contains(req.GetQ()),
			prisma.User.Name.Contains(req.GetQ()),
		))
	}

	if req.GetName() != "" {
		params = append(params, prisma.User.Name.Equals(req.GetName()))
	}

	if req.GetEmail() != "" {
		params = append(params, prisma.User.Email.Equals(req.GetEmail()))
	}

	skip, take := utils.BuildPagination(int(req.GetPage()), int(req.GetLimit()))

	res, err := u.db.User.
		FindMany(params...).
		Skip(skip).
		Take(take).
		Exec(ctx)

	return res, err
}

func (u *userUseCaseImpl) FindUser(ctx context.Context, req *userpb.UserReq) (*prisma.UserModel, error) {
	if req.GetId() == "" {
		return nil, prisma.ErrNotFound
	}
	user, err := u.db.User.FindUnique(
		prisma.User.ID.Equals(req.GetId()),
	).Exec(ctx)

	return user, err
}
