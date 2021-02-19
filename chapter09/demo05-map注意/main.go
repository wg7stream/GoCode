package main
import {
	"fmt"
}

func modify(map1 map[int]int) {
	map1[10] = 900
}

type Stu struct {
	Name string
	Age int
	Address string
}

func main() {
	map1 := make(map[int]int, 2)
	map1[1] = 90
	map1[2] = 77
	map1[10] = 44
	map1[20] = 489
	modify(map1)

	fmt.Println()

	students := make(map[string]Stu, 10)

	stu1 := Stu{"tom", 18, "北京"}
	stu2 := Stu{"mary", 28, "上海"}
	students["no1"] = stu1
	students["no2"] = stu2
	fmt.Println(students)

}