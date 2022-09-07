/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Sun Jun 05 2022
 */

// Runtime Error
// REFERENCE: https://leetcode.com/problems/decoded-string-at-index/

package main

import (
	"fmt"
)

func decodeAtIndex(s string, k int) string {

	strLen := len(s)
	decodedStringLen := 0
	index := 0
	decodedString := make([]string, 1)

	decodedStringAtIndexK := ""

	for i := 0; i < strLen; i++ {

		stringValue := string(s[i])
		intValue := int(s[i]) - 48

		if (intValue >= 2) && (intValue <= 9) {
			for j := 0; j < intValue-1; j++ {

				currentIndex := index

				for m := 0; m <= currentIndex; m++ {

					index++
					currentDecodedStrLen := decodedStringLen
					decodedString = append(decodedString, "")
					decodedString[index] += decodedString[m]
					decodedStringLen += len(decodedString[m])

					if decodedStringLen >= k {
						stringIndex := k - currentDecodedStrLen
						decodedStringAtIndexK = string(decodedString[m][stringIndex])
						return decodedStringAtIndexK
					}
				}
			}
		} else {
			decodedString[index] += stringValue
			decodedStringLen++
		}

		if decodedStringLen >= k {
			decodedStringAtIndexK = stringValue
			return decodedStringAtIndexK
		}
	}

	return decodedStringAtIndexK
}

func main() {
	var str string
	var k int
	for {
		fmt.Scanln(&str)
		if str == "0" {
			break
		}
		fmt.Scanln(&k)

		fmt.Println("output : ", decodeAtIndex(str, k))
	}
}

// "leet2code3"
// 10
// "o"
// "ha22"
// 5
// "h"
// "a2345678999999999999999"
// 1
// "a"
// "y959q969u3hb22odq595"
// 222280369
// "y"
