package main

import (
	"fmt"
	"time"
)

func main() {
	// generate integer channel
	ch := make(chan int)

	go func() {
		ch <- 123 // send 123 to channel
	}()

	var i int
	i = <-ch // Receive 123 from channel
	println(i)

	// buffering

	c := make(chan int)
	c <- 1           //no send routine deadlock
	fmt.Println(<-c) //if comment deadlock (because except no go routine)

	// no send version
	ch1 := make(chan int, 1)

	// if no receiver can send
	ch1 <- 101

	fmt.Println(<-ch1)

	// channel parameter

	ch2 := make(chan string, 1)
	sendChan(ch2)
	receiveChan(ch2)

	// chnnel close

	ch3 := make(chan int, 2)

	// send channel
	ch3 <- 1
	ch3 <- 2

	// 채널을 닫는다
	close(ch3)

	// 채널 수신
	println(<-ch3)
	println(<-ch3)

	if _, success := <-ch3; !success {
		println("no more data.")
	}

	// channel range

	ch4 := make(chan int, 2)

	// send cha4
	ch4 <- 1
	ch4 <- 2

	// closed cha4
	close(ch4)

	// select1
	// Continue listening until it detects that the channel is closed
	/*
	   for {
	       if i, success := <-ch; success {
	           println(i)
	       } else {
	           break
	       }
	   }
	*/

	// select2
	// Channel range statement same as above expression
	for i := range ch {
		println(i)
	}

	// channel select

	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

EXIT:
	for {
		select {
		case <-done1:
			println("run1 done")

		case <-done2:
			println("run2 done")
			break EXIT
		}
	}

}

func sendChan(ch2 chan<- string) {
	ch2 <- "Data"
	// x := <-ch2 // error
}

func receiveChan(ch2 <-chan string) {
	data := <-ch2
	fmt.Println(data)
}

// channel select func
func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}
