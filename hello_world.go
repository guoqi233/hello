package main

import "fmt"

var global int

func main() {
	fmt.Println("hello, golang")
	fmt.Println(global)
	fmt.Println(FunctionName())
	fmt.Println(global)
}

func FunctionName() (a int, b int) {
	var global = 3
	a = 10
	b = 20
	fmt.Println(global)
	return a, b
}
