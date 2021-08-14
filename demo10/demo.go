package main

import "fmt"

func main() {
	// for range用来迭代
	nums := []int{4,5,6}
	num := 0
	// 第一个值是索引，用不到就用_代替
	for _, v := range nums {
		fmt.Println(v)
		num += v
	}
	fmt.Println(num)
	// 遍历map
	m := map[string]int{"k1": 10, "k2": 20, "k3": 30}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 遍历字符串
	s := "hello world"
	for i,v := range s {
		fmt.Println(i, string(v))
	}
}