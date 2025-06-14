// Гриценко Ірина Валеріївна №9б
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var n = (r.Intn(9) + 1) * 2

func createArr() [20]float64 {
	var arr [20]float64
	for i := 0; i < n; i++ {
		arr[i] = r.Float64() * 10
	}
	return arr
}

func adderArrs(a [20]float64) [10]float64 {
	var t [10]float64

	for i := 0; i < n/2; i++ {
		t[i] = a[i] + a[n-1-i]
	}
	return t
}

func main() {
	a := createArr()
	fmt.Println("a:", a)
	s := adderArrs(a)
	fmt.Println("s:", s)
}
