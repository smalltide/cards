package main

import (
	"fmt"
	"io"
	"os"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: printfile <filename>")
		os.Exit(1)
	}
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
