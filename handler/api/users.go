package api

import (
	v1 "app/api/v1"
	"app/internal/user/param/request"
	"app/internal/user/port/driver"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type UserApiHandler struct {
	v1.UnimplementedUserServer

	userWriter driver.UserWriterUsecase
	log        log.Logger
}

func NewUserApiHandler(writer driver.UserWriterUsecase, log log.Logger) *UserApiHandler {
	return &UserApiHandler{
		userWriter: writer,
		log:        log,
	}
}

func (h UserApiHandler) CreateUser(ctx context.Context, params *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	userID, err := h.userWriter.CreateUser(ctx, &request.CreateUser{
		Username:    params.Username,
		PhoneNumber: params.PhoneNumber,
		Name:        params.Name,
		Password:    params.Password,
		Gender:      params.Gender,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &v1.CreateUserResponse{
		Id: userID,
	}, nil
}
