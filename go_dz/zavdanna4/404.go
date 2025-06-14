// Hrytsenko Iryna Valeriivna Task4
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(99) + 1
	fmt.Println("n = ", n)
	a := r.Float64() * 1000
	a = math.Round(a) / 100
	fmt.Println("a = ", a)

	x := VeeeeryHuge(float64(n), a)
	fmt.Println("Result = ", x)

}

func VeeeeryHuge(n, a float64) float64 {
	res := 0.0
	for k := 1; float64(k) <= n; k++ {
		znak := math.Pow(-1, float64(k-1))
		num := math.Pow(a, float64(2*k-1))
		d := denom(2*k - 1)
		res += znak * num / float64(d)
	}
	return res
}

func denom(q int) int {
	res := 0
	for k := 1; k <= q; k++ {
		res += k
	}
	return res
}
