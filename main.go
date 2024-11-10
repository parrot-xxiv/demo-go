package main

import (
	"fmt"
)

func checkLength() bool {
	var password string
	fmt.Println("Enter password:")
	fmt.Scanln(&password)
	return len(password) >= 5
}

// Dummy functions to simulate options
func optionA() {
	fmt.Println("Option A selected")
}

func optionB() {
	fmt.Println("Option B selected")
}

func main() {
	for {
		var x int
		fmt.Println("Choose app: ")
		fmt.Println("1 Option A")
		fmt.Println("2 Option B")
		fmt.Println("3 Exit")
		fmt.Scan(&x)

		switch x {
		case 1:
			optionA()
		case 2:
			optionB()
		case 3:
			fmt.Println("Exited")
			return // return instead of break
		default:
			fmt.Println("Not an option")
		}
	}
}
