package env

import (
	"os"

	"github.com/anazh/gofunc/encryp"
	"github.com/yumaojun03/dmidecode"
)

// 获取程序执行文件所在的目录
func GetExePwd() string {
	data, err := os.Getwd()
	if err == nil {
		return data
	}
	return ""
}

// linux:注意这个函数需要以root用户运行
// winodws:也可以获取
// CPUID+CPUNUM+硬盘序列号
func GetLinuxHardwareID() string {
	//获取CPUID
	dmi, err := dmidecode.New()
	if err != nil {
		return ""
	}
	infos, err := dmi.Processor() //windows是一起输出；linux是每个核心单独输出；但是不管是单个还是多个输出可以简单的将core count相加
	if err != nil {
		return ""
	}
	boardInfos, _ := dmi.BaseBoard()
	return encryp.MD5(infos[0].ID.String()) + encryp.MD5(boardInfos[0].SerialNumber)
}
