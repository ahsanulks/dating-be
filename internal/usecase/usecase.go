package usecase

import "github.com/google/wire"

// ProviderSet is usecase providers.
var ProviderSet = wire.NewSet(NewGreeterUsecase)
