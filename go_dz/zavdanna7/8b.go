// Гриценко Ірина Валеріївна №8б
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var n = (r.Intn(4) + 1) * 2

func createArr() [10]float64 {
	var arr [10]float64
	for i := 0; i < n; i++ {
		arr[i] = r.Float64() * 10
	}
	return arr
}

func changer(a [10]float64) [10]float64 {
	var t [10]float64

	for i := 0; i < n/2; i++ {
		t[2*i] = a[i]
	}
	for i := 0; i < n/2; i++ {
		t[2*i+1] = a[n-1-i]
	}
	return t
}

func main() {
	a := createArr()
	fmt.Println("a:", a)
	a = changer(a)
	fmt.Println("a:", a)
}
