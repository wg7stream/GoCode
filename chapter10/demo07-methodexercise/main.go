package main
import (
	"fmt"
)

type Circle struct {
	radius float64
}

// 返回面积的函数
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c *Circle) area2() float64 {
	fmt.Printf(c)
	c.eadius = 10;
	return 3.14 * c.radius * c.radius
}

func main() {
	var c Circle
	fmt.Printf(&c)
}