package authorize

import "github.com/gin-gonic/gin"

type AuthorizeRouteRegister struct {
	route *gin.RouterGroup
}

func (s *AuthorizeRouteRegister) Handle(r *gin.Engine) {
	s.route = r.Group("/authorizing")

	(new(loginRouteHandle)).Handle(s.route)

}
