package main

import "fmt"

func main00() {
	fmt.Println("Please, input  three-digit number:")
	var x uint32
	fmt.Scan(&x)
	a := x / 100
	b := (x / 10) % 10
	c := x % 10
	var r uint32
	if a%2 == 0 {
		r = a * 11
	} else {
		r = 9 - a
	}

	if b%2 == 0 {
		r = r*100 + b*11
	} else {
		r = r*10 + (9 - b)
	}

	if c%2 == 0 {
		r = r*100 + c*11
	} else {
		r = r*10 + (9 - c)
	}
	fmt.Println(r)
}

//  go run main2.go
