package main

import "fmt"

func main(){
	var a = []int{1,2,3,4,5}
	res := mapFunc(multiplyFunc(9), a)
	fmt.Println(a)
	fmt.Println(res)
}

func multiplyFunc(a int)(func(num int)(res int)){
	return func(num int) (res int) {
		res = num*a
		return
	}
}

func mapFunc(function func(num int)(res int), a []int)(res []int){
	res = make([]int, len(a))
	for index, value := range a{
		res[index] = function(value)
	}
	return
}
