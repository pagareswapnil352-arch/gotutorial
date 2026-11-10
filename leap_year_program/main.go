package main

import "fmt"

func main() {

	var year = 2028

	if year%4 == 0 || year%100 != 0 {

		fmt.Println("the 2028 is leap")

	} else {

		fmt.Println("2028  is not a leap year")

	}

}
