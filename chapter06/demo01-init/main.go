package main
import (
	"fmt"
	// 引入包
	"GoCode/chapter06/demo01-init/utils"
)

var age = test()

// 全局变量是先被初始化 然后才是init函数
func test() int {
	fmt.Println("test()...") // 1
	return 90
}

// init函数，通常可以再init函数中完成初始化工作
func init() {
	fmt.Println("init()...") // 2
}


func main() {
	fmt.Println("main()...age = ", age) // 3
	fmt.Println("Age =", utils.Age, "Name = ", utils.Name)
}