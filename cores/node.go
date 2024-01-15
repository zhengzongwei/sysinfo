package cores

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
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

func GetDiskInfo() []disk.UsageStat {
	var DiskInfo []disk.UsageStat
	DiskParts, err := disk.Partitions(false)
	if err != nil {
		fmt.Printf("获取磁盘分区信息错误：%v", err)
		return nil
	}
	for _, DiskPart := range DiskParts {
		diskUsage, err := disk.Usage(DiskPart.Mountpoint)
		if err != nil {
			fmt.Printf("获取磁盘信息错误：%v", err)
			continue
		}
		DiskInfo = append(DiskInfo, *diskUsage)
	}
	return DiskInfo
}

func GetNetInfo() []net.InterfaceStat {
	NETInfo, err := net.Interfaces()
	if err != nil {
		fmt.Printf("获取网络信息错误：%v", err)
		return nil
	}
	return NETInfo

}

func GetHostInfo() *host.InfoStat {
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
