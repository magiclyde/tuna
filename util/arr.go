/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/10/5 上午12:20
 * @note:
 */

package util

// ReverseStringArray 数组倒序函数
func ReverseStringArray(arr *[]string) {
	var temp string
	length := len(*arr)
	for i := 0; i < length/2; i++ {
		temp = (*arr)[i]
		(*arr)[i] = (*arr)[length-1-i]
		(*arr)[length-1-i] = temp
	}
}
