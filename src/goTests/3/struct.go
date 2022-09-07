/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Sat May 21 2022
 */

package main

import "fmt"

type Contact struct {
	name string
	age  int
}

func (x Contact) welcome() Contact {
	x.age += 1
	fmt.Println("Welcome : ", x.name, " ", x.age)
	return x
}

func main() {
	// contactOfJames := Contact{"James", 42}
	contact0 := Contact{name: "James", age: 42}

	fmt.Println("First contact: ", contact0)

	fmt.Println("name of contact: ", contact0.name)

	contact0.age += 1

	fmt.Println("age of ", contact0.name, " now : ", contact0.age)

	p := &contact0

	// Pointers to structs are automatically dereferenced, meaning we can access the field values by simply using a dot:
	fmt.Println("name of contact from pointer: ", p.name)

	// Now newP is a pointer to the newly created struct
	newP := &Contact{"Bob", 45}

	fmt.Println("name of new contact: ", newP.name)

	// Methods are simply functions with a special receiver argument.
	newP.welcome()
	// won't increase the original value here
	newP.welcome()

	// can't call similarly on pointer

	for i := 0; i < 5; i++ {
		// will increase the original value here
		contact0 = contact0.welcome()
	}
}
