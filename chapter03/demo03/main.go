package main
import "fmt"

// 定义全局变量
var n1 = 100
var n2 = 200
var name = "zcy"
// 上面的声明方式，也可以改成一次性声明
var (
	n3 = 300
	n4 = 900
	name2 = "mary"
)


func main() {
	// golong一次声明多个变量

	// 方式1：
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	// 方式2：
	var n4, name1, n5 = 100, "zhangchengyu", 789
	fmt.Println("n4=", n4, "name1=", name1, "n5=", n5)

	// 方式3：使用类型推导
	n6, name2, n7 := 123, "gjj", 456
	fmt.Println("n6=", n6, "name2=", name2, "n7=", n7)
}