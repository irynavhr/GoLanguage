// Гриценко Ірина Валеріївна №7в
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

func changeArr(a [10]float64) [10]float64 {
	var t [10]float64

	k := n
	if n%2 == 1 {
		k = n/2 + 1
	} else {
		k = n / 2
	}

	for i := 0; i < k; i++ {
		t[i] = a[n/2+i]

	}
	for i := 0; i < (n / 2); i++ {
		t[k+i] = a[i]

	}

	return t
}

func main() {
	a := createArr()
	fmt.Println("a:", a)
	a = changeArr(a)
	fmt.Println("a:", a)
}
