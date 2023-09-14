package middleware

/**
// 单个路由使用缓存
r.GET("/user/:id", middleware.CachePage(5*time.Minute, func(c *gin.Context) string {
	return c.Param("id")
}), func(c *gin.Context) {
	// 处理请求
})

// 分组路由使用缓存
admin := r.Group("/admin")
admin.Use(func(c *gin.Context) {
	// 获取全局缓存
	cache, ok := c.MustGet("cache").(*cache.Cache)
	if !ok {
		panic("cache middleware is required")
	}
	c.Set("cache", cache)
	c.Next()
})
*/

//func CachePage(duration time.Duration, keyGenerator func(*gin.Context) string) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		key := keyGenerator(c)
//		if key != "" {
//			if page, found := cache.Get(key); found {
//				fmt.Println("from cache")
//				c.Data(http.StatusOK, "text/html; charset=utf-8", page.([]byte))
//				return
//			}
//			c.Writer = &cacheWriter{
//				header:         c.Writer.Header(),
//				ResponseWriter: c.Writer,
//				duration:       duration,
//				key:            key,
//			}
//			c.Next()
//		} else {
//			c.Next()
//		}
//	}
//}
//
//type cacheWriter struct {
//	header http.Header
//	gin.ResponseWriter
//	duration time.Duration
//	key      string
//}
//
//func (c *cacheWriter) Write(data []byte) (int, error) {
//	if len(c.key) > 0 && len(data) > 0 {
//		cache.Set(c.key, data, c.duration)
//	}
//	return c.ResponseWriter.Write(data)
//}
