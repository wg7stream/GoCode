package main
import (
	"fmt"
	"encoding/json"
)

type A struct {
	Num int
}

type B struct {
	Num int
}

type Monster struct {
	Name string `json:"name"`// `json:"name"` 就是struct tag
	Age int `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	var a A
	var b B
	a = A(b)// 可以进行转换
	fmt.Println(a, b)

	monster := Monster{"牛魔王", 500, "芭蕉扇~"}

	// 将monster变量序列化为json格式字串
	// json.Marshl 函数中使用反射
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json 处理错误 ", err)
	}
	fmt.Println("jsonStr", string(jsonStr))
}