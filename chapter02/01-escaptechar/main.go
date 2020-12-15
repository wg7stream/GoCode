package main

import "fmt" //fmt中提供格式化，输出，输入函数 
func main(){
	//演示转义字符的使用 
	fmt.Println("tom\tjack")
	fmt.Println("hello\nworld")
	fmt.Println("c:\\user\\asdsad\\desktop")
	fmt.Println("tom说\"i love you\"")
	// \r 回车  表示从当前行的最前面开始输出，覆盖掉以前的内容
	fmt.Println("天龙八部雪山飞狐\r张飞")
}

// go run main.go
// go build main.go 生成可执行文件

