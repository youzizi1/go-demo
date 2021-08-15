package main

import (
	"fmt"
)
// channel buffering
func main() {
	// 建立有三个缓冲区的通道，当缓冲区消息达到三个时会阻塞当前线程
	message := make(chan string, 3)

	message <- "消息1"
	message <- "消息2"
	message <- "消息3"

	// 接受消息
	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}