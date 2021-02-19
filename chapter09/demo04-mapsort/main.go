package main
import (
	"fmt"
	"sort"
)

func main() {
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 899

	fmt.Println(map1)

	// 按照map的key的顺序进行排序
	// 1. 先将map的key放入到切片中
	// 2. 对切片排序
	// 3. 遍历切片，然后按照key来输出map的值
	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	fmt.Println(keys)

	for _, k := range keys {
		fmt.Printf("map1[%v] = %v \n", k, map1[k])
	}
}