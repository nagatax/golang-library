package utils

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

// 実行中のOS名を取得する
func GetOSName() string {
	return runtime.GOOS
}

// バージョン番号を表示するか判定する
func IsPrintVersion(args ...string) bool {
	var isPrint bool

	// isPrint is true when commandline argument includes -v or -version
	var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	CommandLine.BoolVar(&isPrint, "v", false, "show version")
	CommandLine.BoolVar(&isPrint, "version", false, "show version")
	CommandLine.Parse(args)

	return isPrint
}

// バージョン番号を表示する
func PrintVersion(version string) {
	fmt.Println("version:", version)
}
