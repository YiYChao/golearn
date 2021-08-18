package main

import "fmt"

// 比较运算 和 逻辑运算
func main() {
	/*比较运算*/
	fmt.Println(1 + 2 == 3)			// true
	fmt.Println(1 + 2 != 3)			// false
	fmt.Println(1 + 2 > 3)			// false
	fmt.Println(1 + 2 >= 3)			// true
	fmt.Println(1 + 2 < 3)			// false
	fmt.Println(1 + 2 <= 3)			// true
	/*逻辑运算*/
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)

	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)

	fmt.Println(!true)
	fmt.Println(!false)
}