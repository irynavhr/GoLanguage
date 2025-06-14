// Гриценко Ірина Валеріївна Ex 8.5
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	a := r.Float64() * 10
	b := r.Float64() * 10
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	u := fun5(a, b)
	v := fun5(a*b, a+b)
	fmt.Println(u)
	fmt.Println(v)
	fmt.Println(fun5(u+v*v, 3.1417))
}

func fun5(a, b float64) float64 {
	if a > b {
		return b
	}
	return a
}
