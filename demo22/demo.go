package main

import (
	"fmt"
)

func foo(val interface{})  {
	switch val.(type) {
	case bool:
		fmt.Println("布尔")
	case int:
		fmt.Println("整数")
	case string:
		fmt.Println("字符串")
	default:
		fmt.Println("其他")
	}
}

func main() {
	foo(1)
	foo(true)
	foo("1")
	foo(make(chan bool))

	// 接口类型转换
	var x interface{} = 100
	// 类型转换
	x1, err := x.(int)
	if !err {
		fmt.Println("转换失败", err)
	} else {
		fmt.Println("转换成功", x1, x1 + 200)
	}
}