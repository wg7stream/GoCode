package main
import (
	"fmt"
)

// 函数外部声明/定义的变量叫全局变量
// 作用域在整个包都有效，如果其首字母为大写，则作用域在整个程序有效、
var age int = 50
var Name string = "jack~"

// 函数
func test() {
	// age和Name的作用域只在test函数内部
	age ：= 10
	Name := "fdsfsd"
	fmt.Println(age)
	fmt.Println(Name)
}

func main {
	fmt.Println(age)
	fmt.Println(Name)
	test()
}