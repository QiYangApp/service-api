// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2018 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package auth

import (
	"frame/modules/resp"
	"github.com/gin-gonic/gin"
	"service-api/conf"
	"service-api/internal/net/validator/auth"
	captchaValidator "service-api/internal/net/validator/captcha"
	captchaService "service-api/internal/services/captcha"
)

func SignUp(ctx *gin.Context) *resp.Response {
	return resp.Success(ctx, "")
}

func SignUpPost(ctx *gin.Context, form *auth.SignUpForm, captchaForm *captchaValidator.CaptchaVerifyForm) *resp.Response {

	if {

	}

	if r := captchaService.Verify(ctx, conf.CaptchaFeatureSignUp, captchaForm.Token, captchaForm.Key, captchaForm.Answer, true); r != nil {
		return r
	}

	return resp.Success(ctx, form)
}
