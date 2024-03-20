// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2019 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package auth

import (
	"context"
	"ent/models"
	"ent/types/auth"
	"strings"
)

func UserSingIn(ctx context.Context, username, nickname string) (*models.User, *models.Source, error) {
	trimmedUsername := strings.TrimSpace(username)

	if trimmedUsername == "" {
		return nil, nil, auth.ErrUserNotExist{Name: username}
	}

	sourceType = auth.SMTP
	if strings.Contains(trimmedUsername, "@") {

	}

	return nil, nil, nil
}
