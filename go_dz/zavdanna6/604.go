// Гриценко Ірина Валеріївна №4
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var n = r.Intn(99) + 1

func creater() [100]float64 {
	var arr [100]float64
	for i := 0; i < n; i++ {
		arr[i] = r.Float64() * 10
	}
	return arr
}

func main() {

	fmt.Println("n:", n)

	x := creater()
	fmt.Println("Масив X:\n", x)

	y := creater()
	fmt.Println("Масив Y:\n", y)

	s := 0.0
	for i := 0; i < n; i++ {
		s += x[i] * y[i]
	}

	fmt.Println("Сума добутків відповідних елементів масивів X та Y:\n", s)
}
