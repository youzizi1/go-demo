package main

import "fmt"

// 不能省略var
var a = 1
var (
	b = 2
	c = "hello world"
)

var m, n = 10, 20
var l, s string = "hello", "world"

func demo() {
	fmt.Println(m, n, l, s)
	// 只初始化，int默认为0，string默认为""
	var u int
	var v string
	fmt.Println(u,v)
	// 初始化并赋值
	var a int = 3
	var b string = "abc"
	// 省略type，编译器自动推断
	var m = 1
	var n = "hello"
	// 省略var，只能在函数内部省略var
	c,d,e := true, 4, 5
	e = 6
	fmt.Println(a,b,c,d,e,m,n)
	// 指针
	var foo *int
	// 获取m的内存地址
	foo = &m
	fmt.Println(foo)
	// 获取内存对应的值
	fmt.Println(*foo)
}

func main() {
	demo()
	fmt.Println("hello world")
	// 定义常量，不能被重新赋值
	const s = "HELLO"
	fmt.Println(s)
}