package main

import "fmt"

func fn1(a int, b int) int  {
	return a + b
}

func fn2(a, b, c int) int {
	return a + b + c
}

func fn3(a, b int) (int, int, int, int)  {
	return a + b, a - b, a * b, a / b
}

func fn4(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

// 递归调用
func fn5(num int) int {
	if num == 1 {
		return num		
	}
	return fn5(num - 1) + num
}

func main() {
	fmt.Println(fn1(1, 2))
	fmt.Println(fn2(1, 2, 3))
	fmt.Println(fn3(1, 2))
	fmt.Println(fn4(1, 2, 3, 4, 5))
	fmt.Println(fn4(1, 2, 3))
	fmt.Println(fn5(10))
}