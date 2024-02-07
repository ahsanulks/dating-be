package infra

import (
	"context"

	"app/internal/usecase"

	"github.com/go-kratos/kratos/v2/log"
)

type GreeterRepo struct {
	db  *PostgresDB
	log *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(db *PostgresDB, logger log.Logger) *GreeterRepo {
	return &GreeterRepo{
		db:  db,
		log: log.NewHelper(logger),
	}
}

func (r *GreeterRepo) Save(ctx context.Context, g *usecase.Greeter) (*usecase.Greeter, error) {
	return g, nil
}

func (r *GreeterRepo) Update(ctx context.Context, g *usecase.Greeter) (*usecase.Greeter, error) {
	return g, nil
}

func (r *GreeterRepo) FindByID(context.Context, int64) (*usecase.Greeter, error) {
	return nil, nil
}

func (r *GreeterRepo) ListByHello(context.Context, string) ([]*usecase.Greeter, error) {
	return nil, nil
}

func (r *GreeterRepo) ListAll(context.Context) ([]*usecase.Greeter, error) {
	return nil, nil
}
