package main
import (
	"fmt"
)

type Cat struct {
	Name string
	Age int
	Color string
	Hobby string
}

func main() {
	var cat1 Cat
	cat1.Name = "sdf"

	fmt.Println(cat1)
}