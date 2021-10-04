/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/10/4 下午10:19
 * @note: 通过 code 换取网页授权 access_token
 */

package oa_web_apps

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/magiclyde/tuna/util"
	"github.com/magiclyde/tuna/weixin"
)

// GetAccessToken
func GetAccessToken(code, appId, appSecret string) (*weixin.GetAccessTokenRsp, error) {
	apiUrl := util.StrBuilder("https://api.weixin.qq.com/sns/oauth2/access_token?",
		"appid=", appId, "&",
		"secret=", appSecret, "&",
		"code=", code, "&",
		"grant_type=authorization_code",
	)
	body, err := util.HTTPGet(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("weixin.oauth.GetAccessToken, HTTPGet err: %s", err.Error())
	}
	var rsp weixin.GetAccessTokenRsp
	if err := jsoniter.Unmarshal(body, &rsp); err != nil {
		return nil, fmt.Errorf("weixin.oauth.GetAccessToken, jsoniter.Unmarshal err: %s", err.Error())
	}
	if rsp.ErrCode != 0 {
		return nil, fmt.Errorf("weixin.oauth.GetAccessToken, errcode: %d, errmsg: %s", rsp.ErrCode, rsp.ErrMsg)
	}
	return &rsp, nil
}
