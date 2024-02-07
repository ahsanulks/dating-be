package fake

import (
	"app/internal/user/entity"
	"app/internal/user/param/response"
	"app/internal/user/port/driven"
	"errors"
)

var _ driven.TokenProvider[*entity.User] = new(FakeTokenProvider)

type FakeTokenProvider struct{}

// Generate implements driven.TokenProvider.
func (*FakeTokenProvider) Generate(user *entity.User) (*response.Token, error) {
	if user.Username == "wrongUsername" {
		return nil, errors.New("invalid")
	}
	return &response.Token{
		Token:     "1231313213213131",
		ExpiresIn: 3600,
		Type:      "Bearer",
	}, nil
}

func (*FakeTokenProvider) ValidateToken(tokenString string) (map[string]interface{}, error) {
	return nil, nil
}
