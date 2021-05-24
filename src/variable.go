package main

func main() {
	const c int = 10
	const s string = "Hi"

	const (
		Visa   = "visa"
		Master = "MasterCard"
		Amex   = "American Express"
	)

	const (
		Apple  = iota // identifier start point is 0
		Grape         // 1
		Orange        // 2
	)

	println(c)
	println(s)
	println(Visa, Master, Amex)
	println(Apple, Grape, Orange)
}
