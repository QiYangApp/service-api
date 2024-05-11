package routers

import (
	"frame/modules/router"
	"github.com/gin-gonic/gin"
	v1 "service-api/internal/app/api/v1"
)

type TestRouter struct {
}

func (*TestRouter) Handle(r *gin.RouterGroup) {
	// 获取前置数据
	r.GET("/v1/test", router.Bind(v1.TestSession))

}
