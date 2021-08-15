package main

import (
	"fmt"
	"errors"
)

func fn1(num int) (int, error) {
	if num < 0 {
		return -1, errors.New("参数错误")
	}
	return 2 * num, nil
}

// 自定义错误
type argError struct {
	arg int
	msg string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d -> %s", e.arg, e.msg)
}

func fn2(num int) (int, error) {
	if num < 0 {
		return -1, &argError{num, "参数不能为负数"}
	}
	return 2 * num, nil
}

func main() {
	// fn1
	result, err := fn1(10)
	fmt.Println(result, err)
	result, err = fn1(-1)
	fmt.Println(result, err)
	// fn2
	result1, err1 := fn2(10)
	fmt.Println(result1, err1)
	result1, err1 = fn2(-1)
	fmt.Println(result1, err1)
}