package main

import (
	"fmt"
)

func main() {
	var (
		name string = "pravee"
		age  int    = 20
	)
	var height, width int = 30, 30
	const (
		NAME string = "Broo"
	)
	lol := "Greeting!"
	fmt.Println(height, width)
	fmt.Println(name, age)
	fmt.Println(lol)

	fmt.Printf("%v, %T\n", name, age)
	fmt.Printf("%v, %T", NAME, NAME)

}
