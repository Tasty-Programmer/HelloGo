package main

func main() {
	var a [3]int // declared Array Integer  3 elements have
	a[0] = 1
	a[2] = 2
	a[2] = 3
	println(a[1])	// 2 print

	// Array reset
	var a1 = [3]int{1 , 2 , 3}
	var a3 = [...]int{1 , 2 , 3} 	// Array size Automatic

	// MultiArray
	var multiArray = [3][4][5]int 	// Justify
	multiArray[0][1][2]



	// Multidimensional Array

	var a = [2][3]int{
		{1 , 2 , 3},
		{4 , 5 , 6},	// edge , comma add
	}
	println(a[1][2])

}


