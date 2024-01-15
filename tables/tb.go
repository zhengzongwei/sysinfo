package tables

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/shirou/gopsutil/v3/net"
	"os"
	"strings"
	"sysinfo/cores"
	"sysinfo/utils"
)

var cpuTable = tablewriter.NewWriter(os.Stdout)
var memTable = tablewriter.NewWriter(os.Stdout)
var hostTable = tablewriter.NewWriter(os.Stdout)
var netTable = tablewriter.NewWriter(os.Stdout)
var diskTable = tablewriter.NewWriter(os.Stdout)

func init() {
	// Initialize CPU table header
	cpuTable.SetHeader([]string{"CPU", "Name", "VendorID", "Family", "Model", "Cores", "Mhz"})

	// Initialize MEM table header
	memTable.SetHeader([]string{"Total", "Available", "Used", "UsedPercent", "Free", "Active", "Inactive", "Wired"})
	hostTable.SetHeader([]string{"HostName", "UpTime", "BootTime", "Procs", "OS", "Platform", "platformFamily",
		"platformVersion", "KernelArch", "KerverVersion", "virtualizationSystem", "virtualizationRole", "hostId"})

	netTable.SetHeader([]string{"name", "mac", "mtu", "addrs"})
	diskTable.SetHeader([]string{"path", "fstype", "total", "free", "used", "usedPercent"})
}
func CPUTable() {

	CPUInfo := cores.GetCPUInfo()
	for _, info := range CPUInfo {
		row := []string{
			fmt.Sprintf("CPU %d", info.CPU),
			info.ModelName,
			info.VendorID,
			fmt.Sprint(info.Family),
			fmt.Sprint(info.Model),
			fmt.Sprint(info.Cores),
			fmt.Sprint(info.Mhz),
		}
		cpuTable.Append(row)
	}
	cpuTable.Render()
}

func MEMTable() {
	MEMInfo := cores.GetMemInfo()
	row := []string{
		utils.ConvertSize(MEMInfo.Total, "MB"),
		utils.ConvertSize(MEMInfo.Available, "MB"),
		utils.ConvertSize(MEMInfo.Used, "MB"),
		fmt.Sprintf("%.2f%%", MEMInfo.UsedPercent),
		utils.ConvertSize(MEMInfo.Free, "MB"),
		utils.ConvertSize(MEMInfo.Active, "MB"),
		utils.ConvertSize(MEMInfo.Inactive, "MB"),
		utils.ConvertSize(MEMInfo.Wired, "MB"),
	}
	memTable.Append(row)
	memTable.Render()
}

func HostTable() {
	HostInfo := cores.GetHostInfo()
	row := []string{
		HostInfo.Hostname,
		fmt.Sprint(HostInfo.Uptime),
		fmt.Sprint(HostInfo.BootTime),
		fmt.Sprint(HostInfo.Procs),
		HostInfo.OS,
		HostInfo.Platform,
		HostInfo.PlatformFamily,
		HostInfo.PlatformVersion,
		HostInfo.KernelArch,
		HostInfo.KernelVersion,
		HostInfo.VirtualizationSystem,
		HostInfo.VirtualizationRole,
		HostInfo.HostID,
	}
	hostTable.Append(row)
	hostTable.Render()
}

func NetTable() {
	NetInfo := cores.GetNetInfo()
	for _, info := range NetInfo {
		row := []string{
			info.Name,
			info.HardwareAddr,
			fmt.Sprint(info.MTU),
			joinAddresses(info.Addrs),
		}
		netTable.Append(row)
	}
	netTable.Render()

}

func DiskTable() {
	DiskInfo := cores.GetDiskInfo()
	for _, info := range DiskInfo {
		row := []string{
			info.Path,
			info.Fstype,
			utils.ConvertSize(info.Total, "GB"),
			utils.ConvertSize(info.Free, "GB"),
			utils.ConvertSize(info.Used, "GB"),
			fmt.Sprintf("%.2f%%", info.UsedPercent),
		}
		diskTable.Append(row)
	}
	diskTable.Render()
}

func joinAddresses(addrs net.InterfaceAddrList) string {
	var addrStrings []string
	for _, addr := range addrs {
		addrStrings = append(addrStrings, addr.Addr)
	}
	return strings.Join(addrStrings, "\n")

}
