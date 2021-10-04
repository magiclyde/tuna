/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/10/5 上午12:19
 * @note:
 */

package util

import "regexp"

// ValidatePhone 校验手机号格式
func ValidatePhone(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(16[5,6])|(17[0,3,5-8])|(18[0-9])|(19[1,3,5,8,9]))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}
