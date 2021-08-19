package main

import "fmt"

func main() {
	fmt.Println("打开文件A准备读取。。。")
	defer fmt.Println("关闭文件A")			// 延时执行
	fmt.Println("打开文件B准备读取~~~")
	defer fmt.Println("关闭文件B")
	fmt.Println("进行文件的拷贝等操作@@@")
}