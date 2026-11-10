package main

import "fmt"

//var command = "go inside"

//switch command {

//case "go inside":
//	fmt.Println("your entering the cave")

//case "go outside":
//fmt.Println("your getting out of cave")

//fallthrough

//case "stay inside":
//	fmt.Println("your inside the cave ")

//	}

func main() {
	var command = "go inside"

	switch command {

	case "go inside":
		fmt.Println("your entering the cave")
		fallthrough

	case "go outside":
		fmt.Println("your getting out of cave")

	case "stay inside":
		fmt.Println("your inside the cave ")
	}
}
