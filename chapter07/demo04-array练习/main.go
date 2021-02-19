package main
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 创建一个数组，存放'A'-'Z'
	var myChars [26]byte
	for i := 0; i < 26; i ++ {
		myChars[i] = 'A' + byte(i) 
	}

	for i := 0; i < 26; i ++ {
		fmt.Printf("%c", myChars[i])
	}

	// 输出数组中最大的数
	fmt.Println()// 换行
	var intArr [6]int = [...]int {1, -1, 9, 90, 11, 9000}
	maxVal := intArr[0]
	maxValIndex := 0

	for i := 0; i < len(intArr); i ++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValIndex = i
		}
	}
	fmt.Printf("maxVal = %v  maxValIndex = %v", maxVal, maxValIndex)

	// 求平均值
	var intArr2 [5]int = [...]int {1, -1, 9, 90, 12}
	sum := 0
	for _, val := range intArr2 {
		sum += val
	}

	fmt.Printf("sum = %v  平均值 = %v", sum, float64(sum) / float64(len(intArr2)))


	// 随机生成5个数，然后反转打印
	var intArr3 [5]int 
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(intArr3); i ++ {
		intArr3[i] = rand.Intn(100)
	}
	fmt.Println(intArr3)
}