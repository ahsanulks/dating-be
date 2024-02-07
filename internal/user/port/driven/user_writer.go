package driven

import (
	"app/internal/user/entity"
	"context"
)

type UserWriter interface {
	Create(ctx context.Context, user *entity.User) (id int64, err error)
}
