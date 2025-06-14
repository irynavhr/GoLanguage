// Гриценко Ірина Валеріївна Ex 8.2
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := r.Float64() * 10
	t := r.Float64() * 10
	fmt.Println("s:", s)
	fmt.Println("t:", t)

	g1 := fun2(1.2, s)
	fmt.Println(g1)

	g2 := fun2(t, s)
	fmt.Println(g2)

	g3 := fun2(2*s-1, s*t)
	fmt.Println(g3)

	fmt.Println(g1 + g2 + g3)
}

func fun2(a, b float64) float64 {
	num := a*a + b*b
	denom := a*a + 2*a*b + 3*b*b + 4
	return num / denom
}
