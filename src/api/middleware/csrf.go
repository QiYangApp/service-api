package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
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
			csrfToken, err = GetCSRFToken()
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
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
				_ = c.AbortWithError(http.StatusBadRequest, errors.New("missing referer header"))
				return
			}

			err := CheckCSRF(csrfToken, c.Request)
			if err != nil {
				c.AbortWithStatus(http.StatusForbidden)
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
