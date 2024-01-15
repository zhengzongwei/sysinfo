package cores

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func GetCPUInfo() []cpu.InfoStat {
	CPUInfo, err := cpu.Info()
	if err != nil {
		fmt.Printf("获取CPU信息错误：%v", err)
		return nil
	}
	return CPUInfo
}

func GetMemInfo() *mem.VirtualMemoryStat {
	MEMInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("获取内存信息错误：%v", err)
		return nil
	}
	return MEMInfo
}

func GetDiskInfo() {

}

func GetHostInfo() *host.InfoStat {
	// 当前系统
	//fmt.Println(runtime.GOOS)
	//fmt.Println(runtime.GOARCH)
	HostInfo, err := host.Info()
	if err != nil {
		fmt.Printf("获取主机信息错误：%v", err)
		return nil
	}
	return HostInfo

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
