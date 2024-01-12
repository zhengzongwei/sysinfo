package tables

import (
	"Golang/cores"
	"Golang/utils"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

func CPUInfoTable() {

	// 创建表格对象
	table := tablewriter.NewWriter(os.Stdout)

	// 设置表格头部
	table.SetHeader([]string{"CPU", "Name", "VendorID", "Family", "Model", "Cores", "Mhz"})

	// 将 CPU 信息添加到表格

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
		table.Append(row)
	}
	table.Render()
}

func MEMInfoTable() {
	// 创建表格对象
	table := tablewriter.NewWriter(os.Stdout)

	// 设置表格头部
	table.SetHeader([]string{"Total", "Available", "used", "usedPercent", "free", "active", "inactive", "wired"})
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
	table.Append(row)
	table.Render()
}
