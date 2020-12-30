package main

import(
	"fmt"
)
 
func main() {
	for i := 1; i <= 10; i ++ {
		fmt.Println("swafdasd")
	}
	
	j := 1
	for j <= 10 {
		fmt.Println("swafdasd")
		j ++
	}

	k := 1
	for {
		if k < 10 {
			fmt.Println("ewfwef")
		} else {
			break
		}
		k ++
	}

	// 字符串遍历
	var str string = "helloworld"

	for i := 0; i < len(str); i ++ {
		fmt.Printf("%c \n", str[i])
	}

	for index, val := range str {
		fmt.Printf("index = %d, val = %c \n", index, val)
	}
}
