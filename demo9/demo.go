package main

import "fmt"

func main() {
	// 定义一个map
	m := make(map[string]int)
	// 赋值
	m["k1"] = 1
	m["k2"] = 2
	fmt.Println(m, m["k1"] + m["k2"], len(m))
	// 删除key
	delete(m, "k2")
	// 结果为0
	fmt.Println(m, m["k2"])

	// 定义且初始化
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)
}