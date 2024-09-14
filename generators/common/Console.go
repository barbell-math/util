package common

import "fmt"


func PrintRunningInfo(fmtStr string, args ...any) {
	fmtStr = " |-" + fmtStr + "\n"
	fmt.Printf(fmtStr, args...)
}

func PrintRunningError(fmtStr string, args ...any) {
	fmtStr = " |-ERROR: " + fmtStr + "\n"
	fmt.Printf(fmtStr, args...)
}
