package platform

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/config"
)

type PlatformInfo struct {
	AppId  string
	Secret string
	Token  string
	Aeskey string
}

func (p *PlatformInfo) New() *officialaccount.OfficialAccount {
	wc := wechat.NewWechat()
	cfg := &config.Config{
		AppID:          p.AppId,
		AppSecret:      p.Secret,
		Token:          p.Token,
		EncodingAESKey: p.Aeskey,
		Cache:          cache.NewMemory(),
	}
	return wc.GetOfficialAccount(cfg)
}
