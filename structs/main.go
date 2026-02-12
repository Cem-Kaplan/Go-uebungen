package main

import "fmt"

type Person struct {
	name  string
	alter int
}

func main() {
	fmt.Println("structs")
	Cem := Person{"Cem Kaplan", 19}
	fmt.Println(Cem)
}
