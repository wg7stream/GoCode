package main
import (
	"fmt"
	"errors"
)

func test() {
	// 使用defer + recover来捕获和处理异常
	defer func() {
		err := recover() // recover()内置函数，可以捕获到异常
		if err != nil { // 说明捕获到错误
			fmt.Println("err = ", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res = ", res)
}

// 自定义错误
// 函数去读取以配置文件init.conf的信息、
// 如果文件名传入不正确，我们就返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		// 读取..
		return nil
	} else {
		// 返回一个自定义错误
		return errors.New("读取文件错误..")
	}
}

func test02() {
	err := readConf("config2.ini")
	if err != nil {
		// 如果读取发生错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Println("test02()继续执行....")
}


func main() {
	// test()
	// fmt.Println("main()下面的代码...")
	test02()
	fmt.Println("main()下面的代码...")
}