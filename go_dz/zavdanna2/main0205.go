package main

import "fmt"

func main5() {
	var x, y, z float64
	fmt.Println("This program can solve an expression that depends on what number is greater.Input 2 float numbers, please")
	fmt.Scan(&x, &y)
	if x > y {
		z = x - y
		fmt.Println("Result: ", z)
	} else {
		z = y - x + 1
		fmt.Println("Result: ", z)
	}
}
