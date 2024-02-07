package driver

import (
	"app/internal/user/param/request"
	"app/internal/user/param/response"
	"context"
)

type UserWriterUsecase interface {
	CreateUser(ctx context.Context, params *request.CreateUser) (id int64, err error)
	GenerateUserToken(ctx context.Context, params *request.GenerateUserToken) (*response.Token, error)
}
