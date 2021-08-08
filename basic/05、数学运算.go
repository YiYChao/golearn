package main

import (
	"fmt"
	"math"
)

func main() {
	/*加、减、乘、除、取余*/
	fmt.Println("5 + 3 = ", 5 + 3)		// 8
	fmt.Println("5 - 3 = ", 5 - 3)		// 2
	fmt.Println("5 * 3 = ", 5 * 3)		// 15
	fmt.Println("5 / 3 = ", 5 / 3)		// 1
	fmt.Println("5 % 3 = ", 5 % 3)		// 2
	/*乘方、开方*/
	fmt.Println("5的3次方 = ", math.Pow(5, 3))					// 125
	fmt.Println("125开3次方 = ", math.Pow(125.0, 1.0 / 3.0))	// 5.000000000000001
	/*四舍五入*/
	fmt.Println("3.49四舍五入为：", math.Round(3.49))		// 3
	fmt.Println("3.51四舍五入为：", math.Round(3.51))		// 4
	fmt.Println("-3.49四舍五入为：", math.Round(-3.49))	// -3
	fmt.Println("-3.51四舍五入为：", math.Round(-3.51))	// -4
	fmt.Println("3.51舍入为：", math.Ceil(3.51))			// 4
	fmt.Println("3.51舍出为：", math.Floor(3.51))			// 3
	/*绝对值*/
	fmt.Println("3.49的绝对值为：", math.Round(3.49))		// 3
	fmt.Println("-3.49的绝对值为：", math.Round(-3.49))	// -3
	/*三角函数，参数必须是弧度而不是角度*/
	fmt.Println("30°的正弦为：", math.Sin((30.0 / 180) * math.Pi))		// 0.5
	fmt.Println("30°的余弦为：", math.Cos((30.0 / 180) * math.Pi))		// 0.8660254037844386
	fmt.Println("30°的正切为：", math.Tan((30.0 / 180) * math.Pi))		// 0.5773502691896258
	fmt.Println("30°的余弦为：", 1.0 / math.Tan((30.0 / 180) * math.Pi))	// 1.732050807568877
	/*反三角函数*/
	fmt.Println("正弦为0.5角度是：", math.Asin(0.5) * 180 / math.Pi)	// 30.000000000000004
}