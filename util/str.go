/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/10/4 下午10:32
 * @note:
 */

package util

import (
	"reflect"
	"strings"
	"unsafe"
)

// StrBuilder join string
func StrBuilder(ss ...string) string {
	var b strings.Builder
	for _, s := range ss {
		b.WriteString(s)
	}
	return b.String()
}

// S2B converts string to a byte slice without memory allocation
func S2B(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// B2S converts a byte slice to string without memory allocation.
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
