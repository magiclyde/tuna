/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/10/4 下午10:29
 * @note: 拉取用户信息(需 scope 为 snsapi_userinfo)
 */

package oa_web_apps

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/magiclyde/tuna/util"
	"github.com/magiclyde/tuna/weixin"
)

// GetUserInfo
func GetUserInfo(accessToken, openId string) (*weixin.UserInfoRsp, error) {
	apiUrl := util.StrBuilder("https://api.weixin.qq.com/sns/userinfo?",
		"access_token=", accessToken, "&",
		"openid=", openId,
	)
	body, err := util.HTTPGet(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("weixin.oauth.GetUserInfo, HTTPGet err: %s", err.Error())
	}
	var rsp weixin.UserInfoRsp
	if err := jsoniter.Unmarshal(body, &rsp); err != nil {
		return nil, fmt.Errorf("weixin.oauth.GetUserInfo, jsoniter.Unmarshal err: %s", err.Error())
	}
	if rsp.ErrCode != 0 {
		return nil, fmt.Errorf("weixin.oauth.GetUserInfo, errcode: %d, errmsg: %s", rsp.ErrCode, rsp.ErrMsg)
	}
	return &rsp, nil
}
