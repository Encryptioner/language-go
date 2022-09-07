/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Fri May 20 2022
 */

package main

import "fmt"

func main() {
	x := 42
	p := &x

	fmt.Println("Address of x in pointer variable p: ", p)
	fmt.Println("value of x from pointer variable p: ", *p)

	*p = *p - 8
	fmt.Println("Value of x is updated from pointer variable p: ", *p)
}
