package platform

import (
	"errors"

	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/oauth"
)

type scope string

const (
	OnlyOpenid scope = "snsapi_base"
	UserInfo   scope = "snsapi_userinfo"
)

type Oauth struct {
	Account *officialaccount.OfficialAccount
	Url     string //oauth reback url
	State   string //oauth url param
}

// url oauth reback url
// state parameter
func (w *Oauth) WebOauthUrl(scopeGet scope) (string, error) {
	if w.Url == "" || w.State == "" {
		return "", errors.New("url or state is empty")
	}
	auth := oauth.NewOauth(w.Account.GetContext())
	return auth.GetRedirectURL(w.Url, string(scopeGet), w.State)
}

// 网页授权获取用户信息，通过code换取access_token
func (w *Oauth) GetResAccessToken(code string) (oauth.ResAccessToken, error) {
	auth := oauth.NewOauth(w.Account.GetContext())
	return auth.GetUserAccessToken(code)
}
