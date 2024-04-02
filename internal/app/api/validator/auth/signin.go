// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2018 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package auth

type SignInForm struct {
	UserName string `form:"userName" json:"userName" binding:"required,max=254"`
	Passwd   string `form:"passwd" json:"passwd" binding:"max=255"`
	Remember bool   `form:"remember" json:"remember"`
}

type SignIn2FAJump struct {
	WEBAUTHN  bool `json:"WEBAUTHN"`
	TwoFactor bool `json:"TWO_FACTOR"`
}

type SignInVerifyError struct {
	ProhibitLogin     bool `json:"prohibitLogin,omitempty"`
	ActiveYourAccount bool `json:"activeYourAccount,omitempty"`
}
