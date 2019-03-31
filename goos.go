package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	var goos string = runtime.GOOS
	fmt.Println("The operating system is: %s", goos)
	path := os.Getenv("PATH")
	fmt.Println("Path is %s", path)
}
