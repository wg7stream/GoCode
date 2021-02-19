package main
import(
	"fmt"
)

func main() {
	var intArr [3]int// int占8字节
	// 数组各个元素默认为0
	fmt.Println(intArr)// 打印数组的所有元素

	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30

	fmt.Println(intArr)
	fmt.Printf("intArr的地址 = %p intArr[0] 地址%p intArr[1] 底子%p intArr[2] 地址%p",
		&intArr, &intArr[0], &intArr[1], &intArr[2])


	// 数组的4种初始化方式
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Printf("numArr01 = ", numArr01)

	var numArr02 = [3]int{5, 6, 7}
	fmt.Println("numArr02 = ", numArr02)

	var numArr03 = [...]int{8, 9, 10}
	fmt.Println("numArr03 = ", numArr03)

	var numArr04 = [...]int{1 : 800, 0 : 900, 2 : 999}
	fmt.Println("numArr04 = ", numArr04)

	// 类型推导
	strArr05 := [...]string{1: "tom", 0: "jack", 2: "mary"}
	fmt.Println("strArr05 = ", strArr05)
}