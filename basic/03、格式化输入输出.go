package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b string			// 定义两个变量用于接受用户的输入
	fmt.Print("请输入两个数：")
	fmt.Scan(&a, &b)
	fmt.Println("字符串 + : ", a + b)
	aInt, _ := strconv.ParseInt(a, 0, 32)
	bInt, _ := strconv.ParseInt(b, 0, 32)
	fmt.Println("Int32 + : ", aInt + bInt)
}