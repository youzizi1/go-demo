package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./demo25/demo.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开错误", err)
	}
	file.WriteString("hello world")
	// file.Close()
}