package usecase

import (
	"app/internal/user/port/driven"
)

type UserWriterUsecase struct {
	userWriter driven.UserWriter
	encryptor  driven.Encyptor
}

func NewUserWriterUsecase(
	userWriter driven.UserWriter,
	encryptor driven.Encyptor,
) *UserWriterUsecase {
	return &UserWriterUsecase{
		userWriter: userWriter,
		encryptor:  encryptor,
	}
}
