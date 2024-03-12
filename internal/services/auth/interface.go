// Copyright 2019 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package auth

import (
	"context"
	"net/http"
	"service-api/internal/ent"
)

// DataStore represents a data store
type DataStore any

// SessionStore represents a session store
type SessionStore any

// Method represents an authentication method (plugin) for HTTP requests.
type Method interface {
	// Verify tries to verify the authentication data contained in the request.
	// If verification is successful returns either an existing user object (with id > 0)
	// or a new user object (with id = 0) populated with the information that was found
	// in the authentication data (username or email).
	// Second argument returns err if verification fails, otherwise
	// First return argument returns nil if no matched verification condition
	Verify(http *http.Request, w http.ResponseWriter, store DataStore, sess SessionStore) (*ent.User, error)

	Name() string
}

// PasswordAuthenticator represents a source of authentication
type PasswordAuthenticator interface {
	Authenticate(ctx context.Context, user *ent.User, login, password string) (*ent.User, error)
}

// LocalSkipper represents a source of authentication that can skip local 2fa
type LocalSkipper interface {
	IsSkipLocal() bool
}

// SynchronizableSource represents a source that can synchronize users
type SynchronizableSource interface {
	Sync(ctx context.Context, updateExisting bool) error
}
