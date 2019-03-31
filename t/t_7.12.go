package main

import "fmt"

func main() {
	fmt.Println(split("hello, world!", 15))
}

func split(str string, i int) (str1, str2 string) {
	if i > len(str) {
		return
	}
	slice := make([]byte, len(str))
	copy(slice, str)
	slice1, slice2 := slice[:i], slice[i:]
	str1 = string(slice1)
	str2 = string(slice2)
	return
}
