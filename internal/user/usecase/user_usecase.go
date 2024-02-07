package usecase

import (
	"app/internal/user/port/driven"
)

type UserUsecase struct {
	userWriter driven.UserWriter
	encryptor  driven.Encyptor
}

func NewUserUsecase(
	userWriter driven.UserWriter,
	encryptor driven.Encyptor,
) *UserUsecase {
	return &UserUsecase{
		userWriter: userWriter,
		encryptor:  encryptor,
	}
}
