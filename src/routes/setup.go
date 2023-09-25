package routes

import (
	"github.com/gin-gonic/gin"
	"service-api/src/routes/admin"
	"service-api/src/routes/authorize"
)

type RouterRegisterMange interface {
	Handle(r *gin.Engine)
}

type RouterRegister interface {
	Handle(r *gin.RouterGroup)
}

func SetupRoutes(router *gin.Engine) {
	// 定义路由
	(new(admin.AdminRouteRegister)).Handle(router)
	(new(authorize.AuthorizeRouteRegister)).Handle(router)
}
