package tables

import (
	"Golang/utils"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

func CPUInfoTable() {

	CPUInfo := utils.GetCPUInfo()
	// 创建表格对象
	table := tablewriter.NewWriter(os.Stdout)

	// 设置表格头部
	table.SetHeader([]string{"CPU", "Name", "VendorID", "Family", "Model", "Cores", "Mhz"})

	// 将 CPU 信息添加到表格
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
