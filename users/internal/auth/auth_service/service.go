package auth_service

import (
	"github.com/divilla/tproto/users/internal/auth/auth_domain"
	"github.com/divilla/tproto/users/internal/auth/auth_repository"
	"github.com/divilla/tproto/users/internal/containers"
	"github.com/labstack/echo/v4"
)

type (
	Service struct {
		log  echo.Logger
		rbs  randomBytesService
		repo repository
	}

	randomBytesService interface {
		URLBase64(size int) (string, error)
	}

	repository interface {
		Login(email, password string) (*auth_domain.User, error)
	}
)

func New(mc *containers.Main) *Service {
	return &Service{
		log: mc.Log(),
		rbs: mc.RandomBytesService(),
		repo: auth_repository.New(mc),
	}
}

func (s *Service) Login(u *LoginDTO) (*auth_domain.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	user, err := s.repo.Login(u.Email, u.Password)
	if err != nil {
		return nil, err
	}

	token, err := s.rbs.URLBase64(64)
	if err != nil {
		return nil, err
	}

	user.SetAuthorizationToken(token)
	return user, nil
}
