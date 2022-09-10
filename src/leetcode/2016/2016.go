/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Sat Sep 10 2022
 */

// Accepted
// REFERENCE: https://leetcode.com/problems/maximum-difference-between-increasing-elements/

package main

func maximumDifference(nums []int) int {
	maxDiffVal := -1
	lowestIVal := nums[0]

	for _, num := range nums {
		diffVal := num - lowestIVal
		if diffVal > maxDiffVal && diffVal > 0 {
			maxDiffVal = diffVal
		}
		if num < lowestIVal {
			lowestIVal = num
		}
	}
	return maxDiffVal
}

func main() {
	arr := [5]int{1, 5, 3, 6, 4}
	maximumDifference(arr[:])
}

// Test
// [7,1,5,4]
// [9,4,3,2]
// [1,5,2,10]
