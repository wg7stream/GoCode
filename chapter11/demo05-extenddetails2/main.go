package main

import (
	"fmt"
)

type A struct {
	Name string
	age int
}
type B struct {
	Name string
	Score float64
}
type C struct {
	A
	B
}
type D struct {
	a A// 有名结构体
}

type Goods struct {
	Name string
	Price float64
}

type Brand struct {
	Name string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

type Monster struct {
	Name string
	Age int
}

type E struct {
	Monster
	int
	n int
}

func main() {
	var c C
	c.A.Name = "tom"
	fmt.Println("c")

	var d D
	d.a.Name = "jack"

	tv := TV{Goods{"电视机001", 5000.99}, Brand{"海尔", "山东"},}
	fmt.Println(tv.Goods.Name)
	fmt.Println(tv.Price)

	tv2 := TV{
		Goods{
			Price : 500.99,
			Name : "电视机",
		},
		Brand{
			Name : "夏普",
			Address : "北京",
		},
	}

	fmt.Println("tv", tv)
	fmt.Println("tv2", tv2)

	tv3 := TV2{&Goods{"电视机",7000.99}, &Brand{"床位", "河南"},}

	tv4 := TV2{
		&Goods{
			Name : "电视机",
			Price : 9000.99,
		},
		&Brand{
			Name : "长虹",
			Address : "四川",
		},
	}

	fmt.Println("tv3", *tv3.Goods, *tv3.Brand)
	fmt.Println("tv4", *tv4.Goods, *tv4.Brand)
}

