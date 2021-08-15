package main

import "fmt"

func byValue(value int)  {
	value = 1
}

func byPointer(ptr *int)  {
	*ptr = 10
}

func main() {
	i := 20
	fmt.Println(i)
	// 值传递
	byValue(i)
	// 结果还是10，不会影响原变量
	fmt.Println(i)
	// 地址传递，重新赋值会影响原变量
	byPointer(&i)
	fmt.Println(i, &i)
}