package main
import (
	"fmt"
	"strings"
)

// 累加器
func AddUpper() func(int) int {
	// 闭包
	var n int = 10;
	var str = "hello"
	return func(x int) int {
		n = n + x;
		str += string(36)
		fmt.Println("str = ", str)
		return n
	}
}

// 闭包练习
func makeSuffix(suffix string) func(string) string {
	return func (name string) string{
		// 如果name没有指定后缀 则加上 否则就返回原来的名字
		if !strings.HasSuffix(name, suffix) {
			return name +suffix
		}

		return name 
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))

	f2 := makeSuffix(".jpg")
	fmt.Println("文件处名理后的 = ", f2("winter"))
}