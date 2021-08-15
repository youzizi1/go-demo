package main

import "fmt"

type rect struct {
	width, height int
}

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