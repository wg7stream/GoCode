package main
import(
	"fmt"
	"strconv"
)

func main() {
	// 1. 统计字符串长度，按字节 len(str)
	// golong的统一编码为utf-8（ascii的字符（字母和数字）占一个字节 汉子占3个字节）
	str := "hello暗"
	fmt.Println("str len = ", len(str))// 8

	str2 := "hello北京"
	// 2. 字符串遍历 同时处理有中文的问题 r := []rune(str)
	r := []rune(str2)
	for i := 0; i < len(r); i ++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	// 3. 字符串转整数： n, err := strconv.Atoi("12")
	n, err := strconv.Atoi("hello")
	if err != nil {
		fmt.Println("转换错误", err)
	}else {
		fmt.Println("转换的结果是", n)
	}

	// 4. 整数转字符串
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v, str=%T\n", str, str)

	// 5. 字符串转 []byte: var bytes = []byte("hello go")
	var bytes = []byte("hello go")
	fmt.Printf("bytes = %v\n", bytes)

	// 6. []byte 转 字符串： str = string([]byte{97, 98, 99})
	str = string([]byte{97, 98, 99})
	fmt.Printf("str = %v\n", str)

	// 7. 十进制转 二、八、十六进制： str = strconv.FormatInt(123, 2), 返回对应的字符串
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制为 = %v\n", str)
} 
