// Гриценко Ірина Валеріївна №11г
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var n = (r.Intn(19) + 1)

func createArr() [20]int {
	var arr [20]int
	for i := 0; i < n; i++ {
		arr[i] = r.Intn(8) + 1
	}
	return arr
}

func dellInArr(a [20]int) [20]int {
	var t [20]int
	k := 0
	for i := 0; i < n; i++ {
		if a[i]%2 == 1 {
			t[k] = a[i]
			k++
		}
	}

	return t
}

func main() {
	x := createArr()
	fmt.Println("x:", x)
	n := dellInArr(x)
	fmt.Println("n:", n)
}
