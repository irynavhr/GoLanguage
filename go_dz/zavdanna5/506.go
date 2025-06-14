// Гриценко Ірина Валеріївна Ex 8.6
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

	n := fun6(s-t, s*t)
	n0 := n * n
	m := fun6(s-t, s+t)
	m0 := m * m * m * m

	h := fun6(s, t) + math.Max(n0, m0) + fun6(1, 1)
	fmt.Println(h)
}

func fun6(a, b float64) float64 {
	x1 := a / (1 + b*b)
	x2 := b / (1 + a*a)
	x3 := (a - b) * (a - b) * (a - b)
	return x1 + x2 - x3
}
