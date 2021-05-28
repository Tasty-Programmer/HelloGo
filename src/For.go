package main

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)

	n := 1
	for n < 100 {
		n *= 2
		// if n > 90 {
		// 		break
		//}
	}
	println(n)

	for {
		println("Infinite loop")
	}

	names := []string{"홍길동", "이순신", "강감찬"}

	for index, name := range names {
		println(index, name)
	}

	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue // for loop continue

		}
		a++
		if a > 10 {
			break // exit loop
		}
		if a == 11 {
			goto END // goto example
		}
		println(a)
	}

END:
	println("END")

L1:
	for {

		if i == 0 {
			break L1
		}
	}
	println("OK")
}
