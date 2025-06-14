// Гриценко Ірина Валеріївна №1
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	n := r.Intn(91) + 10
	fmt.Println("\nn:", n, "\n")

	fmt.Println(arrA(n), "\n")
	fmt.Println(arrB(n), "\n")
	fmt.Println(arrV(n), "\n")
	fmt.Println(arrG(n), "\n")
	fmt.Println(arrD(n), "\n")
	fmt.Println(arrE(n), "\n")
	fmt.Println(arrYE(n), "\n")
	fmt.Println(arrJ(n), "\n")
	fmt.Println(arrZ(n), "\n")
	fmt.Println(arrI(n), "\n")
	fmt.Println(arrY(n), "\n")
	fmt.Println(arrK(n), "\n")

}

func arrA(n int) [100]float32 {
	fmt.Println("МАСИВ А \n")
	var a [100]float32
	for i := 0; i < n; i++ {
		a[i] = float32(r.Intn(10)+10) + r.Float32()
	}
	return a
}

func arrB(n int) [100]float32 {
	fmt.Println("МАСИВ Б \n")
	var a [100]float32
	for i := 0; i < n; i++ {
		var t float32 = 20.3
		for t >= 17 && t < 43 {
			t = float32(r.Intn(89)-35) + r.Float32()
		}
		a[i] = t
	}
	return a
}

func arrV(n int) [100]float32 {
	fmt.Println("МАСИВ В \n")
	a := (r.Intn(150) - 75)
	fmt.Println("a = ", a)
	b := (r.Intn(74-a) + a + 1)
	fmt.Println("b = ", b)
	var x [100]float32
	for i := 0; i < n; i++ {
		x[i] = float32(r.Intn(b-a)+a) + r.Float32()
	}
	return x
}

func arrG(n int) [100]uint8 {
	fmt.Println("МАСИВ Г \n")
	var a [100]uint8
	for i := 0; i < n; i++ {
		a[i] = uint8(r.Int31n(21)) + 20
	}
	return a
}

func arrD(n int) [100]int8 {
	fmt.Println("МАСИВ Д \n")
	a := (r.Intn(40) - 20)
	fmt.Println("a = ", a)
	b := (r.Intn(19-a) + a + 1)
	fmt.Println("b = ", b)
	var x [100]int8
	for i := 0; i < n; i++ {
		x[i] = int8(r.Intn(b-a+1) + a)
	}
	return x
}

func arrE(n int) [100]int8 {
	fmt.Println("МАСИВ Е \n")
	var a [100]int8
	for i := 0; i < n; i++ {
		t := 1
		for t%4 != 0 {
			t = r.Intn(100) + 1
		}
		a[i] = int8(t)
	}
	return a
}

func arrYE(n int) [100]int16 {
	fmt.Println("МАСИВ Є \n")
	var a [100]int16
	for i := 0; i < n; {

		t := r.Intn(999) + 1
		x := t % 10
		y := t / 10 % 10
		z := t / 100

		if (2*x)%10 == 0 && (2*y)%10 == 0 && (2*z)%10 == 0 {
			a[i] = int16(t)
			i++
		}
	}
	return a
}

func arrJ(n int) [100]int16 {
	fmt.Println("МАСИВ Ж \n")
	var arr [100]int16
	a, b, c := r.Intn(10), r.Intn(10), r.Intn(10)
	fmt.Println("a:", a, "  b:", b, "  c:", c)
	for i := 0; i < n; {

		t := r.Intn(499) + 1

		x := t % 10
		y := t / 10 % 10
		z := t / 100

		condition := (x == a || x == b || x == c) && (y == a || y == b || y == c) && (z == a || z == b || z == c)

		if condition {
			arr[i] = int16(t)
			i++
		}
	}
	return arr
}

func arrZ(n int) [100]uint8 {
	fmt.Println("МАСИВ З \n")
	var a [100]uint8
	for i := 0; i < n; i++ {
		a[i] = uint8(r.Int31n(2))
	}
	return a
}

func arrI(n int) [100]float32 {
	fmt.Println("МАСИВ І \n")
	values := [...]float32{0, 0.2, 0.4, 0.6, 0.8, 1.0}
	var arr [100]float32
	for i := 0; i < n; i++ {
		randIdx := r.Intn(len(values))
		arr[i] = values[randIdx]
	}
	return arr
}

func arrY(n int) [100]int {
	fmt.Println("МАСИВ Й \n")
	var arr [100]int
	for i := 0; i < n; {
		t := r.Intn(999) + 1
		x := t % 10
		y := t / 10 % 10
		z := t / 100

		if (x + y + z) == 10 {
			arr[i] = t
			i++
		}
	}
	return arr
}

func arrK(n int) [100]int {
	fmt.Println("МАСИВ К \n")
	sumOfDid := r.Intn(40) + 1
	fmt.Println("sumOfDid = ", sumOfDid)
	var arr [100]int
	for i := 0; i < n; {
		t := r.Intn(99999) + 1
		x := t % 10
		y := t / 10 % 10
		z := t / 100 % 10
		q := t / 1000 % 10
		u := t / 10000

		if (x + y + z + q + u) == sumOfDid {
			arr[i] = t
			i++
		}
	}
	return arr
}
