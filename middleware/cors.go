package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/gin-gonic/gin"
)

// Cors 设置请求头中间件，解决跨域问题
func CorsGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置响应头
		c.Header("X-Custom-Header", "HeaderValue")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, User-Agent, Keep-Alive, X-Requested-With, If-Modified-Since, Cache-Control, X-Custom-Header")
		// 暴露自定义头部字段
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, Authorization, X-Custom-Header")

		// 如果是 OPTIONS 请求，直接返回 200
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		// 继续处理请求
		c.Next()
	}
}

// hertz版
func Cors(ctx context.Context, c *app.RequestContext) {
	c.Response.Header.Set("X-Custom-Header", "HeaderValue")
	c.Response.Header.Set("Access-Control-Allow-Origin", "*")
	c.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type, Authorization, User-Agent, Keep-Alive, X-Requested-With, If-Modified-Since, Cache-Control, X-Custom-Header")
	// 暴露自定义头部字段
	c.Response.Header.Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, Authorization, X-Custom-Header")
	if string(c.Request.Header.Method()) == consts.MethodOptions {
		c.AbortWithStatus(consts.StatusOK)
		return
	}
	c.Next(ctx)
}
