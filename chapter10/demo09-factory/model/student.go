package main

// 定义一个结构体
type student Struct{
	Name string
	score float64
}

// student结构体的首字母为小写 因此只能在model中使用
// 通过工厂模式来解决该问题
func  NewStudent(n string, s float64) *student {
	return &student{
		Name : n,
		score : s,
	}
}

// 前面写参数是给该结构体绑定一个参数
func (s *student) GetScore() float64 {
	return s.score
}