package main

import "fmt"

func partition(a []int, start int, end int) (i int) {
	pivot := a[start]
	index := start + 1
	for i = index; i <= end; i++ {
		if a[i] < pivot {
			a[i], a[index] = a[index], a[i]
			index += 1
		}
	}
	index -= 1
	a[start], a[index] = a[index], a[start]
	i = index
	return
}

func main() {
	array := []int{4, 3, 9, 0, 2, 1, 9, 5, 8, 9, 5, 6, 5, 3, 1, 2}
	index := partition(array, 0, len(array)-1)
	fmt.Println(index, array[index])
	fmt.Println(array)
}
