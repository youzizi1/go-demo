package main

import "fmt"

type rect struct {
	width, height int
}
// go定义结构体的方法不是在结构体内部，而是在结构体外面
func (r *rect) area() int  {
	return r.width * r.height
}

func (r rect) perim() int  {
	return 2 * (r.width + r.height)
}

func main() {
	r1 := rect{width: 10, height: 20}
	fmt.Println(r1.area(), r1.perim())

	// 指针
	r2 := &r1
	fmt.Println(r2.area(), r2.perim())
}