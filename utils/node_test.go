package utils

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetCPUInfo(t *testing.T) {
	CPUInfo := GetCPUInfo()
	CPUInfoJSON, err := json.MarshalIndent(CPUInfo, "", "  ")
	if err != nil {
		fmt.Println("JSON marshaling failed:", err)
		return
	}
	fmt.Printf("%s", CPUInfoJSON)
}
