package interfaces

type RandomBytesService interface {
	Bytes(size int) ([]byte, error)
	Base64(size int) (string, error)
	URLBase64(size int) (string, error)
}
