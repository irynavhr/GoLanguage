// Гриценко Ірина Валеріївна №3
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var n = r.Intn(99) + 1

func sum(ptrA [100]*float64) float64 {
	sum := 0.00
	for i := 0; i < n; i++ {
		sum += *ptrA[i]
	}
	return sum
}
func createrA() [100]float64 {
	var arr [100]float64
	for i := 0; i < n; i++ {
		arr[i] = r.Float64() * 10
	}
	return arr
}
func createrB(a [100]*float64) [100]float64 {
	sum := sum(a)

	var b [100]float64
	for i := 0; i < n; i++ {
		b[i] = (sum - *a[i]) / float64(n-1)
	}
	return b
}
func main() {

	fmt.Println("n:", n)

	a := createrA()
	fmt.Println("Масив A:\n", a)

	var ptrA [100]*float64
	for ptr := 0; ptr < len(ptrA); ptr++ {
		ptrA[ptr] = &a[ptr]
	}

	b := createrB(ptrA)

	fmt.Println("Масив B:\n", b)
}
