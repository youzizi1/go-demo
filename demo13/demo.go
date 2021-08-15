package main

import "fmt"

type user struct {
	name string
	password string
	age int
}

func main() {
	ugu := user{name: "ugu", password: "123", age: 20}
	fmt.Println(ugu, ugu.name)

	// 指针
	xiaoming := &ugu
	fmt.Println(xiaoming, xiaoming.name)

	// 修改
	ugu.age = 30
	fmt.Println(xiaoming.age)
}