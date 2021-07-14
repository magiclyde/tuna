/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/7/14 下午3:27
 * @note: custom gin logger with zap
 * @refer: https://github.com/gin-contrib/zap
 */

package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/syyongx/php2go"
	"go.uber.org/zap"
	"io/ioutil"
	"time"
)

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w responseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func GinZapLogger(logger *zap.Logger, skipPaths []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		respWriter := &responseWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = respWriter

		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		if php2go.InArray(path, skipPaths) {
			return
		}

		end := time.Now()
		latency := end.Sub(start)

		reqId := c.Request.Header.Get("X-Request-Id")
		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		respBody := respWriter.body.String()

		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("latency", latency),
			zap.String("reqId", reqId),
			zap.ByteString("reqBody", reqBody),
			zap.String("respBody", respBody),
		)
	}
}
