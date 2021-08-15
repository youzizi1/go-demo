package main

import (
	"fmt"
	"strconv"
)
// 协程通信
func main() {
	message := make(chan string)
	
	go func() {
		for i := 0; i < 3; i++ {
			message <- (strconv.Itoa(i) + ".hello world")
		}
	}()

	result := ""
	result = <-message
	fmt.Println(result)
	result = <-message
	fmt.Println(result)
	result = <-message
	fmt.Println(result)
}