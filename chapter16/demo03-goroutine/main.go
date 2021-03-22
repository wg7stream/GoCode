package main

import (
	"fmt"
	_ "time"
	"sync"
)

// 实现计算1~200的各个数的阶乘 并将结果放入map


var (
	myMap = make(map[int]int, 10)
	// 声明一个全局互斥锁
	// lock 是一个全局的互斥锁
	// sync 是包：synchornized 同步
	// mutex：互斥
	lock sync.Mutex
)

func test(n int) {
	// 计算n的阶乘
	res := 1
	for i := 1; i <= n; i ++ {
		res *= i
	}

	// 将结果放入map
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {
	// 开启多个协程
	for i := 1; i <= 20; i ++ {
		go test(i)
	}

	lock.Lock()
	for i, v := range myMap {
		fmt.Println("map[%d] = %d\n", i, v)
	}
	lock.Unlock()
}