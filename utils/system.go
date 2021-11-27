// Package utils system
package utils

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// GetOSName - 実行中のOS名を取得する
func GetOSName() string {
	return runtime.GOOS
}

// IsPrintVersion - バージョン番号を表示するか判定する
func IsPrintVersion(args ...string) bool {
	var isPrint bool

	// isPrint is true when commandline argument includes -v or -version
	var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	CommandLine.BoolVar(&isPrint, "v", false, "show version")
	CommandLine.BoolVar(&isPrint, "version", false, "show version")
	CommandLine.Parse(args)

	return isPrint
}

// PrintVersion - バージョン番号を表示する
func PrintVersion(version string) {
	fmt.Println("version:", version)
}

// GetRandomSeed - 乱数生成時に使用するシード値を取得する
func GetRandomSeed() int64 {
	return time.Now().UnixNano()
}

// JoinFilePath - ファイルパスを結合する
func JoinFilePath(elem []string) string {
	return filepath.Join(elem...)
}
