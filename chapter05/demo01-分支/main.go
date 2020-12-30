package main
import(
	"fmt"
)

func main() {
	var n1 int32 = 10
	var n2 int32 = 50

	if n1 + n2 >= 50 {
		fmt.Println("hello world")
	} else {
		fmt.Println("heiheihei")
	}

	var score int
	fmt.Scanln(&score)

	var b bool = true
	// 错误写法
	if b = false {
		fmt.Println("wdwdwd")
	}
 
}

