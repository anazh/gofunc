package platform

import (
	"encoding/json"

	"github.com/silenceper/wechat/v2/officialaccount"
)

type UrlToken struct {
	Url     string
	Account *officialaccount.OfficialAccount
}

// out to be json marshaled bytes
func (u *UrlToken) CreateJSToken() ([]byte, error) {
	js := u.Account.GetJs()
	v, err := js.GetConfig(u.Url)
	if err != nil {
		return nil, err
	}
	return json.Marshal(v)
}
