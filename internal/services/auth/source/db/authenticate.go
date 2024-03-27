// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package db

import (
	"context"
	"ent/models"
	usertype "ent/types/user"
	util "ent/utils"
	"fmt"
	"service-api/internal/modules/setting"
	usermodel "service-api/internal/repo/user"
)

// ErrUserPasswordNotSet represents a "ErrUserPasswordNotSet" kind of error.
type ErrUserPasswordNotSet struct {
	UID  int64
	Name string
}

func (err ErrUserPasswordNotSet) Error() string {
	return fmt.Sprintf("user's password isn't set [uid: %d, name: %s]", err.UID, err.Name)
}

// Unwrap unwraps this error as a ErrInvalidArgument error
func (err ErrUserPasswordNotSet) Unwrap() error {
	return util.ErrInvalidArgument
}

// ErrUserPasswordInvalid represents a "ErrUserPasswordInvalid" kind of error.
type ErrUserPasswordInvalid struct {
	UID  int64
	Name string
}

func (err ErrUserPasswordInvalid) Error() string {
	return fmt.Sprintf("user's password is invalid [uid: %d, name: %s]", err.UID, err.Name)
}

// Unwrap unwraps this error as a ErrInvalidArgument error
func (err ErrUserPasswordInvalid) Unwrap() error {
	return util.ErrInvalidArgument
}

// Authenticate authenticates the provided user against the DB
func Authenticate(ctx context.Context, user *models.User, login, password string) (*models.User, error) {
	if user == nil {
		return nil, usertype.ErrUserNotExist{Name: login}
	}

	if !usermodel.IsPasswdSet(user) {
		return nil, ErrUserPasswordNotSet{UID: user.ID, Name: user.Name}
	} else if !usermodel.ValidatePassword(user, password) {
		return nil, ErrUserPasswordInvalid{UID: user.ID, Name: user.Name}
	}

	// Update password hash if server password hash algorithm have changed
	// Or update the password when the salt length doesn't match the current
	// recommended salt length, this in order to migrate user's salts to a more secure salt.
	if user.PasswdHashAlgo != setting.SecretSetting.PasswdHashAlgo || len(user.PasswdSalt) != usermodel.SaltByteLength*2 {
		if err := usermodel.SetPassword(user, password); err != nil {
			return nil, err
		}

		var err error
		if user, err = usermodel.UpdatePassword(ctx, user); err != nil {
			return nil, err
		}
	}

	// WARN: DON'T check user.IsActive, that will be checked on reqSign so that
	// user could be hinted to resend confirm email.
	if user.ProhibitLogin {
		return nil, usertype.ErrUserProhibitLogin{
			UserId: user.ID,
			Name:   user.Name,
		}
	}

	return user, nil
}
