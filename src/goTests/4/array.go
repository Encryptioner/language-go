/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Sat May 21 2022
 */

package main

import "fmt"

func main() {
	// var arr [5]int
	arr := [5]int{0, 2, 4, 6, 8}

	fmt.Println("Printing arr :")
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, " ", arr[i])
	}

	// var slicedArray []int = arr[1:3]
	// var slicedArray []int = arr[:3]
	var slicedArray []int = arr[1:]

	fmt.Println("Printing slicedArray :")
	for i := 0; i < len((slicedArray)); i++ {
		fmt.Println(i, " ", slicedArray[i])
	}

	// A slice does not store any data; it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.

	arr[2] = 10

	fmt.Println("Printing slicedArray after arr[2] :")

	for i, v := range slicedArray {
		fmt.Println(i, " ", v)
	}

	// Go provides a make() function to create slices.
	// This is how you create dynamically-sized arrays

	// dynamicArr := make([]int, 0)
	dynamicArr := make([]int, 2)
	dynamicArr = append(dynamicArr, 1, 2)

	fmt.Println("Printing dynamicArr :")

	sum := 0
	// If you want only the values, you can skip the index using an underscore:
	for _, v := range dynamicArr {
		if v%2 == 0 {
			fmt.Println("value is even")
		}
		fmt.Println(v)
		sum += v
	}

	// printing string

	helloString := "hello"

	fmt.Println("Printing helloString: ")
	for _, v := range helloString {
		fmt.Printf("character %c \n", v)
		fmt.Println("unicode code points: ", v)
	}

	var bool0 [5]bool

	println("bool0[3] = ", bool0[3])

	fmt.Scanln(&bool0[3])

	println("bool0[3] = ", bool0[3])
}
