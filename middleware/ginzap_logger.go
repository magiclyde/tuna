/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/7/14 下午3:27
 * @note: custom gin logger with zap
 * @refer: https://github.com/gin-contrib/zap
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/syyongx/php2go"
	"go.uber.org/zap"
	"time"
)

func GinZapLogger(logger *zap.Logger, skipPaths []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		if php2go.InArray(path, skipPaths) {
			return
		}

		end := time.Now()
		latency := end.Sub(start)

		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("latency", latency),
		)
	}
}
