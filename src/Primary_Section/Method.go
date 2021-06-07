package main

// Rect - struct define

type Rect struct {
	width, height int
}

// Rect area() method
func (r Rect) area() int {
	return r.width * r.height
}

// Pointer Receiver
func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func main() {

	rect := Rect{10, 20}
	area := rect.area2() // method call
	println(area)

	rect2 := Rect{10, 20}
	area2 := rect2.area2()      // method call
	println(rect2.width, area2) // 11 220 print
}
