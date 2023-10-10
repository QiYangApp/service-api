package authorize

import (
	"github.com/gin-gonic/gin"
	"service-api/src/app/entity/http"
	"service-api/src/core/helpers/response"
)

type AuthorizedType[P http.VerifyType] interface {
	Authorizing(p P, c *gin.Context) *response.Response[any]

	Authorized(p P, c *gin.Context) *response.Response[any]
}
