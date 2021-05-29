package main

func main() {
	msg := "Hello"
	say(msg)
}

func say(msg string) {
	println(msg)
}

func main2() {
	msg := "Hello"
	say(&msg)
	println(msg) // changed Message print
}

func say2(msg *string) {
	println(*msg)
	*msg = "Changed" // changed Message

}

func main3() {
	say("This", "is", "a", "book")
	say("Hi")

}

func say3(msg ...string) {
	for _, s := range msg {
		println(s)
	}

}

func main4() {
	total := sum(1, 7, 3, 5, 9)
	println(total)
}

func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func main5() {
	count2, total2 := sum(1, 7, 3, 5, 9)
	println(count2, total2)
}

func sum2(nums ...int) (int, int) {
	s := 0     // Sum
	count := 0 // element
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

func sum3(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}
