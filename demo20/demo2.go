package main	

import (
	"fmt"
)

func foo() func(int) int {
	count := 0
	return func (num int) int {
		count += num
		return count
	}
}

func main() {
	bar := foo()
	fmt.Println(bar(1))
	fmt.Println(bar(2))
	fmt.Println(bar(3))
}