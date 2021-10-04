/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/10/4 下午10:37
 * @note: 获取微信网页授权链接
 */

package oa_web_apps

import (
	"github.com/magiclyde/tuna/util"
	"net/url"
)

// GetAuthorizeUrl https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html
func GetAuthorizeUrl(appId, redirectUri, scope, state string) string {
	base := "https://open.weixin.qq.com/connect/oauth2/authorize?"
	query := util.StrBuilder(
		"appid=", appId, "&",
		"redirect_uri=", url.QueryEscape(redirectUri), "&",
		"response_type=code", "&",
		"scope=", scope, "&",
		"state=", state,
	)
	fragment := "#wechat_redirect"
	return util.StrBuilder(base, query, fragment)
}
