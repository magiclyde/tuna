/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/10/4 下午10:39
 * @note: 获取公众号的全局唯一接口调用凭据 (access_token)
 */

package weixin

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/magiclyde/tuna/util"
)

// GetAccessToken https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html
func GetAccessToken(appid, secret string) (*GetAccessTokenRsp, error) {
	apiUrl := util.StrBuilder("https://api.weixin.qq.com/cgi-bin/token?",
		"grant_type=client_credential&",
		"appid=", appid, "&",
		"secret=", secret,
	)
	body, err := util.HTTPGet(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("weixin.GetAccessToken, HTTPGet err: %s", err.Error())
	}
	var rsp GetAccessTokenRsp
	if err := jsoniter.Unmarshal(body, &rsp); err != nil {
		return nil, fmt.Errorf("weixin.GetAccessToken, jsoniter.Unmarshal err: %s", err.Error())
	}
	if rsp.ErrCode != 0 {
		return nil, fmt.Errorf("weixin.GetAccessToken, errcode: %d, errmsg: %s", rsp.ErrCode, rsp.ErrMsg)
	}
	return &rsp, nil
}
