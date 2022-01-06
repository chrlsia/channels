package main

import "fmt"

func main() {
	go doSomething("Hello, world")

	for {
		// do something
	}
}

func doSomething(s string) {
	for {
		fmt.Println("s is", s)
	}
}