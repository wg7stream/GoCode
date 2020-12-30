package main
import "fmt"

func main() {
	// while的实现
	var i int = 1
	for {
		if i > 10 {
			break
		}
		fmt.Println("hello world", i)
		i ++
	}

}