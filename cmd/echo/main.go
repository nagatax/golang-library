// Package main
package main

import (
	"fmt"
	"os"

	"github.com/nagatax/golang-library/utils"
)

const (
	// Version of software
	Version = "0.0.1"
)

var (
	revision string
)

func main() {
	// print version
	if utils.IsPrintVersion(os.Args[1:]...) {
		utils.PrintVersion(Version)
		return
	}

	// print command line arguments
	var (
		args         = ""
		length       = len(os.Args)
		startingArgs = 1
	)
	for i := startingArgs; i < length; i++ {
		args += os.Args[i]
	}
	fmt.Printf("%v ", args)

	// print line feed code
	fmt.Println()

	// finish(status code 0)
	os.Exit(0)
}
