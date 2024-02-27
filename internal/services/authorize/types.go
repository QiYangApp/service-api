// Copyright 2019 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package authorize

import "service-api/internal/app/http/validator"

type Authorizing interface {
	PreAuthData(req validator.PreAuthDataRequest) any

	PreAuthVerify(req validator.PreAuthVerifyRequest) any

	Authorizing(req validator.AuthorizingRequest) any

	Authorized(req validator.AuthorizedRequest) any
}
