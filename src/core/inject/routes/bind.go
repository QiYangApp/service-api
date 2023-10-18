package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	verifyType "service-api/src/app/entity/http"
)

type HandlerFuncParams func(*gin.Context, verifyType.VerifyType) *gin.Context
type HandlerFuncDefault func(*gin.Context) *gin.Context

func Exce(fn gin.HandlerFunc) {

}

func Bind(handlerFunc HandlerFuncParams, p any) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var pt string

		pt = c.GetHeader("Content-Type")

		switch pt {
		case "application/json":
			err = c.ShouldBindJSON(&p)
			break
		case "application/xml":
			err = c.ShouldBindXML(&p)
			break
		case "application/form":
			err = c.ShouldBindXML(&p)
			break
		default:
			err = c.ShouldBind(&p)
			break
		}

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
		}

		handlerFunc(c, p)

		c.Abort()
	}
}
