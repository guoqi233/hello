package main

import (
	"fmt"
	"strings"
)

func main() {
	//for i :=0; i<10 ; i++  {
	//	fmt.Printf("The num is %d\n", i)
	//}
	//
	//i := 0
	//for i <25{
	//	i += 1
	//	fmt.Println(strings.Repeat("G", i))
	//}
	//
	//for i := 0; i <= 10; i++ {
	//	fmt.Printf("the complement of %b is: %b\n", i, ^i)
	//}

	//const (
	//	FIZZ		= 3
	//	BUZZ		= 5
	//	FIZZBUZZ	= 15
	//)
	//for i := 1; i < 100; i++{
	//	msg := ""
	//	switch  {
	//	case i % FIZZBUZZ == 0:
	//		msg = "FIZZBUZZ"
	//	case i % FIZZ == 0:
	//		msg = "Fizz"
	//	case i % BUZZ == 0:
	//		msg = "Buzz"
	//	default:
	//		msg = strconv.Itoa(i)
	//	}
	//	fmt.Println(msg)
	//}

	w, h := 20, 10
	temp := strings.Repeat("*", w) + "\n"
	content := strings.Repeat(temp, h)
	fmt.Println(content)

	a, b := 6, 3
	fmt.Printf("a: %d; b: %d\n", a, b)
	Switch(&a, &b)
	fmt.Printf("a: %d; b: %d\n", a, b)
	slice := []int{1, 2, 3, 4}
	What(1, slice)
}

func MaxMin(x int, y int) (max int, min int) {
	if x > y {
		max, min = x, y
	} else {
		max, min = y, x
	}
	return
}

func Switch(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func What(a int, li []int) {
	fmt.Println(a)
	fmt.Println(li)
}
