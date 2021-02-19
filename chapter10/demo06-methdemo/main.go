package main

import (
	"fmt"
)

type Person struct{
	Name string
}

// 给Person结构体添加speak方法
func (p Person) speak() {
	fmt.Println(p.Nmae)
}

func (p Person) jisuan() {
	res := 0
	for i := 0; i <= 1000; i ++ {
		res += i
	}
	fmt.Println(p.Name)
}


func (p Person) getSum() {
	return n1 + n2
}

func main() {
	var p Person
	p.Name = "safdsdf"
	p.speak();
}