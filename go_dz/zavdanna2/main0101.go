package main

import (
	"fmt"
	"math"
)

func maina() {
	var x, y, z float64

	fmt.Println("Input 3 numbers, please")
	fmt.Scan(&x, &y, &z)

	var denom_a float64
	denom_a = 1 + x*x/2 + y*y/4
	var num_a float64
	num_a = math.Sqrt(math.Abs((x - 1))) - math.Cbrt(math.Abs(y))
	var a float64
	a = num_a / denom_a

	var b float64
	b = x * (math.Atan(z) + math.Exp(-(x + 3)))

	fmt.Println("a =", a)
	fmt.Println("b =", b)

}
