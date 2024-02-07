package usecase

import (
	"app/internal/user/entity"
	"app/internal/user/port/driven"
)

type UserWriterUsecase struct {
	userWriter    driven.UserWriter
	encryptor     driven.Encyptor
	userGetter    driven.UserGetter
	tokenProvider driven.TokenProvider[*entity.User]
}

func NewUserWriterUsecase(
	userWriter driven.UserWriter,
	encryptor driven.Encyptor,
	userGetter driven.UserGetter,
	tokenProvider driven.TokenProvider[*entity.User],
) *UserWriterUsecase {
	return &UserWriterUsecase{
		userWriter:    userWriter,
		encryptor:     encryptor,
		userGetter:    userGetter,
		tokenProvider: tokenProvider,
	}
}
