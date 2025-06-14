// Hrytsenko Iryna Valeriivna Task10
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

	x := sin1(float64(n), a)
	fmt.Println("a) ", x)
	x = sin2(float64(n), a)
	fmt.Println("b) ", x)
	x = sin3(float64(n), a)
	fmt.Println("c) ", x)
	x = sin4(float64(n), a)
	fmt.Println("d) ", x)

}

func sin1(n, a float64) float64 {
	res := 0.0
	for k := 1; float64(k) <= n; k++ {
		res += math.Pow(math.Sin(a), float64(k))
	}
	return res
}
func sin2(n, a float64) float64 {
	res := 0.0
	for k := 1; float64(k) <= n; k++ {
		res += math.Sin(math.Pow(a, float64(k)))
	}
	return res
}

func sin3(n, a float64) float64 {
	res := 0.0
	p := a
	for k := 1; float64(k) <= n; k++ {
		p = math.Sin(p)
		res += p
	}
	return res
}

func sin4(n, x float64) float64 {
	res, pn, pp := math.Sin(x), math.Sin(x), x
	for i := 1; float64(i) < math.Ceil(n/2); i++ {
		pn = math.Sin(math.Cos(pn))
		res += pn
	}
	for i := 1; float64(i) < math.Ceil(n+1/2); i++ {
		pp = math.Sin(math.Cos(pp))
		res += pp
	}
	return res
}
