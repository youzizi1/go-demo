package main

import (
	"fmt"
	"time"
)

func main() {
	i := 10
	switch i {
	case 5:
		fmt.Println("equal 5")
	case 10:
		fmt.Println("equal 10")
	default:
		fmt.Println("default case")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("before none")
	default:
		fmt.Println("after noon")
	}
}