// Гриценко Ірина Валеріївна №10
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var n = (r.Intn(4) + 1) * 2

func sign(x int) int {
	if x >= 0 {
		return 1
	}
	return -1
}

func creater() [10]int {
	var arr [10]int
	for i := 0; i < n; i++ {
		arr[i] = r.Intn(19) - 9
	}
	return arr
}

func main() {
	x := creater()
	fmt.Println("x:", x)
	result := true
	for i := 0; i < n-1; i++ {
		if sign(x[i])*sign(x[i+1]) > 0 {
			result = false
			break
		}
	}
	if result == true {
		fmt.Println("Масив x є знакозмінним")
	} else {
		fmt.Println("Масив x не є знакозмінним")
	}

}
