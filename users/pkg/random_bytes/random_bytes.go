package random_bytes

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/labstack/echo/v4"
)

type Service struct {
	log echo.Logger
}

func New(l echo.Logger) *Service {
	return &Service{
		log: l,
	}
}

func (r *Service) Bytes(size int) ([]byte, error) {
	val, err := r.make(size)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (r *Service) Base64(size int) (string, error) {
	val, err := r.make(size)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(val), nil
}

func (r *Service) URLBase64(size int) (string, error) {
	val, err := r.make(size)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(val), nil
}

func (r *Service) make(size int) ([]byte, error) {
	val := make([]byte, size)
	_, err := rand.Read(val)
	if err != nil {
		return nil, fmt.Errorf("failed to create random bytes: %w", err)
	}

	return val, nil
}
