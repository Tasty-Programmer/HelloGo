package main

import "fmt"

func main() {
	var m map[int]string

	m = make(map[int]string)
	// add or renewal
	m[901] = "Apple"
	m[134] = "Grape"
	m[777] = "Tomato"

	//  find key about value
	str := m[134]
	println(str)

	noData := m[999] // if no value nil or zero returned
	println(noData)

	// delete
	delete(m, 777)

	tickers := map[string]string{
		"Good": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   " FaceBook",
		"AMZN": " Amazon",
	}

	// map key check
	val, exists := tickers["MSFT"]
	if !exists {
		println("No MSFT ticker")
	}
	myMap := map[string]string{
		"A": "Apple",
		"B": "Banana",
		"C": "Charlie",
	}
	// use for range whole map elements print
	// Map is unorodered so order is random
	for key, val := range myMap {
		fmt.Println(key, val)
	}
}
