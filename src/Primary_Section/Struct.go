package main

import "fmt"

// struct define
type person struct {
	name string
	age  int
}

type dict struct {
	data map[int]string
}

// Construct function define
func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d // Pointer use
}
func main() {
	// person object produce
	p := person{}

	// field value
	p.name = "Lee"
	p.age = 10

	dic := newDict() // struct call
	dic.data[1] = "A"

	fmt.Println(p)
}
