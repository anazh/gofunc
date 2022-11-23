package net

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func NetInit() {
	Nets = NetStatus{}
	go Nets.netCheck()
	g.Log().Debug(context.Background(), "net init end")
}
