package main

import "fmt"

func main(){
	half("hello,world!")
}

func half(str string){
	fmt.Println(str)
	length := len(str)
	fmt.Println(str[length/2:] + str[:length/2])
}