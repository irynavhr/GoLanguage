package main

import (
	"fmt"
	"math"
)

func main4() {
	fmt.Println("This program can solve an expression that consist of three variabls.Input 3 float numbers, please")
	var x, y, z float64
	fmt.Scan(&x, &y, &z)

	var min float64
	min = math.Min(x+y+z/2, x*y*z)

	var result float64
	result = min*min + 1
	fmt.Println("result = ", result)

}
