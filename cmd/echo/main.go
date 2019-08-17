package main

import (
	"fmt"
	"os"
)

func main() {

	// print command line arguments
	length := len(os.Args)
	for i := 1; i < length; i++ {
		fmt.Printf("%v ", os.Args[i])
	}

	// print line feed code
	fmt.Println()

	// finish(status code 0)
	os.Exit(0)
}
