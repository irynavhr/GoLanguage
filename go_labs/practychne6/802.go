// Гриценко Ірина Валеріївна Ex 8.2
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(50)
	diameter, circumference, area := sphereStuff(n)
	fmt.Println("diameter:", diameter)
	fmt.Println("circumference:", circumference)
	fmt.Println("area:", area)
}

func sphereStuff(radius int) (d int, c, s float32) {
	d = radius * 2
	c = 2 * math.Pi * float32(radius)
	s = 4 * math.Pi * float32(radius) * float32(radius)
	return
}
