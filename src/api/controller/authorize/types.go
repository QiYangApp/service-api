package authorize

import (
	"github.com/gin-gonic/gin"
	"service-api/src/app/entity/http"
)

type AuthorizedType interface {
	// Check 检测
	Check(c *gin.Context, p http.ReqType) *gin.Context

	// Authorizing 授权前置
	Authorizing(c *gin.Context, p http.ReqType) *gin.Context

	// Authorized 授权完成
	Authorized(c *gin.Context, p http.ReqType) *gin.Context
}
