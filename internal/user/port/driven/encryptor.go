package driven

type Encyptor interface {
	Encrypt(data []byte, cost int) ([]byte, error)
	CompareEncryptedAndData(encrypted, data []byte) error
}
