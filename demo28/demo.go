package main

import (
	"fmt"
	"time"
)
// 协程接受到消息后并处理
func worker(name string, msg chan int)  {
	for {
		val, ok := <-msg
		if ok {
			fmt.Println(name, val, ok)
		} else {
			break
		}
	}
}

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	defer close(chan1)
	defer close(chan2)

	go worker("玩家1", chan1)
	go worker("玩家2", chan2)
	// 主线程循环给协程发送消息
	i := 1
	for i <= 5 {
		chan1 <- i
		chan2 <- i
		i++
		time.Sleep(time.Second)
	}
}