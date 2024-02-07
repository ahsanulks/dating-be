package infra

import (
	"app/infra/database"
	"app/infra/encryption"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(database.NewPostgresDB, new(encryption.BcyrpEncryption))
