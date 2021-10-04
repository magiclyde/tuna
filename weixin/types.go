/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/10/4 下午10:19
 * @note:
 */

package weixin

type CommRsp struct {
	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}

type GetAccessTokenRsp struct {
	CommRsp
	AccessToken  string `json:"access_token,omitempty"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	OpenId       string `json:"openid,omitempty"`
	Scope        string `json:"scope,omitempty"`
}

type UserInfoRsp struct {
	CommRsp
	OpenId     string `json:"openid,omitempty"`
	Nickname   string `json:"nickname,omitempty"`
	Sex        int    `json:"sex,omitempty"` // 用户性别，值为1时是男性，值为2时是女性，值为0时是未知
	Province   string `json:"province,omitempty"`
	City       string `json:"city,omitempty"`
	Country    string `json:"country,omitempty"`
	HeadImgUrl string `json:"headimgurl,omitempty"`
	UnionId    string `json:"unionid,omitempty"`
}
