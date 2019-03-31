package main

import "fmt"

func main() {
	var arr = []int{3, 9, 6, 8, 3, 4, 7, 5, 1}
	bubble(arr)
	fmt.Println(arr)
}

func bubble(a []int) {
	length := len(a)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
	}
}
