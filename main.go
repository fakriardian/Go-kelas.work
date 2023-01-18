package main

import "fmt"


type shape interface {
    getArea() float64
    getPerimeter() float64
}

type rectangle struct {
    width float64
    height float64
}

func (r rectangle) getArea() float64{
    return r.width * r.height
}

func (r rectangle) getPerimeter() float64{
    return 2*r.width + 2*r.height
}

func measureShape(s shape){
    fmt.Printf("get area %v\n", s.getArea())
    fmt.Printf("get perimeter %v\n", s.getPerimeter())
}

func main() {
    myRectangle := rectangle{
        width: 60,
        height: 30,
    }

    measureShape(myRectangle)
}