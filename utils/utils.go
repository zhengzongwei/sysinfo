package utils

import "fmt"

//func BytesToSize(bytes uint64) string {
//	const (
//		_  = iota
//		KB = 1 << (10 * iota)
//		MB
//		GB
//		TB
//		PB
//		EB
//	)
//
//	switch {
//	case bytes >= EB:
//		return fmt.Sprintf("%.2f EB", float64(bytes)/float64(EB))
//	case bytes >= PB:
//		return fmt.Sprintf("%.2f PB", float64(bytes)/float64(PB))
//	case bytes >= TB:
//		return fmt.Sprintf("%.2f TB", float64(bytes)/float64(TB))
//	case bytes >= GB:
//		return fmt.Sprintf("%.2f GB", float64(bytes)/float64(GB))
//	case bytes >= MB:
//		return fmt.Sprintf("%.2f MB", float64(bytes)/float64(MB))
//	case bytes >= KB:
//		return fmt.Sprintf("%.2f KB", float64(bytes)/float64(KB))
//	default:
//		return fmt.Sprintf("%d B", bytes)
//	}
//}

func ConvertSize(size uint64, sizeType string) string {
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
	)

	var divisor uint64
	switch sizeType {
	case "B", "Bytes":
		divisor = 1
	case "KB":
		divisor = KB
	case "MB":
		divisor = MB
	case "GB":
		divisor = GB
	case "TB":
		divisor = TB
	case "PB":
		divisor = PB
	case "EB":
		return fmt.Sprintf("%.2f EB", float64(size)/float64(EB))
	default:
		return ""
	}

	result := float64(size) / float64(divisor)
	return fmt.Sprintf("%.2f %s", result, sizeType)
}
