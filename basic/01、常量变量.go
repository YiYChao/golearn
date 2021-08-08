package main

import "fmt"

// 一次性声明多个常量
const(
	Speed = 320.0
	Distance int = 120
)
// 一次性声明多个变量
var(
	method = "Train"
	passenger = 1000
)


func main0101() {
	const pi float32 = 3.14			// 定义一个圆周率常量
	var radius float32 = 10			// 定义一个可以变换的半径变量

	var area = pi * radius * radius // 计算圆的面积的表达式

	fmt.Println("面积是：", area)

	// 动态判断数据的类型
	const MoneyRate = 7
	var deposit = 123.45
	var flag= false
	fmt.Println(deposit, flag)
}

func main01_02() {
	price := 12.7		// 变量的声明和赋值合二为一
	fmt.Printf("type=%T, value=%v", price, price)	// type=float64, value=12.7
}

func main() {
	a := 4i + 3		// 定义一个复数
	b := 3i + 4
	fmt.Println(a * b)
}