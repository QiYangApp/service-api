package middlewares

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"framework/cache"
	"framework/log"
	"framework/utils/optional"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type responseCache struct {
	Status int
	Header http.Header
	Data   []byte
}

var (
	PageCachePrefix = "api.response.cache"
)

func Cache(duration time.Duration, handle func(*gin.Context) string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var resp optional.Option[responseCache]
		var err error
		key := CreateKey(c.Request.URL.RequestURI())

		if resp, err = cache.Get[responseCache](key); err != nil {
			// 在请求处理之后记录响应信息
			log.Client.Sugar().Debugw(
				"cache",
				zap.String("url", c.Request.URL.String()),
				zap.Any("headers", c.Request.Header),
				zap.String("method", c.Request.Method),
				zap.Any("body", c.Request.Body),
				zap.Int("status", c.Writer.Status()),
				zap.String("cache key", key),
			)

			c.Writer = &cacheWriter{
				c.Writer,
				c.Writer.Header(),
				duration,
				key,
				c.Writer.Status(),
			}

			handle(c)

			// Drop caches of aborted contexts
			if c.IsAborted() {
				_ = cache.Del(key)
			}
		} else {
			if !resp.Has() {
				log.Client.Sugar().Warn("response cache data is empty")
				handle(c)
				return
			}

			c.Writer.WriteHeader(resp.Value().Status)
			for k, vals := range resp.Value().Header {
				for _, v := range vals {
					c.Writer.Header().Set(k, v)
				}
			}
			_, _ = c.Writer.Write(resp.Value().Data)

		}
	}
}

type cacheWriter struct {
	gin.ResponseWriter
	header   http.Header
	duration time.Duration
	key      string
	status   int
}

func (c *cacheWriter) Write(data []byte) (int, error) {
	ret, err := c.ResponseWriter.Write(data)
	if err != nil {
		return ret, err
	}

	if c.Status() < 300 {
		var resp optional.Option[responseCache]
		if _, err := cache.Get[responseCache](c.key); err == nil && resp.Has() {
			data = append(resp.Value().Data, data...)
		}

		val := responseCache{
			c.Status(),
			c.Header(),
			data,
		}

		_ = cache.SetNx(c.key, val, c.duration)
	}

	return ret, err
}

func (c *cacheWriter) WriteString(s string) (int, error) {
	ret, err := c.ResponseWriter.WriteString(s)

	if err == nil && c.Status() < 300 {
		val := responseCache{
			c.Status(),
			c.Header(),
			[]byte(s),
		}
		_ = cache.SetNx(c.key, val, c.duration)
	}

	return ret, err
}

func (c *cacheWriter) WriteHeader(code int) {
	c.status = code
	c.ResponseWriter.WriteHeader(code)
}

// CreateKey creates a package specific key for a given string
func CreateKey(u string) string {
	return urlEscape(PageCachePrefix, u)
}

func urlEscape(prefix, u string) string {
	//key := url.QueryEscape(u)
	//if len(key) > 200 {
	//	h := sha1.New()
	//	_, _ = io.WriteString(h, u)
	//	key = string(h.Sum(nil))
	//}

	path, err := getRequestUriIgnoreQueryOrder(u)
	if err != nil {
		return ""
	}

	h := sha1.New()
	h.Write([]byte(path))
	key := hex.EncodeToString(h.Sum(nil))

	var buffer bytes.Buffer
	buffer.WriteString(prefix)
	buffer.WriteString(":")
	buffer.WriteString(key)

	return buffer.String()
}

func getRequestUriIgnoreQueryOrder(requestURI string) (string, error) {
	parsedUrl, err := url.ParseRequestURI(requestURI)
	if err != nil {
		return "", err
	}

	values := parsedUrl.Query()

	if len(values) == 0 {
		return requestURI, nil
	}

	queryKeys := make([]string, 0, len(values))
	for queryKey := range values {
		queryKeys = append(queryKeys, queryKey)
	}
	sort.Strings(queryKeys)

	queryVals := make([]string, 0, len(values))
	for _, queryKey := range queryKeys {
		sort.Strings(values[queryKey])
		for _, val := range values[queryKey] {
			queryVals = append(queryVals, queryKey+"="+val)
		}
	}

	return parsedUrl.Path + "?" + strings.Join(queryVals, "&"), nil
}
