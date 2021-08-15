package main

import (
	"fmt"
	"math"
	"reflect"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry)  {
	fmt.Println(reflect.TypeOf(g), g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r1 := rect{width: 20, height: 10}
	c1 := circle{radius: 5}
	measure(r1)
	measure(c1)
}