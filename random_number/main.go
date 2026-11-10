package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var num = rand.Intn(10) + 1

	fmt.Println("the number:", num)

	var num1 = rand.Intn(10) + 1

	fmt.Println("the number:", num1)

	// random distance

	var distance = rand.Intn(35000000) + 56000000

	fmt.Println("the distance:", distance)

}
