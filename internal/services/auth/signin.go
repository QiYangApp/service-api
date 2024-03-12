// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package auth

import (
	"context"
	"service-api/internal/ent"
	accountmodel "service-api/internal/models/accounts"
	usermodel "service-api/internal/models/user"

	"strings"
)

// UserSignIn validates user name and password.
func UserSignIn(ctx context.Context, username, password string) (*ent.User, *ent.Source, error) {

	var user *ent.User
	isEmail := false
	singinType := accountmodel.Account

	trimmedUsername := strings.TrimSpace(username)
	if len(trimmedUsername) == 0 {
		return nil, nil, accountmodel.ErrUserNotExist{
			Name: username,
		}
	}

	// check same email
	if strings.Contains(trimmedUsername, "@") {
		trimmedUsername = strings.ToLower(trimmedUsername)
		singinType = accountmodel.Email
	}

	// check same account
	userAccount, err := accountmodel.Find(trimmedUsername, singinType, ctx)
	if err != nil {
		return nil, nil, err
	}

	if !userAccount.IsActivated {
		return nil, nil, accountmodel.ErrUserNotExist{
			Name: username,
		}
	}

	user = &ent.User{ID: userAccount.ID}

	if user != nil {
		hasUser, err := usermodel.FindById(ctx, user.ID)
		if err != nil {
			return nil, nil, err
		}

		if hasUser {
			source, err := auth.GetSourceByID(ctx, user.LoginSource)
			if err != nil {
				return nil, nil, err
			}

			if !source.IsActive {
				return nil, nil, oauth2.ErrAuthSourceNotActivated
			}

			authenticator, ok := source.Cfg.(PasswordAuthenticator)
			if !ok {
				return nil, nil, smtp.ErrUnsupportedLoginType
			}

			user, err := authenticator.Authenticate(ctx, user, user.LoginName, password)
			if err != nil {
				return nil, nil, err
			}

			// WARN: DON'T check user.IsActive, that will be checked on reqSign so that
			// user could be hint to resend confirm email.
			if user.ProhibitLogin {
				return nil, nil, user_model.ErrUserProhibitLogin{UID: user.ID, Name: user.Name}
			}

			return user, source, nil
		}
	}

	sources, err := db.Find[auth.Source](ctx, auth.FindSourcesOptions{
		IsActive: optional.Some(true),
	})
	if err != nil {
		return nil, nil, err
	}

	for _, source := range sources {
		if !source.IsActive {
			// don't try to authenticate non-active sources
			continue
		}

		authenticator, ok := source.Cfg.(PasswordAuthenticator)
		if !ok {
			continue
		}

		authUser, err := authenticator.Authenticate(ctx, nil, username, password)

		if err == nil {
			if !authUser.ProhibitLogin {
				return authUser, source, nil
			}
			err = user_model.ErrUserProhibitLogin{UID: authUser.ID, Name: authUser.Name}
		}

		if user_model.IsErrUserNotExist(err) {
			log.Debug("Failed to login '%s' via '%s': %v", username, source.Name, err)
		} else {
			log.Warn("Failed to login '%s' via '%s': %v", username, source.Name, err)
		}
	}

	if isEmail {
		return nil, nil, user_model.ErrEmailAddressNotExist{Email: username}
	}

	return nil, nil, user_model.ErrUserNotExist{Name: username}
}
