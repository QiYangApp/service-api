package v1

import (
	"frame/modules/router"
	"service-api/internal/net/api/v1/auth"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
}

func (*AuthRouter) Handle(r *gin.RouterGroup) {

	// 获取前置数据
	// 登录
	router.GetBind(r, "v1/auth/signIn", auth.SignIn)
	router.PostBind(r, "v1/auth/signIn", auth.SignInPost)

	// email and password sign up
	router.GetBind(r, "v1/auth/signUp", auth.SignUp)
	router.PostBind(r, "v1/auth/signUp", auth.SignUpPost)

	// 2fa
	router.GetBind(r, "v1/auth/2fa", auth.SignInPost)
	router.PostBind(r, "v1/auth/2fa", auth.SignInPost)

}
