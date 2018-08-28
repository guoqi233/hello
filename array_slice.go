package main

import "fmt"

func main()  {
	var sl1 = make([]byte, 5, 5)
	//var sl2 = []byte{'a', 'b'}
	fmt.Println(cap(sl1[1:3]), sl1)
}
func fun1()  {
	var arr1 [6]int
	var slice1 = arr1[2:5]
	var slice2 = &arr1
	for index, value := range arr1{
		fmt.Printf("Slice at %d is %d\n", index, value)
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	//slice1 = slice1[:cap(slice1)]
	//fmt.Printf("The length of slice1 is %d\n", len(slice1))
	//fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
	slice1[0] = 1
	fmt.Println(arr1)
	fmt.Println(slice1)
	fmt.Printf("The capacity of slice2 is %d\n", cap(slice2))
}

func fun2()  {
	var slice1 = make([]int, 10)
	for index, _ := range slice1{
		slice1[index] = 5 * index
	}

	for index, value := range slice1{
		fmt.Printf("Slice at %d is %d\n", index, value)
	}
	fmt.Printf("\nThe length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}

//func Append(slice, data []byte)(res []byte) {
//	fmt.Println(cap(res))
//	lenSlice := len(slice)
//	capSlice := cap(slice)
//
//	lenData :=len(data)
//
//	lenTotal := lenSlice + lenData
//	fmt.Println(lenTotal)
//	if lenTotal <= capSlice {
//		res = append(slice, data...)
//	}else {
//		res = make([]byte, lenTotal, lenTotal)
//		fmt.Println(len(res))
//		res = append(slice, data...)
//	}
//	fmt.Println(cap(res))
//	return res
//}