package main
import (
	"fmt"
)

func main() {
	var arr3 = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	// for循环遍历
	for i := 0; i < len(arr3); i ++ {
		for j := 0; j < len(arr3[i]); j ++ {
			fmt.Println("%v\t", arr3[i][j])
		}
		fmt.Println()
	}

	// for range方式
	for i, v := range arr3 {
		for j, v2 := arnge v {
			fmt.Printf("arr3[%v][%v] = %v \t", i, j, v2)
		}
		fmt.Println()
	}
}