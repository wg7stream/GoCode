package main

import (
	"fmt"
	"GOCODE/chapter11/demo03-encapexercise/model"
)


func main() {
	account := model.NewAccount("zhangcy", "00ddd0", 40)
	if account != nil {
		fmt.Println("创建成功")
	} else {
		fmt.Println("创建失败")
	}
}
