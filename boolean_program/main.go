package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("the cave story")

	var command = "walk outside"

	var exist = strings.Contains(command, "outside")

	fmt.Println("the answer:", exist)

	var option1 = true

	var option2 = false

	var result = option1 || option2

	var result1 = option1 && option2

	defer fmt.Println("the answer :", result)
	fmt.Println("the answer :", result1)

}
