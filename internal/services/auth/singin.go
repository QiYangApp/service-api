// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2019 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package auth

import (
	"context"
	"ent/models"
	authtype "ent/types/auth"
	usertype "ent/types/user"
	"service-api/internal/repo"
	usermodel "service-api/internal/repo/user"
	"strings"
)

func UserSingIn(ctx context.Context, username, passwd string) (*models.User, *models.Source, error) {
	var user *models.User
	var err error

	trimmedUsername := strings.TrimSpace(username)

	if trimmedUsername == "" {
		return nil, nil, authtype.ErrUserNotExist{Name: username}
	}

	userSourceType := authtype.NoType
	if strings.Contains(trimmedUsername, "@") {
		userSourceType = authtype.SMTP
		trimmedUsername = strings.ToLower(trimmedUsername)
	}

	var account *models.Accounts
	if account, err = usermodel.GetSingleAccountByName(ctx, trimmedUsername); err != nil {
		return nil, nil, err
	}

	// verify user account is exist and is active
	if account == nil || !account.IsActivated {
		return nil, nil, usertype.ErrAccountNotExist{
			Account: username,
		}
	}

	if user, err = repo.Client.User.Get(ctx, account.UserID); err != nil {
		return nil, nil, err
	}

	if user != nil {
		var sourceType *models.Source
		if sourceType, err = repo.Client.Source.Get(ctx, user.LoginSource); err != nil {
			return nil, nil, err
		}

		if sourceType == nil || !sourceType.IsActive {
			return nil, nil, authtype.ErrAuthSourceNotActivated
		}

		var authenticator PasswordAuthenticator
		var ok bool
		if authenticator, ok = sourceType.Cfg.Cfg.(PasswordAuthenticator); !ok {
			return nil, nil, authtype.ErrUnsupportedLoginType
		}

		if user, err := authenticator.Authenticate(ctx, user, user.LoginName, passwd); err != nil {
			return nil, nil, err
		}

		if user.ProhibitLogin {
			return nil, nil, usertype.ErrUserProhibitLogin{UserId: user.ID, Name: user.Name}
		}

		return user, sourceType, nil
	}

	if userSourceType

	return nil, nil, nil
}

func userSingInAllSource(ctx context.Context, username, passwd string) (*models.User, *models.Source, error) {

}