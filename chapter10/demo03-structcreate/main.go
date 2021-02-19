package main
import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func main() {
	p2 := Person{"mary", 20}
	fmt.Println(p2)

	var p3 *Person = new(Person)
	// (*p3).Name = "zcy" 也可以写为 p3.Name = "zcy"

	(*p3).Name = "zcy"
	p3.Name = "sdf"

	var person *Person = &Person{}
	(*person).Nmae = "sdfv"
	person.Nmae = "sdff"


}