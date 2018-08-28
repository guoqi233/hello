package main

import "fmt"

func main(){
	//var a = 10
	//var p  = &a
	//
	//fmt.Println(a)
	//fmt.Println(*p)
	//a = 20
	//fmt.Println(*p)

	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
}
