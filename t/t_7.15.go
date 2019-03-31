package main

import "fmt"

func main() {
	fmt.Println(string(copyDiff("aabbcccd")))
}

func copyDiff(str string) (res []byte) {
	length := len(str)

	slice := make([]byte, length)
	copy(slice, str)

	for i := 1; i < length; i++ {
		if slice[i] != slice[i-1] {
			res = append(res, slice[i])
		}
	}
	return
}
