package driven

import "app/internal/user/param/response"

type TokenProvider[T any] interface {
	Generate(data T) (*response.Token, error)
	ValidateToken(tokenString string) (map[string]any, error)
}
