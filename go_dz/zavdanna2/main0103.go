package main

import (
	"fmt"
	"math"
)

func mainc() {
	var x, y, z float64

	fmt.Println("Input 3 numbers, please")
	fmt.Scan(&x, &y, &z)

	var denom_a float64
	denom_a = math.Exp(-x-2) + 1/(x*x+4)
	var num_a float64
	num_a = x + y/(x*x+4)
	var a float64
	a = (1 + y) * num_a / denom_a

	var denom_b float64
	denom_b = math.Pow(x, 4)/2 + math.Sin(z)*math.Sin(z)
	var num_b float64
	num_b = 1 + math.Cos(y-2)
	var b float64
	b = num_b / denom_b

	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
