package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// WaitGroup produce . 2 go routine wait
	var wait sync.WaitGroup
	wait.Add(2)

	// AnonFunctino used Goroutine
	go func() {
		defer wait.Done() // at finished called .Done()
		fmt.Println("Hello")
	}()

	// AnonFunction send parameter
	go func(msg string) {
		defer wait.Done() // at finished called .Done()
		fmt.Println(msg)
	}("Hi")

	wait.Wait() // All go rountine finished wait

	// 4 Cpu use
	runtime.GOMAXPROCS(4)

	//......
}
