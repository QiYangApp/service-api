package group

import (
	"frame/modules/router"
	"service-api/internal/api/v1/auth"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
}

func (*AuthRouter) Handle(r *gin.RouterGroup) {

	// 获取前置数据
	router.GetBind(r, PrefixV1+"/signIn", auth.SignIn)

	// 登录
	router.PostBind(r, PrefixV1+"/signIn", auth.SignInPost)

	// 2fa
	router.GetBind(r, PrefixV1+"/2fa", auth.SignInPost)
	router.PostBind(r, PrefixV1+"/2fa", auth.SignInPost)

}
