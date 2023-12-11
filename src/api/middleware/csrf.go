package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"service-api/src/core/helpers/response"
	"service-api/src/core/logger"
)

/**
// 单路由使用
router.POST("/", middleware.CSRFMiddleware(), func(c *gin.Context) {
	// Handle POST request
})

// 分组路由
protected := router.Group("/")
protected.Use(middleware.CSRFMiddleware())

protected.POST("/", func(c *gin.Context) {
	// Handle POST request
})
*/

const (
	csrfCookieName = "csrf_token"
	csrfHeaderName = "X-CSRF-Token"
)

func CSRF() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get CSRF token from cookie or generate a new one
		csrfToken, err := c.Cookie(csrfCookieName)
		if err != nil || csrfToken == "" {
			logger.S().Errorln("url: %s, err: %v", c.Request.URL, zap.Error(err.(error)))

			csrfToken, err = GetCSRFToken()
			if err != nil {
				c.AbortWithStatusJSON(
					http.StatusInternalServerError,
					response.RFail(
						c,
						err,
						http.StatusInternalServerError,
						"",
					).ToStruct(),
				)
				return
			}
			c.SetCookie(csrfCookieName, csrfToken, 0, "", "", false, true)
		}

		// Set CSRF token in header
		c.Writer.Header().Set(csrfHeaderName, csrfToken)

		// Check CSRF token for POST, PUT and DELETE requests
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
			referer := c.Request.Header.Get(csrfHeaderName)
			if referer == "" {
				logger.S().Errorln("url: %s, referer is empty", c.Request.URL)

				c.AbortWithStatusJSON(
					http.StatusBadRequest,
					response.RFail(
						c,
						err,
						http.StatusInternalServerError,
						"missing referer header",
					).ToStruct(),
				)
				return
			}

			err := CheckCSRF(csrfToken, c.Request)
			if err != nil {
				c.AbortWithStatusJSON(
					http.StatusForbidden,
					response.RFail(
						c,
						err,
						http.StatusForbidden,
						"",
					).ToStruct(),
				)
				return
			}
		}

		// Continue processing the request
		c.Next()
	}
}

// GetCSRFToken returns a new CSRF token.
func GetCSRFToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// CheckCSRF returns an error if the CSRF token does not match the value in the request header.
func CheckCSRF(csrfToken string, r *http.Request) error {
	if csrfToken == "" {
		return http.ErrNoCookie
	}

	csrfHeader := r.Header.Get(csrfHeaderName)
	if csrfToken != csrfHeader {
		return http.ErrNotSupported
	}

	return nil
}
