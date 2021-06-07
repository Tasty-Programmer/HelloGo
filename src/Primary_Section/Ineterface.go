package main

type Shape interface {
	area() float64
	perimeter() float64
}


// Rect define
type Rect struct {
	width , height float64
}

// Circle define
type Circle struct {
	radius float64
}

// Shape Interface Implementation about Rect type 
func (r Rect) area() float64 { return r.width * r.height}
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func ()  {
	
}

func main() {

}
