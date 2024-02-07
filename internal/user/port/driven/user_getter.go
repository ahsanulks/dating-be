package driven

import (
	"app/internal/user/entity"
	"context"
)

type UserGetter interface {
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
}
