package main

import (
	"fmt"
)

func main() {
	var arr1 *[5]byte
	for index, _ := range arr1 {
		fmt.Println(index)
		//fmt.Println(index, value, reflect.TypeOf(value))
	}
}
