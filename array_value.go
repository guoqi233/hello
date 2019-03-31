package main

import "fmt"

func main() {
	var arr1 = []int{0}
	var arr2 = arr1

	arr1[0] = 5

	fmt.Println(arr1)
	fmt.Println(arr2)
	change(arr2)
	fmt.Println(arr2)
	change_(arr2[:2])

}

func change_(s []int) {
	fmt.Println(s)
}

func change(arr []int) {
	arr[0] = 5
}

// 当数组赋值时，发生了数组内存拷贝。
