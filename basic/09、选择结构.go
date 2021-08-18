package main

import "fmt"

func main0901() {
	/*单分支结构*/
	var birthday string
	fmt.Print("请输入你的生日(如：0101):")
	fmt.Scan(&birthday)
	if birthday >= "0601" && birthday <= "1030"{
		fmt.Println("初级：阁下真乃幸运儿~~")
	}
	fmt.Println("======= 初级诊断结束 =======")
	/*双分支结构*/
	if birthday >= "0601" && birthday <= "1030"{
		fmt.Println("中级：阁下真乃幸运儿~~")
	} else {
		fmt.Println("中级：阁下需再接再厉~~")
	}
	fmt.Println("======= 中级诊断结束 =======")
	/*多分支结构*/
	if birthday >= "0601" && birthday <= "1030"{
		fmt.Println("高级：阁下真乃幸运儿~~")
	} else if birthday >= "0301" && birthday <= "0531" {
		fmt.Println("高级：阁下运气尚存~~")
	} else {
		fmt.Println("高级：阁下运气不足~~")
	}
	fmt.Println("======= 高级诊断结束 =======")
}

func main() {
	/*switch结构*/
	var month uint8
	fmt.Print("请输入出生的月份：")
	fmt.Scan(&month)
	/*单点判断*/
	switch month {
	case 12,1,2:
		fmt.Println("单点判断：寒冬腊月")
	case 3, 4, 5:
		fmt.Println("单点判断：春暖花开")
	case 6, 7, 8 :
		fmt.Println("单点判断：夏日炎炎")
		fallthrough		// switch穿透,只从本层穿透到临近的下一层
	case 9, 10, 11:
		fmt.Println("单点判断：秋高气爽")
	default:
		fmt.Println("单点判断：世外桃源~~")
	}
	/*自由条件判断*/
	switch {
	case month == 12 || month >= 1 && month <= 2:
		fmt.Println("自由条件：寒冬腊月")
	case 3 <= month && month <= 5:
		fmt.Println("自由条件：春暖花开")
	case 6 <= month && month <= 8 :
		fmt.Println("自由条件：夏日炎炎")
		fallthrough
	case 9 <= month && month <= 11:
		fmt.Println("自由条件：秋高气爽")
	default:
		fmt.Println("自由条件：世外桃源~~")
	}
}
