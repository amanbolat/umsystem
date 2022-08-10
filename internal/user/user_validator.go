package user

import (
	"errors"
	"fmt"
)

var ErrInvalidUsername = errors.New("invalid username")
var ErrInvalidPassword = errors.New("invalid password")

type UserValidator struct {
	minPasswordLength uint
	minUsernameLength uint
}

func NewUserValidator(minPasswordLength, minUsernameLength uint) *UserValidator {
	return &UserValidator{
		minPasswordLength: minPasswordLength,
		minUsernameLength: minUsernameLength,
	}
}

func (v UserValidator) ValidateUser(usr User) error {
	if len(usr.username) < int(v.minUsernameLength) {
		return fmt.Errorf("%w: username should be longer than %d", ErrInvalidUsername, v.minUsernameLength)
	}

	if len(usr.password) < int(v.minPasswordLength) {
		return fmt.Errorf("%w: password should be longer than %d", ErrInvalidPassword, v.minPasswordLength)
	}

	return nil
}
