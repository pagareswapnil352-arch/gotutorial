package main

import "fmt"

func main() {

	const lightspeed = 299979

	var marsdistance = 56000000

	var earthdistance = 4000000

	var time1 = marsdistance / lightspeed

	var time2 = earthdistance / lightspeed

	fmt.Println("the time required :", time1, "second")

	fmt.Println("the time required:", time2, "seconds")

	// increment operator
	var age = 18

	age++

	fmt.Println(age)
	// decrement operator

	age--
	fmt.Println(age)
}
