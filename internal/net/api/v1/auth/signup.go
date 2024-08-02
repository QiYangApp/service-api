// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2018 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package auth

import (
	"frame/modules/log"
	"frame/modules/resp"
	"github.com/gin-gonic/gin"
	"service-api/internal/net/validator/auth"
)

func SignUp(ctx *gin.Context) {

}

func SignUpPost(ctx *gin.Context, form *auth.SignUpForm) *resp.Response {
	log.Sugar().Info(form)

	return resp.Success(ctx, form)
}
