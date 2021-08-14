package main

import "fmt"

func main() {
	// 空循环，类似于while语法
	for {
		fmt.Println("hello world")
		break
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for j := 0; j < 10; j++ {
		if j > 6 {
			//break
			return
		}
		fmt.Println(j)
	}
	// return后面的语句不会执行
	fmt.Println("hello world")
}
