// Copyright 2020 The Macaron Authors
// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var cookieConf http.Cookie

func SettingCookieConf(path, domain string, secure, httpOnly bool) {
	cookieConf.Path = path
	cookieConf.Domain = domain
	cookieConf.Secure = secure
	cookieConf.HttpOnly = httpOnly
}

// SetRedirectToCookie convenience function to set the RedirectTo cookie consistently
func SetRedirectToCookie(c *gin.Context, value string) {
	SetSiteCookie(c, "redirect_to", value, 0)
}

// DeleteRedirectToCookie convenience function to delete most cookies consistently
func DeleteRedirectToCookie(c *gin.Context) {
	SetSiteCookie(c, "redirect_to", "", -1)
}

// GetSiteCookie returns given cookie value from request header.
func GetSiteCookie(c *gin.Context, name string) string {
	cookie, err := c.Cookie(name)
	if err != nil {
		return ""
	}

	return cookie
}

// SetSiteCookie returns given cookie value from request header.
func SetSiteCookie(c *gin.Context, name, value string, maxAge int) {
	c.SetCookie(name, value, maxAge, cookieConf.Path, cookieConf.Domain, cookieConf.Secure, cookieConf.HttpOnly)
}
