package admin

import "github.com/gin-gonic/gin"

type AdminRouteRegister struct {
	route *gin.RouterGroup
}

func (s *AdminRouteRegister) Handle(r *gin.Engine) {
	s.route = r.Group("/admin")

	(new(apiRouteHandle)).Handle(s.route)
	(new(testRouteHandle)).Handle(s.route)

}
