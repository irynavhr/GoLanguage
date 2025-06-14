// Гриценко Ірина Валеріївна Ex 8.7
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var a0 = r.Float64() * 10
var a1 = r.Float64() * 10
var a2 = r.Float64() * 10
var a3 = r.Float64() * 10

func main() {
	x := float64(r.Intn(10) + 1)
	fmt.Println("x:", x)
	// fmt.Println("a0:", a0)
	// fmt.Println("a1:", a1)
	// fmt.Println("a2:", a2)
	// fmt.Println("a3:", a3)
	res := fun7(x+1) - fun7(x)
	fmt.Println("res:", res)
}

func fun7(y float64) float64 {
	return a3*y*y*y + a2*y*y + a1*y + a0
}
