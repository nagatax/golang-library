package main

import (
	"fmt"
	"os"
)

func main() {

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
