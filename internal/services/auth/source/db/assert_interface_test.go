// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package db_test

import (
	authtype "ent/types/auth"
	"service-api/internal/services/auth"
	"service-api/internal/services/auth/source/db"
)

// This test file exists to assert that our Source exposes the interfaces that we expect
// It tightly binds the interfaces and implementation without breaking go import cycles

type sourceInterface interface {
	auth.PasswordAuthenticator
	authtype.ConfigType[any]
}

var _ (sourceInterface) = &db.Source[any]{}
