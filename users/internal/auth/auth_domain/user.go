package auth_domain

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	uuid "github.com/satori/go.uuid"
)

type (
	User struct {
		Id                 uuid.UUID `json:"id"`
		Email              string    `json:"email"`
		Password           string    `json:"-"`
		FirstName          string    `json:"firstName"`
		LastName           string    `json:"lastName"`
		AuthorizationToken *string   `json:"authorizationToken,omitempty"`
		SessionToken       *string   `json:"-"`
	}
)

func NewUser(email, password, firstName, lastName string) *User {
	return &User{
		Id: uuid.NewV4(),
		Email: email,
		Password: password,
		FirstName: firstName,
		LastName: lastName,
	}
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email),
	)
}

func (u *User) SetAuthorizationToken(token string) {
	u.AuthorizationToken = &token
}
