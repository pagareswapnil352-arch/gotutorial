package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var number = rand.Intn(100)
	fmt.Println("the numbers are following")

	for i := 1; i <= number; i++ {

		fmt.Println(i)

	}

}
