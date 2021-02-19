package main
import (
	"fmt"
)

type Student struct {
	name string 
	gender string
	age int
	id int
	score float64
}

func (student *Student) say() string {
	infoStr := fmt.Sprintf("student的信息 name=[%v] gender=[%v], age=[%v] id=[%v] score=[%v]",
	student.name, student.gender, student.age, student.id, student.score)

	return infoStr
}

type Box struct {
	len float64
	width float64
	height float64
}

func (box *Box) getVolumn() float64 {
	return box.len * box.width * box.height
}

type Visitor struct {
	Name string
	Age int
}

func (visitor *Visitor) showPrice() {
	if visitor.Age >= 90 || visitor.Age <= 8 {
		fmt.Println("考虑到ddd")
		return
	}
}

func main() {
	
}