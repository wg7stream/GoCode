package main

import (
	"fmt"
	"time"
)

// 向intChan放入1-8000个数
func putNum(intChan chan int) {
	for i := 1; i <= 80000; i ++ {
		intChan<- i
	}
	
	// 关闭intChan
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan

		if !ok {
			break
		}
		flag = true// 先假设是素数
		for i := 2; i < num; i ++ {
			if num % i == 0 {// 说明不是素数
				flag = false
			}
		}

		if flag {
			// 将这个数放入到primeChan
			primeChan<- num
		}
	}

	fmt.Println("有一个primeNum 协程因为娶不到数据，退出")
	// 这里还不能关闭primeChan
	// 向exitChan写入true
	exitChan<- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 20000)
	// 标识退出的管道
	exitChan := make(chan bool, 8)// 4个

	start := time.Now().Unix()

	// 开启一个协程，向intChan放入1-8000个数
	go putNum(intChan)

	// 开启4个协程 从intChan取出数据，并判断是否为素数，如果是
	// 就放入到primeChan
	for i := 0; i < 8; i ++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	// 主线程进行处理
	go func() {
		for i := 0; i < 8; i ++ {
			<-exitChan
		}

		end := time.Now().Unix()
		fmt.Println("使用协程耗时=", end - start)

		close(primeChan)
	}()

	// 遍历primeChan，取出结果
	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
	}

	fmt.Println("main线程退出")
}