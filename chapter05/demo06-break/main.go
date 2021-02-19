package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 给rand设置种子
	rand.Seed(time.Now().UnixNano())
	// 生成0到100之间的一个整数
	n := rand.Intn(100) + 1// [0, 100)
	fmt.Println(n)

	var count int = 0
	for {
		rand.Seed(time.Now().Unix())
		n := rand.Intn(100) + 1
		count ++
		if (n == 99) {
			break// 跳出for循环
		}
	}

	fmt.Println(count)

	// 标签形式使用label
	label2:
	for i := 0; i < 4; i ++ {
		//label1:
		for j := 0; j < 10; j ++ {
			if j == 2 {
				// break 默认跳出最近的for循环
				break label2
			}
			fmt.Println(j)
		}
	}
}