package handler

import (
	"app/handler/api"

	"github.com/google/wire"
)

// ProviderSet is handler providers.
var ProviderSet = wire.NewSet(api.NewUserApiHandler)
