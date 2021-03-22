package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age int
}

func main() {
	// 定义一个可以存放任意数据类型的管道
	allChan := make(chan interface{}, 3)

	allChan<- 10
	allChan<- "zhangchengyu"
	cat := Cat{"小花猫", 4}
	allChan<- cat


	<-allChan
	<-allChan

	newCat := <-allChan

	fmt.Printf("newCat = %T , newCat = %v\n", newCat, newCat)
	// 使用类型断言
	// go里面的类型断言写法：
	// x.(T)  检查x的动态类型是否是T，其中x必须是接口值  
	// 类型断言检查x的动态类型是否等于具体类型T。如果检查成功，类型断言返回的结果是x的动态值，其类型是T
	// 其中x为interface{}类型 T是要断言的类型。
	a := newCat.(Cat)
	fmt.Printf("newCat.Name = %v", a.Name)
}