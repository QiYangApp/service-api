// Copyright 2019 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package auth

import (
	"context"
	"ent/models"
	"net/http"
)

// DataStore represents a data store
type DataStore interface {
	GetData() ContextData
}

type ContextData map[string]any

func (ds ContextData) GetData() ContextData {
	return ds
}

func (ds ContextData) MergeFrom(other ContextData) ContextData {
	for k, v := range other {
		ds[k] = v
	}
	return ds
}

// SessionStore represents a session store
type SessionStore interface {
	// Set sets value to given key in session.
	Set(any, any) error
	// Get gets value by given key in session.
	Get(any) any
	// Delete deletes a key from session.
	Delete() error
	// ID returns current session ID.
	ID() string
	// Release releases session resource and save data to provider.
	Release() error
	// Flush deletes all session data.
	Flush() error
}

// Method represents an authentication method (plugin) for HTTP requests.
type Method interface {
	// Verify tries to verify the authentication data contained in the request.
	// If verification is successful returns either an existing user object (with id > 0)
	// or a new user object (with id = 0) populated with the information that was found
	// in the authentication data (username or email).
	// Second argument returns err if verification fails, otherwise
	// First return argument returns nil if no matched verification condition
	Verify(http *http.Request, w http.ResponseWriter, store DataStore, sess SessionStore) (*models.User, error)

	Name() string
}

// PasswordAuthenticator represents a source of authentication
type PasswordAuthenticator interface {
	Authenticate(ctx context.Context, user *models.User, login, password string) (*models.User, error)
}

// LocalTwoFASkipper represents a source of authentication that can skip local 2fa
type LocalTwoFASkipper interface {
	IsSkipLocalTwoFA() bool
}

// SynchronizableSource represents a source that can synchronize users
type SynchronizableSource interface {
	Sync(ctx context.Context, updateExisting bool) error
}
