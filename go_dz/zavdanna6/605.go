// Гриценко Ірина Валеріївна №5
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var n = r.Intn(99) + 1

func create() [100]float64 {
	var arr [100]float64
	for i := 0; i < n; i++ {
		arr[i] = r.Float64() * 10
	}
	return arr
}

func main() {

	fmt.Println("n:", n)

	a := create()
	fmt.Println("Масив A:\n", a)

	b := create()
	fmt.Println("Масив B:\n", b)

	m := 1.0
	for i := 0; i < n; i++ {
		m *= a[i] + b[n-1-i]
	}

	fmt.Println("Добуток суми  протилежних елементів масивів A та B:\n", m)
}
