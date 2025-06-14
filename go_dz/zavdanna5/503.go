// Гриценко Ірина Валеріївна Ex 8.3
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	y := r.Float64() * 10
	fmt.Println("y:", y)

	t1 := 1.7 * fun3(0.25)
	t2 := 2 * fun3(1+y)
	t3 := 6 - fun3(y*y-1)

	fmt.Println((t1 + t2) / t3)
}

func fun3(x float64) float64 {
	num := 0.0
	for k := 0; k < 11; k++ {
		num += math.Pow(x, float64(2*k+1)) / float64(fuctorial(2*k+1))
	}
	denom := 0.0
	for k := 0; k < 11; k++ {
		denom += math.Pow(x, float64(2*k)) / float64(fuctorial(2*k))
	}
	return num / denom
}

func fuctorial(i int) int {

	if i == 0 {
		return 1
	}
	if i == 1 {
		return 1
	}
	return i * fuctorial(i-1)
}
