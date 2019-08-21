package utils

import (
	"runtime"
)

// 実行中のOS名を取得する
func GetOSName() string {
	return runtime.GOOS
}
