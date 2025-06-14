package main

import "fmt"

func main1() {
	var x, y float64

	fmt.Println("This program can define max and min of two nambers. Please input two numbers:")
	fmt.Scan(&x, &y)

	if x == y {
		fmt.Println("Impossible to define max or min of the numbers.They are equel!")
	} else {
		if x > y {
			fmt.Println("max is:", x)
			fmt.Println("min is:", y)
		} else {
			fmt.Println("max is:", y)
			fmt.Println("min is:", x)
		}

	}
}
