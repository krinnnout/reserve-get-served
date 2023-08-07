package types

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type User struct {
	Id                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"encrypted_password" json:"-"`
}

const (
	bcryptCost      = 12
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPassLen      = 7
)

type UserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (params UserParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.FirstName) < minFirstNameLen {
		errors["firstName"] = fmt.Sprintf("first name cannot be less than %d characters", minFirstNameLen)
	}
	if len(params.LastName) < minLastNameLen {
		errors["lastName"] = fmt.Sprintf("last name cannot be less than %d characters", minLastNameLen)
	}
	if len(params.Password) < minPassLen {
		errors["password"] = fmt.Sprintf("password length cannot be less than %d characters", minPassLen)
	}

	if !isEmailValid(params.Email) {
		errors["email"] = fmt.Sprintf("email not valid")
	}
	return errors
}

func isEmailValid(e string) bool {
	return regexp.MustCompile(`[a-z0-9]+@[a-z]+\.[a-z]{2,3}`).MatchString(e)
}

func NewUserFromParams(params UserParams) (*User, error) {
	encrPsw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encrPsw),
	}, nil
}
