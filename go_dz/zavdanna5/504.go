// Гриценко Ірина Валеріївна Ex 8.4
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	a := r.Float64() * 10
	b := r.Float64() * 10
	c := r.Float64() * 10
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)

	fmt.Println(fun4(a, b, c))
}

func fun4(a, b, c float64) float64 {
	num := math.Max(a, a+b) + math.Max(a, b+c)
	denom := 1 + math.Max(a+b*c, 1.15)
	return num / denom
}
