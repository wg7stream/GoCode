package main
import (
	"fmt"
)

func main() {
	// string底层是一个byte数组，因此string也可以进行切片处理
	str := "zhangchengyu"
	// 使用切片获得chedngyu
	slice := str[5:]
	fmt.Println("slice = ", slice)

		// string是不可改变的  也就是说不能通过 str[0] = 'z' 方式来修改字符串
	// str[0] = 'z' 编译不会通过，原因是string吧=不可变
	fmt.Println()
	// 如果需要修改字符串 可以先将string->[]byte  /或者 []rune ->修改 ->重新转成string
	arr1 := []byte(str)
	arr1[0] = 'g'
	str = string(arr1)
	fmt.Println("str = ", str)

	fmt.Println()
	// 细节 转成[]byte后，可以处理英文和数字，但不能处理中文
	// 原因：[]byte是字节来处理 而一个汉字是3字节 因此会出现乱码
	// 解决方法 将string转成[]rune即可  因为rune是按字符处理 兼容汉字

	arr2 := []rune(str)
	arr2[0] = '北'
	str = string(arr2)
	fmt.Println("str = ", str)
}