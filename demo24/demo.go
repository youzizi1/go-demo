package main

import "fmt"

func foo()  {
	// 先进后出
	defer fmt.Println("foo 1")
	defer fmt.Println("foo 2")
	defer fmt.Println("foo 3")
}

func main() {
	fmt.Println("start")
	defer func ()  {
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()
	fmt.Println("end")

	foo()
}