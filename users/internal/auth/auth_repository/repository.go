package auth_repository

import (
	"github.com/divilla/tproto/users/internal/auth/auth_domain"
	"github.com/divilla/tproto/users/internal/containers"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

var (
	users = []*auth_domain.User{
		{
			Id:        uuid.NewV4(),
			Email:     "vitodivilla@gmail.com",
			Password:  "Pass11",
			FirstName: "Vito",
			LastName:  "Divilla",
		},
		{
			Id:        uuid.NewV4(),
			Email:     "vito.dolfi@gmail.com",
			Password:  "Pass11",
			FirstName: "Vito",
			LastName:  "Dolfi",
		},
	}

	ErrInvalidUsernamePassword = auth_domain.NewHttpError(http.StatusUnprocessableEntity, "invalid email and/or password")
)

type (
	AuthRepository struct {
		log echo.Logger
	}

	RegisterDTO struct {
		Email     string
		Password  string
		FirstName string
		LastName  string
	}
)

func New(mc *containers.Main) *AuthRepository {
	return &AuthRepository{
		log: mc.Log(),
	}
}

func (r AuthRepository) Login(email, password string) (*auth_domain.User, error) {
	for _, user := range users {
		if user.Email == email && user.Password == password {
			return user, nil
		}
	}

	return nil, ErrInvalidUsernamePassword
}
