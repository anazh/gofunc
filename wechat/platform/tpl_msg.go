package platform

import (
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

type TplMsg struct {
	Account *officialaccount.OfficialAccount
}

type MsgContent struct {
	MsgTemplateCode     string
	Openid              string
	MsgUrl              string
	MiniprogramAppId    string
	MiniprogramPagepath string
	Data                map[string]*message.TemplateDataItem
}

func (t *TplMsg) ToSend(m *MsgContent) (bool, string) {
	msg := &message.TemplateMessage{}
	msg.TemplateID = m.MsgTemplateCode
	msg.ToUser = m.Openid
	msg.URL = m.MsgUrl
	msg.MiniProgram.AppID = m.MiniprogramAppId
	msg.MiniProgram.PagePath = m.MiniprogramPagepath
	msg.Data = m.Data
	tpl := t.Account.GetTemplate()
	_, err := tpl.Send(msg)
	if err == nil {
		return true, ""
	}
	return false, err.Error()
}
