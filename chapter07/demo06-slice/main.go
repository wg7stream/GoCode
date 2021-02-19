package main
import (
	"fmt"
)

func main() {
	// 切片的使用
	// 2. make
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[3] = 20

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// 3. 

	fmt.Println()
	var strSlice []string = []string{"das", "fsad", "sdscvd"}
	fmt.Println("strSlice = ", strSlice)

	fmt.Println(len(strSlice))
	fmt.Println(cap(strSlice))

	// 切片的遍历
	// 常规方式
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice1 := arr[1:4]
	for i := 0; i < len(slice1); i ++ {
		fmt.Printf("slice1[%v] = %v", i, slice1[i])
	}

	// for range方式
	fmt.Println()
	for i, v := range slice1 {
		fmt.Printf("i = %v v = %v \n", i, v)
	}


	fmt.Println()

	// append
	var slice3 []i
	
	nt = []int{100, 200, 300}
	slice3 = append(slice3, 400, 500, 600)
	fmt.Println("sloce3: ", slice3)

	// 通过append将切片追加到切片
	slice3 = append(slice3, slice3...)
	fmt.Println("slice3", slice3)

	// 拷贝操作
	fmt.Println()
	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10)
	copy(slice5, slice4)
	fmt.Println("slice4:", slice4)
	fmt.Println("slice5:", slice5)
}

