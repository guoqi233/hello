package main

import "fmt"

func main() {
	fmt.Println(reverse("hello, world!"))
}

func reverse(str string) (res string) {
	slice := make([]byte, len(str))
	copy(slice, str)
	length := len(str)
	for i := 0; i < length/2; i++ {
		temp := slice[i]
		slice[i] = slice[length-1-i]
		slice[length-1-i] = temp
	}
	res = string(slice)
	return
}
