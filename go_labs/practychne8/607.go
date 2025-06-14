// Гриценко Ірина Валеріївна №7б
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var n = r.Intn(9) + 1

func createArr() [10]float64 {
	var arr [10]float64
	for i := 0; i < n; i++ {
		arr[i] = r.Float64() * 10
	}
	return arr
}

func main() {
	a := createArr()
	fmt.Println("a:", a)
	var t [10]float64

	for i := 0; i < n/2; i++ {
		t[i] = a[2*i+1]

	}
	for i := 0; i < n/2; i++ {
		t[n/2+i] = a[2*i]

	}
	for i := 0; i < n; i++ {
		a[i] = t[i]
	}
	fmt.Println("a:", a)
}
