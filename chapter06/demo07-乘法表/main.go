package main
import (
	"fmt"
)

// 编写一个函数调用九九乘法表
func printMulti(num int) {
	// 打印出九九乘法表
	// i表示层数
	for i := 1; i <= num; i ++ {
		for j := 1; j <= i; j ++ {
			fmt.Printf("%v * %v = %v \t", j, i, j * i)
		}
		fmt.Println()
	}
}

func main() {
	var num int
	fmt.Println("输入数字：")
	fmt.Scanln(&num)
	printMulti(num)
}