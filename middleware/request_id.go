/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/7/6 上午10:39
 * @note:
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestId set trace id
func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqId := c.Request.Header.Get("X-Request-Id")
		if reqId == "" {
			reqId = uuid.New().String()
			c.Request.Header.Set("X-Request-Id", reqId)
		}
		c.Next()
	}
}
