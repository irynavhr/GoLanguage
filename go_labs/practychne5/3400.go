// Гриценко Ірина Валеріївна 08.11.2023
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	M = 4
	N = 3
)

func main() {

	array := makeArray()
	fmt.Println(array)
	sums_array := sumsByCols(array)
	fmt.Println(sums_array)
	k := maxCol(sums_array)
	//fmt.Println(k)
	min := array[0][k]
	for i := 1; i < M; i++ {
		if array[i][k] < min {
			min = array[i][k]
		}
	}
	fmt.Println("Найменший елемент у стовпчику з найбільшою сумою: ", min)
}

func maxCol(arr [N]uint8) uint8 {
	var k uint8
	max := arr[0]
	for i := 1; i < N; i++ {
		if arr[i] > max {
			max = arr[i]
			k = uint8(i)
		}
	}
	//fmt.Println(max)
	return k
}

func sumsByCols(arr [M][N]uint8) (res [N]uint8) {

	for row := 0; row < M; row++ {
		for col := 0; col < N; col++ {
			res[col] += arr[row][col]
		}
	}
	return
}

func makeArray() (arr [M][N]uint8) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for row := 0; row < M; row++ {
		for col := 0; col < N; col++ {
			arr[row][col] = uint8(r.Intn(9) + 1)
		}
	}
	return
}
