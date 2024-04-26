package main

import "fmt"

func main() {

	// Array - tamanho fixo
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
//	fmt.Println(array[0])
//	fmt.Println(array[1])
//	fmt.Println(array[0], array[1])
//	fmt.Println(array)

// numPrimos := [6] int{2, 3, 5, 7, 11, 13}
// fmt.Println(numPrimos)
// fmt.Println(numPrimos[0:4])
// fmt.Println(numPrimos[0:6])

// Slice
// var slice []string
slice := make([]string, 5)
slice[0] = "Hello"
slice[1] = "World"
// fmt.Println(slice[0], slice[1])
//fmt.Println(slice[0])
// fmt.Println(slice[1])
// fmt.Println(slice[2])
// slice[2] = "fabricio"
// fmt.Println(slice[2])
// fmt.Println(slice)

// fmt.Println(len(slice))
// fmt.Println(slice[2])
// fmt.Println(slice[3])
// fmt.Println(slice[4])
// fmt.Println(slice[5])

numPares := []int{2, 4, 6, 8, 10, 12}
fmt.Println(numPares)

numPares = append(numPares, 14, 16, 18, 20)
fmt.Println(numPares)

}