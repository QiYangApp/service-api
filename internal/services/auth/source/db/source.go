// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package db

import (
	"context"
	"ent/models"
	authtype "ent/types/auth"
	authmodel "service-api/internal/repo/auth"
)

// Source is a password authentication service
type Source[T any] struct {
	Val T
}

func (source *Source[T]) Value() T {
	return source.Val
}

// Authenticate queries if login/password is valid against the PAM,
// and create a local user if success when enabled.
func (source *Source[T]) Authenticate(ctx context.Context, user *models.User, login, password string) (*models.User, error) {
	return Authenticate(ctx, user, login, password)
}

func init() {
	authmodel.RegisterTypeConfig(authtype.NoType, &Source[any]{})
	authmodel.RegisterTypeConfig(authtype.Plain, &Source[any]{})
}
