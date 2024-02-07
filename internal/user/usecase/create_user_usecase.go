package usecase

import (
	"app/internal/user/entity"
	"app/internal/user/param/request"
	"context"
)

const (
	costEncryption = 10
)

func (uu UserUsecase) CreateUser(ctx context.Context, params *request.CreateUser) (id int64, err error) {
	user, err := entity.NewUser(params)
	if err != nil {
		return id, err
	}

	encryptedPassword, err := uu.encryptor.Encrypt([]byte(user.Password), costEncryption)
	if err != nil {
		return id, err
	}
	user.Password = string(encryptedPassword)
	return uu.userWriter.Create(ctx, user)
}
