package main

import (
	"fmt"
)

func main() {
	// 基本数据类型
	//var flag = true / false		// true对应1，false对应0
	var va complex64 = 3i + 4			// 定义一个复数
	var vb complex64 = 5i + 6
	fmt.Println("复数相加的结果：", va + vb)

	var a uint8 = 'a'
	var b byte = 'b'
	fmt.Printf("a=%v, a=%d, a=%c, b=%v, b=%c", a, a, a, b, b)
}