package authorize

import (
	"github.com/gin-gonic/gin"
	"service-api/src/app/entity/http"
	"service-api/src/core/helpers/response"
)

type AuthorizedType[T http.VerifyType] interface {
	Authorizing(p T, c *gin.Context) *response.Response[any]

	Authorized(p T, c *gin.Context) *response.Response[any]
}
