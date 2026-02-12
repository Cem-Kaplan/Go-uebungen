package main

import (
	"fmt"
	"time"
)

func printHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Hello")
	}
}

func printWorld() {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("World")
	}
}

func main() {
	fmt.Println("Goroutinen")

	go printHello()
	printWorld()
}
