/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/7/6 上午10:59
 * @note:
 */

package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/magiclyde/tuna/jwt"
)

// Authorization authorization with jwt token
func Authorization(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("token")
		if tokenString == "" {
			c.AbortWithError(401, errors.New("Unauthorized"))
			return
		}
		tokenString, err := jwt.RefreshToken(key, tokenString)
		if err != nil {
			c.AbortWithError(401, errors.New("Unauthorized"))
			return
		}
		c.Header("token", tokenString)
		c.Next()
	}
}
