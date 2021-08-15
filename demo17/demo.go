package main

import (
	"os"
	"os/exec"
	"fmt"
	"time"
)

func sayHi(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println(name + "hello world", i)
	}
}

func main() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	// 同步执行
	sayHi("ugu")
	// 异步执行
	go sayHi("ugu1")
	go sayHi("ugu2")
	go sayHi("ugu3")
	go sayHi("ugu4")
	// 异步执行匿名函数
	go func(name string) {
		fmt.Println(name)
	}("lala")

	time.Sleep(time.Second)
}