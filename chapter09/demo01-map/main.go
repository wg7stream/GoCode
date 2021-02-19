package main
import (
	"fmt"
)

func main() {

	// map的使用3种方式

	// 方式1
	var a map[string]string
	a = make(map[string]string, 10)
	a["no1"] = "asdfa"
	a["sdf"] = "sdsdfg"
	fmt.Println(a)

	// 方式2
	cities := make(map[string]string)
	cities["asdf"] = "asdas"
	cities["sdfv"] = "fdvb"
	fmt.Println(cities)

	// 方式3
	hores := map[string]string{
		"df" : "qwedf",
		"vv" : "sadfdsf",
	}
	hores["trgt"] = "dgg"
	fmt.Println(hores)

	// 案例演示
	studentMap := make(map[string]map[string]string)

	studentMap["stu1"] = make(map[string]string, 3)
	studentMap["stu1"]["name"] = "tom"
	studentMap["stu1"]["sex"] = "男"
	studentMap["stu1"]["address"] = "baoji"
}