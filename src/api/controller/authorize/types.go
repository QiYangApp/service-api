package authorize

import (
	"github.com/gin-gonic/gin"
	"service-api/src/app/entity/http"
)

type AuthorizedType[P http.VerifyType] interface {
	// Check 检测
	Check(p P, c *gin.Context) *gin.Context

	// Authorizing 授权前置
	Authorizing(p P, c *gin.Context) *gin.Context

	// Authorized 授权完成
	Authorized(p P, c *gin.Context) *gin.Context
}
