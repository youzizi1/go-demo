package main

import (
	"fmt"
)

func foo() func() {
	return func ()  {
		fmt.Println("hello world")
	}
}

func getMsg() func(string) string {
	return func (name string) string {
		return "hello" + name
	}
}

func main() {
		bar := foo()
		bar()

		fmt.Println(getMsg()("ugu"))
}