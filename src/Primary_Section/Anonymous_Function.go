package main

func main() {
	sum := func(n ...int) int { // Anonymous_Function Justice
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}
	result := sum(1, 2, 3, 4, 5) // Anonymous_Function Called
	println(result)
}

func main2() {
	// Variable add Anonymous Assignment
	add := func(i int, j int) int {
		return i + j
	}

	// add func relay
	r1 := calc(add, 10, 20)
	println(r1)

	// First Parameter Anonymous_Function define
	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	println(r2)

}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

// 원형 define
type calculator func(int, int) int

// calculator 원형 used
func calc(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}
