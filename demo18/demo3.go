package main

import (
	"fmt"
	"time"
)

func action(done chan bool)  {
	fmt.Println("start")
	time.Sleep(2 * time.Second)
	fmt.Println("end")
	done <- true
}

func main() {
	done := make(chan bool)
	go action(done)
	<-done
	// 等待2s后才会打印
	fmt.Println("hello world")
}