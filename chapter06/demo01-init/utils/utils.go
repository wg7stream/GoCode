package utils
import "fmt"

var Age int
var Name string

// Age和Name全局变量，需要再main.go中使用
// 但我们需要初始化Age和Name

//init函数完成初始化工作
func init() {
	fmt.Println("utils包的init()...")
	Age = 100
	Name = "Tom~"
}