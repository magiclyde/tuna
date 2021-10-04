/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/10/4 下午10:55
 * @note: 微信服务器 token 验证
 */

package middleware

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

// WeixinTokenAuth
func WeixinTokenAuth(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		signature := c.Query("signature")
		timestamp := c.Query("timestamp")
		nonce := c.Query("nonce")

		//将 token、timestamp、nonce 三个参数进行字典序排序
		var tempArray = []string{token, timestamp, nonce}
		sort.Strings(tempArray)

		//将三个参数字符串拼接成一个字符串进行 sha1 加密
		var sha1String = ""
		for _, v := range tempArray {
			sha1String += v
		}
		h := sha1.New()
		h.Write([]byte(sha1String))
		sha1String = hex.EncodeToString(h.Sum([]byte("")))

		if sha1String != signature {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
