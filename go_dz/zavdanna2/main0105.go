package main

import (
	"fmt"
	"math"
)

func maine() {
	var x, y, z float64

	fmt.Println("Input 3 numbers, please")
	fmt.Scan(&x, &y, &z)

	var denom_a float64
	denom_a = 1/2 + math.Sin(y)*math.Sin(y)
	var num_a float64
	num_a = 2 * math.Cos(x-3.14/6)
	var a float64
	a = num_a / denom_a

	var num_b float64
	num_b = z * z
	var denom_b float64
	denom_b = 3 + num_b/5

	var b float64
	b = 1 + num_b/denom_b

	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
