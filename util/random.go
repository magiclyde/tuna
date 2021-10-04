/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/10/5 上午12:18
 * @note:
 */

package util

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return B2S(b)
}
