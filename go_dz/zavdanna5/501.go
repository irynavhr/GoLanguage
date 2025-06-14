// Гриценко Ірина Валеріївна Ex 8.1
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := r.Float64() * 10
	t := r.Float64() * 10
	fmt.Println("s:", s)
	fmt.Println("t:", t)

	r1 := fun1(t, (-2 * s), 1.17)
	fmt.Println(r1)

	r2 := fun1(2.2, t, s-t)
	fmt.Println(r2)

	fmt.Println(r1 + r2)
}

func fun1(a, b, c float64) float64 {
	num := 2*a - b - math.Sin(c)
	denom := 5 + math.Abs(c)
	return num / denom
}
