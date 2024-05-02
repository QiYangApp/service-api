package middlewares

import "github.com/gin-gonic/gin"

type Middleware func() gin.HandlerFunc
