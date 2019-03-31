package main

import (
	"fmt"
	"strings"
)

func main() {
	//var str1 = "hello"
	//var str2 = "world!"
	//var a = 0x10
	//fmt.Println("\na")
	//fmt.Println(`\n`)
	//fmt.Println(str1 + " " + str2)
	//strings.HasPrefix(str1, "hel") //  判断字符串 s 是否以 prefix 开头：
	//strings.HasSuffix(str2, "ld!") //  判断字符串 s 是否以 suffix 结尾：
	//fmt.Println(strings.ToLower("Hello World!"))
	//fmt.Println(strings.TrimSpace("  123  "))
	//fmt.Println(strconv.Itoa(a))
	str := "adgfdg判断dhfg"
	//fmt.Println(string(str[0]))
	//for index, runeValue := range str{
	//	if !IsAscii(runeValue){
	//		fmt.Println(index, string(runeValue))
	//	}
	//}
	fmt.Println(strings.Map(replace(0), str))
}

func IsAscii(c rune) rune {
	if c > 255 {
		return 0
	}
	return c
}

func replace(c rune) func(ch rune) rune {
	return func(ch rune) rune {
		if ch > 255 {
			return c
		} else {
			return ch
		}
	}
}
