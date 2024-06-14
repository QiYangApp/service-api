// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package auth

import (
	"errors"
	"fmt"
	"frame/errs"
)

var ErrAuthSourceNotActivated = errors.New("auth source is not activated")
var ErrUnsupportedLoginType = errors.New("Login source is unknown")

// ErrUserAlreadyExist represents a "user already exists" error.
type ErrUserAlreadyExist struct {
	Name string
}

// IsErrUserAlreadyExist checks if an error is a ErrUserAlreadyExists.
func IsErrUserAlreadyExist(err error) bool {
	var errUserAlreadyExist ErrUserAlreadyExist
	ok := errors.As(err, &errUserAlreadyExist)
	return ok
}

func (err ErrUserAlreadyExist) Error() string {
	return fmt.Sprintf("user already exists [name: %s]", err.Name)
}

// Unwrap unwraps this error as a ErrExist error
func (err ErrUserAlreadyExist) Unwrap() error {
	return errs.ErrAlreadyExist
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
	return errs.ErrNotExist
}
