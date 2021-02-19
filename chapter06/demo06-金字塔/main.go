package main
import(
	"fmt"
)

// 将打印金字塔的代码封装到函数中
func printPyramid(totalLevel int) {
	// i表示层数
	for i := 1; i <= totalLevel; i ++ {
		// 在打印前先打空格
		for k:= 1; k <= totalLevel - i; k ++ {
			fmt.Print(" ")
		}

		// j表示每层打印多少
		for j := 1; j <= 2 * i - 1; j ++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Println("请输入金字塔层数：")
	fmt.Scanln(&n)
	printPyramid(n)
}