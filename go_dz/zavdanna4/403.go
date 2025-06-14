// Hrytsenko Iryna Valeriivna Task3
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
	a := r.Float64() * 10000
	a = math.Round(a) / 1000
	fmt.Println("a = ", a)

	x := Huge(float64(n), a)
	fmt.Println("result = ", x)

}

func Huge(n, a float64) float64 {
	res := 1.0
	for k := 1; float64(k) <= n; k++ {
		res *= a + math.Sin(0.1*float64(k)*a)
	}
	return res
}
