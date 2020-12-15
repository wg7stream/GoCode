package main
// import "fmt"
// import "unsafe"

import (
	"fmt"
	"unsafe"
)

func main() {
	// int8的范围 -128~127
	var j int8 = 127
	fmt.Println("j=", j);


	var n1 = 100
	// 查看某个变量的数据类型
	// fmt.Printf() 可以用于格式化输出
	fmt.Printf("n1 的类型 %T\n", n1)

	var n2 int64 = 10
	// 查看某个变量的占用字节大小和数据类型
	// unsafe包的一个函数Sizeof()
	fmt.Printf("n2 的类型 %T n2占用的字节数是 %d ", n2, unsafe.Sizeof(n2))
}