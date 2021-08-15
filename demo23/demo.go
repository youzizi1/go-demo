package main

import "fmt"

func main() {
	loop:
	for {
		for {
			fmt.Println("死循环中...")
			break loop
		}
	}
	fmt.Println("程序结束")
}