// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package db

import (
	"service-api/internal/ent"

	"github.com/gin-gonic/gin"
)

// Source is a password authentication service
type Source struct{}

// FromDB fills up an OAuth2Config from serialized format.
func (source *Source) FromDB(bs []byte) error {
	return nil
}

// ToDB exports the config to a byte slice to be saved into database (this method is just dummy and does nothing for DB source)
func (source *Source) ToDB() ([]byte, error) {
	return nil, nil
}

// Authenticate queries if login/password is valid against the PAM,
// and create a local user if success when enabled.
func (source *Source) Authenticate(ctx gin.Context, user *ent.User, login, password string) (*ent.User, error) {
	return Authenticate(ctx, user, login, password)
}

func init() {
	auth.RegisterTypeConfig(auth.NoType, &Source{})
	auth.RegisterTypeConfig(auth.Plain, &Source{})
}
