package driven

import "app/internal/user/param/response"

type TokenProvider[T any] interface {
	Generate(data T) (*response.Token, error)
}
