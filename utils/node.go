package utils

import (
	"github.com/shirou/gopsutil/v3/cpu"
)

func GetCPUInfo() []cpu.InfoStat {
	CPUInfo, _ := cpu.Info()
	return CPUInfo
}
