package encryption

import (
	"app/internal/user/port/driven"

	"golang.org/x/crypto/bcrypt"
)

var _ driven.Encyptor = new(BcryptEncryption)

type BcryptEncryption struct{}

func NewBcryptEncryption() *BcryptEncryption {
	return &BcryptEncryption{}
}

func (*BcryptEncryption) Encrypt(data []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(data, cost)
}

func (*BcryptEncryption) CompareEncryptedAndData(encrypted, data []byte) error {
	return bcrypt.CompareHashAndPassword(encrypted, data)
}
