package main
import (
	"fmt"
)

func main() {
	// 演示for-range
	heroes := [...]string{"松江", "无用", "卢俊义"}

	for i, v := range heroes {
		fmt.Printf("i = %v  v = %v\n", i, v)
		fmt.Printf("heroes[%d] = %v\n", i, heroes[i])
	}

	for _, v := range heroes {
		fmt.Printf("元素的值 = %v\n", v)
	}
}