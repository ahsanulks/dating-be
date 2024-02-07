package usecase

import (
	customerror "app/internal/custom_error"
	"app/internal/user/param/request"
	"app/internal/user/param/response"
	"context"
)

func (uu UserWriterUsecase) GenerateUserToken(ctx context.Context, params *request.GenerateUserToken) (*response.Token, error) {
	user, err := uu.userGetter.GetByUsername(ctx, params.Username)
	if err != nil {
		return nil, customerror.NewValidationErrorWithMessage("authentication", "wrong username/password")
	}

	err = uu.encryptor.CompareEncryptedAndData([]byte(user.Password), []byte(params.Password))
	if err != nil {
		return nil, customerror.NewValidationErrorWithMessage("authentication", "wrong username/password")
	}

	token, err := uu.tokenProvider.Generate(user)
	if err != nil {
		return nil, err
	}

	err = uu.userWriter.UpdateLoginInformation(ctx, user)
	if err != nil {
		return nil, err
	}
	return token, nil
}
