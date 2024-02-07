package encryption

import (
	"app/internal/user/port/driven"

	"golang.org/x/crypto/bcrypt"
)

var _ driven.Encyptor = new(BcyrpEncryption)

type BcyrpEncryption struct{}

func (*BcyrpEncryption) Encrypt(data []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(data, cost)
}

func (*BcyrpEncryption) CompareEncryptedAndData(encrypted, data []byte) error {
	return bcrypt.CompareHashAndPassword(encrypted, data)
}
