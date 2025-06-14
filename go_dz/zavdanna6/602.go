// Гриценко Ірина Валеріївна №2
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var n = r.Intn(99) + 1

func createArr() [100]float64 {
	var arr [100]float64
	for i := 0; i < n; i++ {
		arr[i] = r.Float64() * 10
	}
	return arr
}

func average(ptrA [100]*float64) float64 {
	sum := 0.00
	for i := 0; i < n; i++ {
		sum += *ptrA[i]
	}
	return sum / float64(n)
}

func formula(ptrA [100]*float64) float64 {
	aver := average(ptrA)
	fmt.Println("averege:", aver)
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += (*ptrA[i] - aver) * (*ptrA[i] - aver)
	}
	return math.Sqrt(sum / float64(n-1))
}

func main() {

	fmt.Println("n:", n)

	a := createArr()

	var ptrA [100]*float64
	for ptr := 0; ptr < len(ptrA); ptr++ {
		ptrA[ptr] = &a[ptr]
	}

	fmt.Println("Result:", formula(ptrA))
}
