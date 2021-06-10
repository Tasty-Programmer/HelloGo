package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

// Rect define
type Rect struct {
	width, height float64
}

// Circle define
type Circle struct {
	radius float64
}

// Shape Interface Implementation about Rect type
func (r Rect) area() float64 { return r.width * r.height }
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Shape Interface Implementation about Circle type
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func main() {
	r := Rect{10., 20.}
	c := Circle{10}

	showArea(r, c)
}
func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area()
		println(a)
	}
}

func Marshal(v interface{}) ([]byte, error)

func Println(a ...interface{}) (n int, err error)

func main2() {
	var x interface{}
	x = 1
	x = "Tom"

	printIt(x)

}

func printIt(v interface{}) {
	fmt.Println(v) // Tom
}

func main3() {
	var a interface{} = 1

	i := a       // a and i is dynamic type, value is 1
	j := a.(int) // j is int type, value is 1

	println(i) // print pointer address
	println(j) // print 1
}
