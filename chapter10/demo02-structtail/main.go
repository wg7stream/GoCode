package main

import "fmt"

// 结构体
type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown, *Point
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}

	fmt.Printf("%p %p %p %p ", &r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	r2 := Rect2{*Poitn{10, 20}, &Point{30, 40}}

	fmt.Printf("%p %p \n", &r2.leftUp, &r2.rightDown)

	fmt.Printf("%p %p\n", r2.leftUp, r2.rightDown)
}

