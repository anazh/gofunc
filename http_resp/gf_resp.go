package http_resp

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type RespContent struct {
	con         context.Context
	ErrCode     int         `json:"error_code"`
	Msg         string      `json:"msg"`
	Data        interface{} `json:"data"`
	ServiceTime int64       `json:"service_time"`
}

func CreateResp(ctx context.Context) *RespContent {
	resp := &RespContent{
		con:         ctx,
		ErrCode:     0,
		Msg:         "",
		Data:        nil,
		ServiceTime: time.Now().Unix(),
	}
	return resp
}

// err return to client
func (r *RespContent) ErrResp(errcode, httpStatus int, errMsg string) {
	r.ErrCode = errcode
	r.Msg = errMsg
	r.ServiceTime = time.Now().Unix()
	g.RequestFromCtx(r.con).Response.WriteStatus(httpStatus)
	g.RequestFromCtx(r.con).Response.WriteJsonExit(r)
}

// success
func (r *RespContent) Success(data interface{}) {
	r.ErrCode = 0
	r.Msg = "ok"
	r.Data = data
	r.ServiceTime = time.Now().Unix()
	g.RequestFromCtx(r.con).Response.WriteJsonExit(r)
}
