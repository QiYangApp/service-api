package middlewares

import (
	"framework/config"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)

func Secure() gin.HandlerFunc {
	return secure.New(secure.Config{
		IsDevelopment:         config.Client.GetBool("debug"),
		AllowedHosts:          []string{config.Client.GetString("domain"), config.Client.GetString("addr")},
		SSLRedirect:           false,
		SSLHost:               config.Client.GetString("domain"),
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
