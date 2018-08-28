package main

import "fmt"

func main()  {
	var month int
	month = 3
	switch {
	case month >0 && month<=3:
		fmt.Println("Spring")
	case month >3 && month<=6:
		fmt.Println("Summer")
	case month >6 && month<=9:
		fmt.Println("Autumn")
	case month >9 && month< 13:
		fmt.Println("Winter")
	default:
		fmt.Println("Bad month")
	}
}
