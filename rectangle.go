package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func area(r Rectangle) float64 {
	return r.width * r.height
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	c.radius = 30
	return c.radius * c.radius * math.Pi
}

func main() {
	R := Rectangle{30, 40}

	// 常规写法
	fmt.Println(area(R))
	// method 写法
	fmt.Println(R.area())

	C := Circle{40}

	fmt.Println(C.area())
	// 值传递
	fmt.Println(C.radius)

}
