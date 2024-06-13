package user

import (
	"errors"
	"fmt"
	util "frame/errors"
)

// ErrAccountNotExist user account not exist
type ErrAccountNotExist struct {
	Account string
}

// IsErrAccountNotExist checks if an error is an ErrEmailAddressNotExist
func IsErrAccountNotExist(err error) bool {
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

// ErrUserNotExist represents a "UserNotExist" kind of error.
type ErrUserNotExist struct {
	UID  int64
	Name string
}

// IsErrUserNotExist checks if an error is a ErrUserNotExist.
func IsErrUserNotExist(err error) bool {
	var errUserNotExist ErrUserNotExist
	ok := errors.As(err, &errUserNotExist)
	return ok
}

func (err ErrUserNotExist) Error() string {
	return fmt.Sprintf("user does not exist [uid: %d, name: %s]", err.UID, err.Name)
}

// Unwrap unwraps this error as a ErrNotExist error
func (err ErrUserNotExist) Unwrap() error {
	return util.ErrNotExist
}

// ErrUserInactive represents a "ErrUserInactive" kind of error.
type ErrUserInactive struct {
	UID  int64
	Name string
}

// IsErrUserInactive checks if an error is a ErrUserInactive
func IsErrUserInactive(err error) bool {
	_, ok := err.(ErrUserInactive)
	return ok
}

func (err ErrUserInactive) Error() string {
	return fmt.Sprintf("user is inactive [uid: %d, name: %s]", err.UID, err.Name)
}

// Unwrap unwraps this error as a ErrPermission error
func (err ErrUserInactive) Unwrap() error {
	return util.ErrPermissionDenied
}
