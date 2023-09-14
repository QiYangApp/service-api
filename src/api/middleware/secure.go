package middleware

import (
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
	"service-api/src/core/config"
)

func Secure() gin.HandlerFunc {
	return secure.New(secure.Config{
		IsDevelopment:         config.Instance.IsDebug(),
		AllowedHosts:          []string{config.Instance.GetDomain()},
		SSLRedirect:           false,
		SSLHost:               config.Instance.GetDomain(),
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
		IENoOpen:              true,
		ReferrerPolicy:        "strict-origin-when-cross-origin",
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
	})
}
