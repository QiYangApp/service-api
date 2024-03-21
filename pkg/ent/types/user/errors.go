package user

import (
	util "ent/utils"
	"errors"
	"fmt"
)

// ErrAccountNotExist user account not exist
type ErrAccountNotExist struct {
	Account string
}

// IsErrEmailAddressNotExist checks if an error is an ErrEmailAddressNotExist
func IsErrEmailAddressNotExist(err error) bool {
	var errAccountNotExist ErrAccountNotExist
	var ok = errors.As(err, &errAccountNotExist)
	return ok
}

func (err ErrAccountNotExist) Error() string {
	return fmt.Sprintf("Email address does not exist [account: %s]", err.Account)
}

func (err ErrAccountNotExist) Unwrap() error {
	return util.ErrNotExist
}

// ErrUserProhibitLogin represents a "ErrUserProhibitLogin" kind of error.
type ErrUserProhibitLogin struct {
	UserId int64
	Name   string
}

// IsErrUserProhibitLogin checks if an error is a ErrUserProhibitLogin
func IsErrUserProhibitLogin(err error) bool {
	var errUserProhibitLogin ErrUserProhibitLogin
	ok := errors.As(err, &errUserProhibitLogin)
	return ok
}

func (err ErrUserProhibitLogin) Error() string {
	return fmt.Sprintf("user is not allowed login [userId: %d, name: %s]", err.UserId, err.Name)
}

// Unwrap unwraps this error as a ErrPermission error
func (err ErrUserProhibitLogin) Unwrap() error {
	return util.ErrPermissionDenied
}
