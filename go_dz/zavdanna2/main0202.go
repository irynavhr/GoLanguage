package main

import (
	"fmt"
	"math"
)

func main2() {
	fmt.Println("This program can define max and min of three nambers.Input 3 float numbers, please")
	var x, y, z float64
	fmt.Scan(&x, &y, &z)
	if x == y && y == z {
		fmt.Println("They are equel")
	} else {
		var Q float64
		Q = math.Max(x, y)
		var q float64
		q = math.Min(x, y)
		var max float64

		max = math.Max(Q, z)
		var min float64
		min = math.Min(q, z)

		fmt.Println("max is ", max)
		fmt.Println("min is ", min)
	}
}
