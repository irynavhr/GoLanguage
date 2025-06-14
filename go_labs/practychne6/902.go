// Гриценко Ірина Валеріївна Ex 9.2
package main

import (
	"fmt"
)

func main() {
	var numbers [10]int
	for x := 0; x < 10; x++ {
		numbers[x] = x
	}
	fmt.Println(numbers)
	fmt.Println(len(numbers))

	words := [7]string{"hi", "how", "are", "you", "let", "it", "snow"}
	fmt.Println(words)
	fmt.Println(len(words))

	date := [...]string{"1.06", "6.12", "1.01", "25.12", "14.03", "7.07", "24.08"}
	fmt.Println(date)
	fmt.Println(len(date))
}
