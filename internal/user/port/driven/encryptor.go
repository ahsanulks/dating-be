package driven

type Encyptor interface {
	Encrypt(data []byte, cost int) ([]byte, error)
}
