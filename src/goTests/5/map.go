/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Sat May 21 2022
 */

package main

import (
	"fmt"
	"reflect"
)

func main() {
	map0 := map[int]int{
		8: 42,
		2: 6,
		4: 9,
		5: 3,
	}
	fmt.Println("map0[4] - map0[5] = ", map0[4]-map0[5])

	delete(map0, 2)
	fmt.Println("map0[2] = ", map0[2])

	keys := []int{}
	for key, _ := range map0 {
		keys = append(keys, key)
	}
	fmt.Println("map0 keys slice = ", keys)

	keysUsingPackage := reflect.ValueOf(map0).MapKeys()
	fmt.Println("map0 keys slice using package = ", keysUsingPackage, keysUsingPackage[0])

	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}

	var str string
	fmt.Scanln(&str)

	results = append(results, str)

	resultMap := map[string]int{
		"w": 3,
		"d": 1,
		"l": 0,
	}

	total := 0

	for _, v := range results {
		total += resultMap[v]
	}
	fmt.Println(total)
}
