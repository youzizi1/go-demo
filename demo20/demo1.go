package main

import (
	"fmt"
)

func foo(cb func())  {
	fmt.Println("foo执行中")
	cb()
}

func main() {
	fmt.Println("start")
	foo(func ()  {
		fmt.Println("回调函数执行中")
	})
	fmt.Println("end")
}