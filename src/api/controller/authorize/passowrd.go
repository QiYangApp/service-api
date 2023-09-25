package authorize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"service-api/src/api/controller"
	"service-api/src/app/entity/http"
	"service-api/src/core/helpers/response"
)

type PasswordController[P http.VerifyType] struct {
	AuthorizedType[P]
	controller.AbstractController
}

func (PasswordController[P]) Authorizing(p P, c *gin.Context) *response.Response[bool] {
	fmt.Println(p)

	return response.RSuccess[bool](c, true)
}

func (PasswordController[P]) Authorized(p P, c *gin.Context) *response.Response[bool] {
	return response.RSuccess[bool](c, true)
}
