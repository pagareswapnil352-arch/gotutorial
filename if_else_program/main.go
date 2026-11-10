package main

import "fmt"

func main() {

	var command string

	fmt.Println("enter your current state")

	fmt.Scanln(&command)

	if command == "walk outside" {

		fmt.Println("the move outside the cave")
	} else if command == "stay" {

		fmt.Println("you are still inside")
	} else {
		fmt.Println("your are lost")

	}

}
