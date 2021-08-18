package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*类型的强制转换*/
	var a int = 123
	aFloat := float32(a)
	fmt.Printf("aFloat: Type=%T, value=%v\n", aFloat, aFloat)	// aFloat: Type=float32, value=123

	var b float64 = 456.789
	bInt16 := int16(b)
	fmt.Printf("bInt16: Type=%T, value=%v\n", bInt16, bInt16)	// bInt16: Type=int16, value=456

	strInt, _ := strconv.ParseInt("345", 10, 16)
	fmt.Printf("strInt: Type=%T, value=%v\n", strInt, strInt)	// strInt: Type=int64, value=345

	strFloat, _ := strconv.ParseFloat("7456.78", 32)
	fmt.Printf("strFloat: Type=%T, value=%v\n", strFloat, strFloat)	// strFloat: Type=float64, value=7456.77978515625
}
