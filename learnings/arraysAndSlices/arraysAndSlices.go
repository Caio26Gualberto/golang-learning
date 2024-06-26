package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array3)

	slice1 := []string{"Caio", "João", "André", "Marrone", "JV", "MatheusBrazil"}
	fmt.Println(slice1)

	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array3))

	slice1 = append(slice1, "Lorenzo")
	fmt.Println(slice1)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(slice2)

	// Interns Arrays

	fmt.Println("----- Interns Arrays -----")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}
