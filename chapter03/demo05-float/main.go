package main
import "fmt"

func main() {

	// 尾数部分可能丢失，造成精度损失
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3=", num3, "num4=", num4)


	var num5 = 1.1
	fmt.Printf("num5的数据类型 %T \n", num5)


	// 科学计数法
	num8 := 5.1234e2
	//num9 := 5.1234E2
	//num10 := 5.1234e-2

	fmt.Println("num8 = ", num8)


}