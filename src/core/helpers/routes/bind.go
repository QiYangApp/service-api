package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service-api/src/core/helpers/response"
)

type HandlerFunc[T interface{}, R interface{}] func(T, *gin.Context) *response.Response[R]

func Bind[T interface{}, R interface{}](handlerFunc HandlerFunc[T, R]) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p T
		var pt string

		pt = c.GetHeader("Content-Type")

		switch pt {
		case "application/json":
			if err := c.ShouldBindJSON(&p); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			break
		case "application/xml":
			if err := c.ShouldBindXML(&p); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			break
		case "application/form":
			if err := c.ShouldBindXML(&p); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			break
		default:
			if err := c.ShouldBind(&p); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			break
		}

		resp := handlerFunc(p, c)

		c.JSON(resp.Code, resp)
		c.Abort()
	}
}
