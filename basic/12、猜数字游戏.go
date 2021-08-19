package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := time.Now().Unix()
	myRand := rand.New(rand.NewSource(seed))
	score := myRand.Intn(100)
	fmt.Println(score)
	var inputN, count int
	for true {
		fmt.Print("请输入一个数：")
		fmt.Scan(&inputN)
		count ++		// 猜测次数，计数用
		if inputN > score {
			fmt.Println("大了》》》》》")
		} else if inputN < score {
			fmt.Println("小了《《《《《")
		} else {
			fmt.Println("恭喜你答对了！！！")
			break
		}
	}
}
