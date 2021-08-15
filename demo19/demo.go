package main	

import (
	"fmt"
)

func main() {
	f := func (a, b int) int {
		return a + b
	}

	fmt.Println(f(1,2))

	result := func (a, b int) int {
		return a + b
	}(100, 200)
	fmt.Println(result)
}