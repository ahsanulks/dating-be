package fake

import (
	"app/internal/user/param/request"
	"app/internal/user/param/response"
	"app/internal/user/port/driver"
	"context"
	"errors"
	"math/rand"

	"github.com/go-faker/faker/v4"
)

var (
	_ driver.UserWriterUsecase = new(FakeUserUsecase)
)

type FakeUserUsecase struct{}

// CreateUser implements driver.UserWriterUsecase.
func (*FakeUserUsecase) CreateUser(ctx context.Context, params *request.CreateUser) (id int64, err error) {
	if params.Username == "test123" {
		return 0, errors.New("cannot create user")
	}
	return faker.NewSafeSource(rand.NewSource(1000)).Int63(), nil
}

// GenerateUserToken implements driver.UserWriterUsecase.
func (*FakeUserUsecase) GenerateUserToken(ctx context.Context, params *request.GenerateUserToken) (*response.Token, error) {
	if params.Username == "test123" {
		return nil, errors.New("cannot create user")
	}
	return &response.Token{
		Token:     faker.Jwt(),
		ExpiresIn: 3600,
		Type:      "Bearer",
	}, nil
}
