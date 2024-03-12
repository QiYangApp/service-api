package accounts

import (
	"errors"
	"fmt"
	"framework/exceptions"
)

// ErrAccountAlreadyUsed represents a "EmailAlreadyUsed" kind of error.
type ErrAccountAlreadyUsed struct {
	Account string
	Type    Type
}

// IsErrEmailAlreadyUsed checks if an error is a ErrEmailAlreadyUsed.
func IsErrEmailAlreadyUsed(err error) bool {
	var errAccountAlreadyUsed ErrAccountAlreadyUsed
	ok := errors.As(err, &errAccountAlreadyUsed)
	return ok
}

func (err ErrAccountAlreadyUsed) Error() string {
	return fmt.Sprintf("account already in use [account: %s]", err.Account)
}

func (err ErrAccountAlreadyUsed) Unwrap() error {
	return exceptions.ErrAlreadyExist
}

// ErrUserNotExist represents a "UserNotExist" kind of error.
type ErrUserNotExist struct {
	UID  int64
	Name string
}

// IsErrUserNotExist checks if an error is a ErrUserNotExist.
func IsErrUserNotExist(err error) bool {
	_, ok := err.(ErrUserNotExist)
	return ok
}

func (err ErrUserNotExist) Error() string {
	return fmt.Sprintf("user does not exist [uid: %d, name: %s]", err.UID, err.Name)
}

// Unwrap unwraps this error as a ErrNotExist error
func (err ErrUserNotExist) Unwrap() error {
	return exceptions.ErrNotExist
}
