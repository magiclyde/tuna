/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/10/4 下午10:27
 * @note:
 */

package util

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

// HTTPTest
func HTTPTest(req *http.Request, router *gin.Engine) (body []byte, err error) {
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应 handler 接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应 body
	body, err = ioutil.ReadAll(result.Body)
	return
}
