package main

import(
	"fmt"
	"GOCODE/chapter11/demo02-capsulate/model"
)

func main() {
	p := model.NewPerson("smith")
	p.SetAge(18)
	fmt.Println(p)
}