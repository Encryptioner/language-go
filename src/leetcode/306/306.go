/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Sun May 22 2022
 */

// incomplete
// REFERENCE: https://leetcode.com/problems/additive-number/

package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func getIntDigitArrFromDigitString(str string) []int {
	strDigitArr := strings.Split(str, "")

	intDigitArr := make([]int, 0)

	for _, v := range strDigitArr {
		digit, _ := strconv.Atoi(v)
		intDigitArr = append(intDigitArr, digit)
	}
	return intDigitArr
}

func isAdditiveNumber(num string) bool {
	intDigitArr := getIntDigitArrFromDigitString(num)

	// TODO: check if num is additive or not

	fmt.Println(intDigitArr)
	return true
}

func main() {
	var num string

	for {
		fmt.Scanln(&num)

		if num == "0" {
			fmt.Println(reflect.TypeOf(num))
			break
		}

		fmt.Println("IS ADDITIVE : ", isAdditiveNumber(num))
	}
}
