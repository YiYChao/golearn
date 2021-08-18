package main

import "fmt"

func main() {
	fmt.Println(7 & 5)		// 0111 & 0101 --> 0101 = 5
	fmt.Println(7 | 5)		// 0111 & 0101 --> 0111 = 7
	fmt.Println(7 ^ 5)		// 0111 & 0101 --> 0010 = 2
	fmt.Println(7 << 1)		// 0111  --> 1110 = 14
	fmt.Println(7 >> 1)		// 0111  --> 0011 = 3
	var a int8 = 3
	print(a << 7)
}
