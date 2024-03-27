// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2018 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package auth

type SignInForm struct {
	UserName string `form:"userName" json:"userName" binding:"required,max=254"`
	Passwd   string `form:"passwd" json:"passwd" binding:"max=255"`
	Remember bool   `form:"remember" json:"remember"`
}
