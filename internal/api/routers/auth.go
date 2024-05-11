package routers

import (
	"frame/modules/router"
	"service-api/internal/app/api/v1/auth"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
}

func (*AuthRouter) Handle(r *gin.RouterGroup) {
	// 获取前置数据
	r.GET("/v1/auth/signIn", router.Bind(auth.SignIn))
	// 登录
	r.GET("/v1/auth/signInPost", router.Bind(auth.SignInPost))

	// 双重校验
	r.GET("/v1/auth/2fa", router.Bind(auth.SignIn))
	r.GET("/v1/auth/2faPost", router.Bind(auth.SignIn))
}
