// Гриценко Ірина Валеріївна Ex 9.7
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const k = 6
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var numbers [k]int
	for x := 0; x < 6; x++ {
		numbers[x] = r.Intn(30)
	}
	fmt.Println(numbers)

	var mofic [k]int
	for i, _ := range mofic {
		if i%2 == 0 {
			mofic[i] = numbers[i]
		} else {
			mofic[i] = numbers[i-1]
		}
	}

	fmt.Println(mofic)

}
