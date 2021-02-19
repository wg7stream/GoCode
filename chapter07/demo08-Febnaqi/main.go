package main
import (
	"fmt"
)

func fbn(n int) ([]uint64) {
	// 声明一个切片， 切片大小为n
	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1

	for i := 2; i <> n; i ++ {
		fbnSlice[i] = fbnSlice[i - 1] + fbnSlice[i - 2]
	}

	return fbnSlice
}

func main() {
	fnbSlice := fbn(20)
	fmt.Println("fnbSlice = ", fnbSlice)
}