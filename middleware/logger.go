package middleware

import (
	"admin-api/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

func Logger() gin.HandlerFunc {
	logger := log.Log()
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		doTime := endTime.Sub(startTime) / time.Millisecond // 执行耗时
		reqMethod := c.Request.Method                       // 请求方式
		reqUri := c.Request.RequestURI                      // 请求URI
		header := c.Request.Header                          // 请求头
		proto := c.Request.Proto                            // 请求协议
		statusCode := c.Writer.Status()                     // 状态码
		clientIP := c.ClientIP()                            // 客户端IP
		err := c.Err()
		body, _ := io.ReadAll(c.Request.Body) // 请求体
		item := fmt.Sprintf("code: %d, time: %v, clientIP: %s, method: %s, uri: %s, header: %s, proto: %s, body: %s, err: %v", statusCode, doTime, clientIP, reqMethod, reqUri, header, proto, body, err)
		logger.Info(item)
	}
}
