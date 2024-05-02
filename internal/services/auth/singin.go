// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2019 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package auth

import (
	"context"
	"ent/models"
	authtype "ent/types/auth"
	usertype "ent/types/user"
	"frame/modules/log"
	"service-api/internal/repo"
	authmodel "service-api/internal/repo/auth"
	usermodel "service-api/internal/repo/user"
	"strings"
)

func UserSignIn(ctx context.Context, username, passwd string) (*models.User, *models.Source, error) {
	var user *models.User
	var sourceType *models.Source
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
	if account == nil {
		return nil, nil, usertype.ErrAccountNotExist{
			Account: username,
		}
	}

	if !account.IsActivated {
		return nil, nil, authtype.ErrAuthSourceNotActivated
	}

	if user, err = repo.Client.User.Get(ctx, account.UserID); err != nil {
		return nil, nil, err
	}

	if user != nil {
		if sourceType, err = repo.Client.Source.Get(ctx, user.LoginSource); err != nil {
			return nil, nil, err
		}

		if sourceType == nil || !sourceType.IsActive {
			return nil, nil, authtype.ErrAuthSourceNotActivated
		}

		var authenticator PasswordAuthenticator
		var ok bool
		if authenticator, ok = sourceType.Cfg.Value().(PasswordAuthenticator); !ok {
			return nil, nil, authtype.ErrUnsupportedLoginType
		}

		if user, err = authenticator.Authenticate(ctx, user, user.LoginName, passwd); err != nil {
			return nil, nil, err
		}

		if user.ProhibitLogin {
			return nil, nil, usertype.ErrUserProhibitLogin{UserId: user.ID, Name: user.Name}
		}

		return user, sourceType, nil
	}

	user, sourceType, err = UserSingInAllSource(ctx, username, passwd)
	if err == nil {
		return user, sourceType, nil
	}

	if userSourceType == authtype.SMTP {
		return nil, nil, usertype.ErrAccountNotExist{Account: username}
	}

	return nil, nil, err
}

func UserSingInAllSource(ctx context.Context, username, passwd string) (*models.User, *models.Source, error) {
	var sources []*models.Source
	var err error
	if sources, err = authmodel.GetAllSourceByIsActive(ctx, true); err != nil {
		return nil, nil, err
	}

	for _, source := range sources {

		var authenticator PasswordAuthenticator
		var ok bool
		if authenticator, ok = source.Cfg.Value().(PasswordAuthenticator); !ok {
			continue
		}

		var user *models.User
		if user, err = authenticator.Authenticate(ctx, nil, username, passwd); err != nil || user == nil {
			continue
		}

		if !user.ProhibitLogin {
			return user, source, nil
		}

		err = usertype.ErrUserProhibitLogin{UserId: user.ID, Name: user.Name}
		if usertype.IsErrUserNotExist(err) {
			log.Sugar().Debugf("Failed to login '%s' via '%s': %v", username, source.Name, err)
		} else {
			log.Sugar().Warnf("Failed to login '%s' via '%s': %v", username, source.Name, err)
		}
	}

	return nil, nil, usertype.ErrUserNotExist{Name: username}
}
