package main	

import (
	"fmt"
	"strconv"
)

func main() {
	message := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			if i == 4 {
				message <- ""
			} else {
				message <- ("hello world" + strconv.Itoa(i))
			}
		}
	}()

	for result := range message {
		if result == "" {
			break
		} else {
			fmt.Println(result)
		}
	}
}