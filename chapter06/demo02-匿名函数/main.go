package main
import (
	"fmt"
)


var (
	// 全局匿名函数
	Func = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	// 匿名函数
	// 在定义匿名函数时就直接调用，这种方式匿名函数只能调用一次

	// 案例演示，求两个数的和  使用匿名函数的方式完成
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println("res1 = ", res1)

	// 将匿名函数func赋值给某个变量
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}

	res2 := a(10, 20)
	fmt.Println("res2 = ", res2)

	// 全局匿名函数
	res4 := Func(4, 9)
	fmt.Println("res4 = ", res4)
}