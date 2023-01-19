package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
	getPerimeter() float64
}

func measureShape(s shape) {
	fmt.Printf("get area %v\n", s.getArea())
	fmt.Printf("get perimeter %v\n", s.getPerimeter())
}

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (r rectangle) getArea() float64 {
	return r.width * r.height
}

func (r rectangle) getPerimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) getArea() float64 {
	return math.Phi * c.radius * c.radius
}

func (c circle) getPerimeter() float64 {
	return 2 * math.Phi * c.radius
}

func main() {
	myRectangle := rectangle{
		width:  60,
		height: 30,
	}

	myCircle := circle{
		radius: 7,
	}

	measureShape(myRectangle)
	measureShape(myCircle)
}
