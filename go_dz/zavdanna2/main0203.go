package main

import (
	"fmt"
	"math"
)

func main3() {
	fmt.Println("This program can compare two expression that consist of three nambers.Input 3 float numbers, please")
	var x, y, z float64
	fmt.Scan(&x, &y, &z)
	var max float64
	max = math.Max(x+y+z, x*y*z)
	if x+y+z == x*y*z {
		fmt.Println("Expressions are equal and = ", max)
	} else {

		if x+y+z > x*y*z {
			fmt.Println("sum > product")
			fmt.Println("sum = ", max)
		} else {
			fmt.Println("product > sum")
			fmt.Println("product = ", max)
		}
	}
}
