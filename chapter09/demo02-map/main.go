package main
import (
	"fmt"
)

func main() {
	cities := make(map[string]string)
	cities["asdf"] = "asdas"
	cities["sdfv"] = "fdvb"

	// 删除
	delete(cities, "asdf")

	// 一次删除所有的key
	// 直接make一个新空间  或者遍历
	cities = make(map[string]string)

	// 查找
	val, ok := cities("eee")
	if ok {
		fmt.Println("存在")
	} else {
		fmt.Println("没有")
	}
}