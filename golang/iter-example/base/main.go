package main

import "fmt"

func main() {
	for n := range seq {
		fmt.Printf("in loop:%d\n", n)
	}
}

func seq(yield func(int) bool) {
	fmt.Println("in seq")
	yield(10)
	yield(20)
}
