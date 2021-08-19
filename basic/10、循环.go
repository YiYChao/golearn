package main

import (
	"fmt"
	"time"
)

func main1001() {
	/*无限循环(死循环)*/
	// 不断输出今天的天气
	for {
		fmt.Println("今天居然在下雨！！！")
		time.Sleep(time.Millisecond * 500)		// 暂停500毫秒
	}
}

func main1002() {
	/*有限循环*/
	// 循环输出工作年限，到10为止
	for i := 0; i < 10; i++ {
		fmt.Printf("这份工作我已经干了将近【%d】年了！\n", i + 1)
	}
}

func main1003() {
	/*跳步循环*/
	// 输出1~10以内的偶数
	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
	}
}

func main1004() {
	/*跳步循环*/
	// 输出1~20以内能被3整除的数
	for i := 1; i <= 20; i++ {
		if i % 3 == 0 {
			fmt.Println(i)
		}
	}
}

func main1005() {
	/*输出9*9乘法表*/
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d \t", j, i, j*i)
		}
		fmt.Println()	// 输出一个换行
	}
}

func main1006() {
	/*goto*/
	// 在一堆数中查找一个符合条件的数(能被同时被3，5，11整除)
	for i := 1; i < 1000; i++ {
		if i % 3 == 0 && i % 5 == 0 && i % 11 == 0 {
			fmt.Println(i)
			goto DONE
		}
	}
	fmt.Println("不确定是否查找到伤心*^____^*")
	DONE:
		fmt.Println("终于找到你了o(*￣▽￣*)ブ")
}

func main1007() {
	/*continue语句的使用*/
	// 输出1~10之间的偶数
	for i := 1; i <= 10; i++ {
		if i % 2 > 0 {
			continue		// 结束本轮循环，继续下次循环
		}
		fmt.Println(i)
	}
}

func main() {
	/*break*/
	// 在一堆数中查找一个符合条件的数(能被同时被3，5，11整除)
	for i := 1; i < 1000; i++ {
		if i % 3 == 0 && i % 5 == 0 && i % 11 == 0 {
			fmt.Println(i)
			break			// 结束本层循环
		}
	}
	fmt.Println("终于找到你了o(*￣▽￣*)ブ")
}