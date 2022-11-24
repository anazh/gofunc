package os

import (
	"context"
	"time"

	"encoding/json"

	"github.com/gogf/gf/v2/os/gproc"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

// MemOut 计算机信息
type MemOut struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"usedPercent"`
	CUsed       float64 `json:"cpuUsed"`
	DiskAll     uint64  `json:"diskAllGB"`
	DiskFee     uint64  `json:"diskFreeGB"`
	CpuTemp     string  `json:"cpu_temp"`
}

// MemInfo 内存等系统信息
func MemInfo() []byte {
	var out MemOut
	v, _ := mem.VirtualMemory()
	cc, _ := cpu.Percent(time.Second, false)
	d, _ := disk.Usage("/")
	out.Free = v.Free
	out.UsedPercent = v.UsedPercent
	out.Total = v.Total
	out.CUsed = cc[0]
	out.DiskAll = d.Total / 1000 / 1000 / 1000
	out.DiskFee = d.Free / 1000 / 1000 / 1000
	out.CpuTemp, _ = gproc.ShellExec(context.Background(), "cat /sys/class/thermal/thermal_zone0/temp") //执行命令
	i, _ := json.Marshal(out)
	return i
}
