package net

import (
	"context"
	"os/exec"
	"runtime"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

//网络情况探明

// 是否已经连接上网
// return false:断网中 true:已经联网
func IsLinkNet() bool {
	var hosts = []string{"baidu.com", "163.com", "ynmdiot.com", "qq.com", "iot.kyunmao.cn"}
	for _, host := range hosts {
		cmd := exec.Command("ping", "-c", "1", "-W", "5", host)
		if runtime.GOOS == "windows" {
			cmd = exec.Command("ping", host, "-w", "5")
		}
		err := cmd.Run()
		if err == nil {
			return true
		} else {
			g.Log().Debug(context.Background(), "net error", err, cmd)
		}
	}
	return false
}

func GetHostUrl(r *ghttp.Request) string {
	if r.Request.Proto == "HTTP/1.1" {
		return "http://" + r.Request.Host
	}
	return "https://" + r.Request.Host
}
