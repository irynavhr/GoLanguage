// Hrytsenko Iryna Valeriivna Task6
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(49) + 1
	fmt.Println("n = ", n)
	a := r.Float64() * 1000
	a = math.Round(a) / 100
	fmt.Println("a = ", a)

	x := ProdDiv(float64(n), a)
	fmt.Println("Result = ", x)

}

func ProdDiv(n, a float64) float64 {
	num := 1.0
	denom := 1.0
	for k := 1; float64(k) <= n; k++ {
		num *= a - math.Pow(2, float64(k))
		denom *= a - (math.Pow(2, float64(k)) - 1)
	}
	res := num / denom
	return res
}
