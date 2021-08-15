package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 错误写法：strAge string := "25"
	strAge := "25"
	intAge, err := strconv.Atoi(strAge)
	if err != nil {
		fmt.Println("转换失败", err)
	} else {
		fmt.Println("转换成功", intAge)
	}
}
