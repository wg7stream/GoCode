package main

import (
	"fmt"
)

func main() {
	// 创建一个可以存放3个整数的管道
	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Println("intChan 的值 = %v intChan本身的地址 = %p\n", intChan, &intChan)


	// 向管道写入数据
	intChan<- 10
	num := 211
	intChan<- num
	intChan<- 50

	// 从channel中取出数据后 可以继续放入
	<-intChan
	intChan<- 98

	// 查看管道的长度和容量
	fmt.Printf("channel len = %v cap = %v \n", len(intChan), cap(intChan))

	// 从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2 = ", num2)
	fmt.Printf("channel len = %v cap = %v \n", len(intChan), cap(intChan))

	// 在没有协程的情况下 如果管道数据已经全部取出，再取就会报deadlock
	num3 := <-intChan
	num4 := <-intChan

	fmt.Println("num3 = ", num3, "num4 = ", num4)
}