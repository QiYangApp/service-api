// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2018 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package auth

import (
	"frame/modules/router"
	"frame/modules/translate"
	"github.com/gin-gonic/gin"
	"service-api/internal/modules/setting"
	"service-api/resources/translate/messages"
)

type SignInForm struct {
	UserName string `form:"userName" json:"userName" binding:"required,max=128"`
	Passwd   string `form:"passwd" json:"passwd" binding:"max=255"`
	Remember bool   `form:"remember" json:"remember"`
	Captcha  struct {
		Key    string                 `form:"key" json:"key"`
		Type   setting.CaptchaFeature `form:"type" json:"type"`
		Token  string                 `form:"token" json:"token"`
		Id     string                 `form:"id" json:"id"`
		Answer string                 `form:"answer" json:"answer"`
	} `form:"captcha" json:"captcha"`
}

func (SignInForm) GetMessage(c *gin.Context) router.ValidatorMessages {
	return router.ValidatorMessages{
		"userName.required": translate.MustGetMessage(c, messages.UserEmailEmpty.ID),
		"userName.max":      translate.MustGetMessage(c, messages.Fail.ID),
		"passwd.max":        translate.MustGetMessage(c, messages.Error.ID),
		"remember.required": translate.MustGetMessage(c, messages.Error.ID),
	}
}

type SignIn2FAJump struct {
	WEBAUTHN  bool `json:"WEBAUTHN"`
	TwoFactor bool `json:"TWO_FACTOR"`
}

type SignInVerifyError struct {
	ProhibitLogin     bool `json:"prohibitLogin,omitempty"`
	ActiveYourAccount bool `json:"activeYourAccount,omitempty"`
}
