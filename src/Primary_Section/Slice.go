package main

import "fmt"

func main() {

	var a []int        // Slice variable declare
	a = []int{1, 2, 3} // slice literal assigment
	a[1] = 10
	fmt.Println(a) // [1 , 10 , 3] 출력

	s := make([]int, 5, 10)
	println(len(s), cap(s)) // len 5 , cap 10

	var s []int

	if s == nil {
		println("Nil Slice")
	}
	println(len(s), cap(s)) // 모두 0

	s := []int{0, 1, 2, 3, 4, 5}
	s := s[2:5]    // 2 , 3 , 4
	s = s[1:]      // 3 , 4
	fmt.Println(s) //  3 , 4 print

	s := []int{0, 1}

	// one expansion
	s = append(s, 2) // 0, 1, 2
	// another elements expansion
	s = append(s, 3, 4, 5) // 0,1,2,3,4,5

	fmt.Println(s)

	// Slice element add

	// len = 0, cap = 3 slice
	sliceA := make([]int, 0, 3)

	// keep elements add
	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		// slice length and size
		fmt.Println(len(sliceA), cap(sliceA))
	}

	fmt.Println(sliceA)

	sliceA := []int{1, 2, 3}
	sliceB := []int{4, 5, 6}

	sliceA = append(sliceA, sliceB...)
	//sliceA = append(sliceA, 4, 5, 6)

	fmt.Println(sliceA) // [1 2 3 4 5 6] print

	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source)
	fmt.Println(target)               // [0 1 2 ] print
	println(len(target), cap(target)) // 3, 6 print
}
