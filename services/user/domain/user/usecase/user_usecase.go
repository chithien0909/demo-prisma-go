package usecase

import (
	"prisma-go/prisma/generated/prisma-client"
)

type UserUseCase interface {
	findUserUseCase
}

type userUseCaseImpl struct {
	prisma *prisma.Client

	//validate *validator.Validate
}

func NewUserUseCase(
	prisma *prisma.Client,
	// requestValidate *validator.Validate,
) UserUseCase {
	return &userUseCaseImpl{
		prisma: prisma,
		//validate: requestValidate,
	}
}
