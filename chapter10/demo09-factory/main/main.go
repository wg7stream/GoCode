package main

import(
	"fmt"
	"go_project/src/GoCode/chapter10/demo09-factory/model"
)

func main() {
	var stu model.NewStudent("tom~", 98.8)

	fmt.Println(*stu)
	fmt.Println("name = ", stu.Name, " score = ", stu.GetScore())
}