package main
import (
	"fmt"
)

func main() {
	// slice表示切片(分片)，例如对一个数组进行切片，取出数组中的一部分值。
	// 每一个slice结构都由3部分组成：容量(capacity)、长度(length)和指向底层数组某元素的指针
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	// 声明/定义一个切片
	// slice := intArr[1:3]
	// 1. slice为切片名
	// 2. intArr[1:3]表示slice引用到intArr这个数组
	// 3. yinyogn intArr数组的起始下标为1，最后下标为3（左开右闭）
	slice := intArr[1:3]
	fmt.Println("intArr = ", intArr)
	fmt.Println("slice的元素是 = ", slice)
	fmt.Println("slice的元素个数 =", len(slice))
	fmt.Println("slice的容量 = ", cap(slice))



}