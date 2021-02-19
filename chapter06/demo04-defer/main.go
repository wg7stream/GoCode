package main
import (
	"fmt"
)

func sum(n1 int, n2 int) int{
	// 当执行到defer时，暂时不会执行，会将defer后面的语句压入到独立的栈
	// 当函数执行完毕后，再从defer栈中按照先入后出的方式出栈、
	defer fmt.Println("ok1 n1 = ", n1)
	defer fmt.Println("ok2 n2 = ", n2)

	n1 ++
	n2 ++

	res := n1 + n2
	fmt.Println("ok3 = ", res)
	return res
}

func main(){
	res :=  sum(10, 20)
	fmt.Println("res := ", res)
}