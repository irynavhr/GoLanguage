// Гриценко Ірина Валеріївна №6
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var n = r.Intn(19) + 1

func creat() [100]float64 {
	var arr [100]float64
	for i := 0; i < n; i++ {
		arr[i] = r.Float64() * 10
	}
	return arr
}

func main() {

	fmt.Println("n:", n)

	a := creat()
	fmt.Println("Масив A:\n", a)

	for i := 0; i < n/2; i++ {
		t := a[i]
		a[i] = a[n-1-i]
		a[n-1-i] = t
	}
	fmt.Println("Обернений масив A:\n", a)
}
