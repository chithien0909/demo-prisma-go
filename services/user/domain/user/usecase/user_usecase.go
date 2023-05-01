package usecase

import (
	"prisma-go/prisma/db"
)

type UserUseCase interface {
	findUserUseCase
	createUserUseCase
	updateUserUseCase
	deleteUserUseCase
}

type userUseCaseImpl struct {
	db *db.PrismaClient

	//validate *validator.Validate
}

func NewUserUseCase(
	db *db.PrismaClient,
	// requestValidate *validator.Validate,
) UserUseCase {
	return &userUseCaseImpl{
		db: db,
		//validate: requestValidate,
	}
}
