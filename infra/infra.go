package infra

import (
	"app/infra/database"
	"app/infra/encryption"
	tokenprovider "app/infra/token_provider"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	database.NewPostgresDB,
	encryption.NewBcryptEncryption,
	database.NewUserRepository,
	tokenprovider.NewUserJwtProvider,
)
