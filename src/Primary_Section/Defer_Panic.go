package main

import (
	"fmt"
	"os"
)

func main() {
	// defer function start
	f, err := os.Open("1.txt")
	if err != nil {
		panic(err)
	}

	// main at last run the close
	defer f.Close()

	// read file
	bytes := make([]byte, 1024)
	f.Read(bytes)
	println(len(bytes))

	// panic function start
	// Invailid input file name
	openFile("Invalid.txt")

	// in openFile( ) panic run
	// under println sentence no run
	println("Done")

	// recover function start
	// Invailid input file name
	openFile1("Invalid.txt")
	// by recover
	// run this sentence
	println("Done")
}

// panic func
func openFile(fn string) {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()
}

// recover func
func openFile1(fn string) {
	// defer func. panic called run
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}
