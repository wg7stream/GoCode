package main

import (
	"runtime"
	"fmt"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("couNum = ", cpuNum)

	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}