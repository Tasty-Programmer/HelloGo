package main

import "fmt"

func main() {


	check(2)

	if k == 1 {
		println("One")
	}else if k == 2 {
		println("Two")
	}else {
		println("Other")
	}

	if val := i * 2; val < max {
		println(val)
	}

	val++

	var name string
	var category = 1

	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
	case 3,4:
		name = "Blog"
	default:
		name = "Other"
	}
	println(name)
	
	//Used Expression
	switch  x := category << 2; x-1 {
		//.....
	}
	
	func grade(score int){
		switch  {
		case score >= 90:
			println("A")
		case score >= 80:
			println("B")
		case score >= 70:
			println("C")
		case score >= 60:
			println("D")
		default:
			println("No Hope")
		}
	}

	switch v.(type) {
	case int:
		println("int")
	case bool:
		println("bool")
	case string:
		println("string")
	default:
		println("unknown")
	}

	func check(val int) {
		switch val {
		case 1:
			fmt.Println("1 이하")
			fallthrough
		case 2:
			fmt.Println("2 이하")
		case 3:
			fmt.Println("3 이하")
			fallthrough
		default:
			fmt.Println("default 도달")



		}
	}

}

