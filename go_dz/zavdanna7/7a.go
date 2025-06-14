// Гриценко Ірина Валеріївна №7a
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

func changArr(a [10]float64) [10]float64 {
	var t [10]float64
	t[0] = a[n-1]
	for i := 1; i < n; i++ {
		t[i] = a[i-1]

	}
	return t
}

func main() {
	a := createArr()
	fmt.Println("a:", a)
	a = changArr(a)
	fmt.Println("a:", a)
}
