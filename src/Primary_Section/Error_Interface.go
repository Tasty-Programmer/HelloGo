package main

import (
	"log"
	"os"
)

type error interface {
	Error() string
}

import (
	"log"
	"os"
)

func main() {
	f, err  := os.Open("C:\\temp\\1.txt")
	if err != nil{
		log.Fatal(err.Error())
	}
	println(f.Name())


	_, err1 := otherFunc()
	switch err1.(type) {
	default:	// no error
		println("ok")
	case MyError:
		log.Println("Log my error")
	case error:
		log.Fatal(err.Error())
	}
}
