package main

import (
	"fmt"
)

func main() {
	defer func ()  {
		if err := recover(); err != nil {
			fmt.Println("捕获错误", err)
		}
	}()

	// 抛出错误
	panic("手动报错")
}