package fake

import (
	"app/internal/user/entity"
	"app/internal/user/port/driven"
	"context"
	"errors"
	"math/rand"

	"github.com/go-faker/faker/v4"
)

var (
	_ driven.UserWriter = new(FakeUserDriven)
	_ driven.UserGetter = new(FakeUserDriven)
)

type ContextType string

type FakeUserDriven struct {
	data           map[int64]*entity.User
	dataByUsername map[string]*entity.User
}

func NewFakeUserDriven() *FakeUserDriven {
	return &FakeUserDriven{
		data:           make(map[int64]*entity.User),
		dataByUsername: make(map[string]*entity.User),
	}
}

func (fud *FakeUserDriven) Create(ctx context.Context, user *entity.User) (id int64, err error) {
	user.ID = faker.NewSafeSource(rand.NewSource(1000)).Int63()
	fud.data[user.ID] = user
	fud.dataByUsername[user.Username] = user
	return user.ID, nil
}

func (fud FakeUserDriven) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	if user, ok := fud.data[id]; ok {
		return user, nil
	}
	return nil, errors.New("resource not found")
}

func (*FakeUserDriven) UpdateLoginInformation(ctx context.Context, user *entity.User) error {
	if val := ctx.Value(ContextType("token_error")); val != nil {
		return errors.New("error")
	}
	return nil
}

// GetByUsername implements driven.UserGetter.
func (fud *FakeUserDriven) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	if user, ok := fud.dataByUsername[username]; ok {
		return user, nil
	}
	return nil, errors.New("resource not found")
}
