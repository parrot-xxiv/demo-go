package main

import (
	"demo-go/greet"
	"fmt"
)

func main() {
	var name string

	fmt.Println("Enter your name:")
	fmt.Scanln(&name)

	greet.Greet(name)
}
