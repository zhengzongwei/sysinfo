package cores

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func GetCPUInfo() []cpu.InfoStat {
	CPUInfo, _ := cpu.Info()
	return CPUInfo
}

func GetMemInfo() *mem.VirtualMemoryStat {
	MEMInfo, _ := mem.VirtualMemory()
	return MEMInfo
}

func IndentJSON(data interface{}) {
	DataJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("JSON marshaling failed:", err)
		return
	}
	fmt.Printf("%s", DataJSON)
}

//type nodeInfo struct {
//
//}
//func GetNodeInfo() []nodeInfo {
//
//}
