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

	// print "Hello World"
	fmt.Println("Hello World")

	// finish(status code 0)
	os.Exit(0)
}
