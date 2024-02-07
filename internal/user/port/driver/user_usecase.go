package driver

import (
	"app/internal/user/param/request"
	"context"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, params *request.CreateUser) (id int64, err error)
}
