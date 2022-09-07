/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Sat May 21 2022
 */

// REFERENCE: https://www.sololearn.com/learning/1164/4883/12540/1

package main

import (
	"fmt"
)

// Variadic functions are functions that can be called with any number of arguments.

func sum(str string, nums ...float32) {
	var total float32

	println(str)
	total = 0.0
	for _, v := range nums {
		total += v
	}
	fmt.Println("Total of ", nums, " : ", total)
}

func main() {
	sum("Sum : ", 1, 2)

	sum("Sum : ", 100, -85)

	sum("Sum : ", -8.8845, -0.1111111115)

	arrSliced := []float32{1, 2, 10.25}
	sum("Sum of Arr: ", arrSliced...)
}
