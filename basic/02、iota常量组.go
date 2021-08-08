package main

import "fmt"

const(
	Monday = iota		// 从0开始自动增长，规则为(iota, iota+1)
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)
const(
	D = 2 << iota		// 2，规则为(2左移一位，二位，三位……)
	C					// 4
	B 					// 6
	A					// 8
)

func main() {
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
	fmt.Println(D, C, B, A)
}