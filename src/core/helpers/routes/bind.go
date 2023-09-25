package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	verifyType "service-api/src/app/entity/http"
)

type HandlerFunc[P verifyType.VerifyType, R *gin.Context] func(*gin.Context, P) R

func Bind[P verifyType.VerifyType, R *gin.Context](handlerFunc HandlerFunc[P, R]) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p P
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
