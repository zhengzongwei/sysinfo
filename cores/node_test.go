package cores

import (
	"testing"
)

func TestGetCPUInfo(t *testing.T) {
	CPUInfo := GetCPUInfo()
	IndentJSON(CPUInfo)
}

func TestGetMemInfo(t *testing.T) {
	MEMInfo := GetMemInfo()
	IndentJSON(MEMInfo)
}

func TestGetHostInfo(t *testing.T) {
	HostInfo := GetHostInfo()
	IndentJSON(HostInfo)
}
