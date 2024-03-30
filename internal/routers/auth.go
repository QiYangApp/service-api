package routers

import (
	"framework/router"
	"github.com/gin-gonic/gin"
	"service-api/internal/app/api/v1/auth"
)

type AuthRouter struct {
}

func (*AuthRouter) Handle(r *gin.RouterGroup) {
	// 获取前置数据
	r.GET("/v1/auth/signIn", router.Bind(auth.SignIn))
	// 登录
	r.GET("/v1/auth/signInPost", router.Bind(auth.SignInPost))
}
