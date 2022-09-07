/*
 * Ankur Mursalin
 *
 * https://encryptioner.github.io/
 *
 * Created on Wed Jul 27 2022
 */

// REFERENCE: https://leetcode.com/problems/flatten-binary-tree-to-linked-list/

package main

import (
	"fmt"
	"reflect"
)

// NOTE: To loop over struct:
// https://go.dev/play/p/mjZLpD63Edf
// https://go.dev/play/p/nleA2YWMj8

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var output TreeNode

func InspectStructV(val reflect.Value, values *[]int) {
	if val.Kind() == reflect.Interface && !val.IsNil() {
		elm := val.Elem()
		if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
			val = elm
		}
	}
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		address := "not-addressable"

		if valueField.Kind() == reflect.Interface && !valueField.IsNil() {
			elm := valueField.Elem()
			if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
				valueField = elm
			}
		}
		if valueField.Kind() == reflect.Ptr {
			valueField = valueField.Elem()
		}
		if valueField.CanAddr() {
			address = fmt.Sprintf("0x%X", valueField.Addr().Pointer())
		}

		fmt.Printf(
			"Field Name: %s,\t Field Value: %v,\t Address: %v\t, Field type: %v\t, Field kind: %v\n",
			typeField.Name, valueField, address, typeField.Type, valueField.Kind(),
		)

		if valueField.Kind() == reflect.Struct {
			InspectStructV(valueField, values)
		} else if valueField.Kind() == reflect.Int {
			intValue := int(valueField.Int())
			*values = append(*values, intValue)
		}
	}
}

func flatten(root *TreeNode) {

	values := make([]int, 0)

	InspectStructV(reflect.ValueOf(*root), &values)

	fmt.Println(values)

	var helper func(int)

	var output TreeNode

	helper = func(index int) {
		if index == len(values)-1 {
			output.Right = nil
			output.Left = nil
			output.Val = values[index]
			fmt.Println(output)
			return
		}

		helper(index + 1)

		fmt.Println(index, output)
		output.Right = &output
		output.Left = nil
		output.Val = values[index]
	}

	helper(0)

	fmt.Println("printing")

	fmt.Println(output)

	fmt.Println(output.Right)
	fmt.Println(output.Right.Right)

}

func main() {

	var root TreeNode

	root = TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			// Right: &TreeNode{
			// 	Val: 4,
			// },
		},
		// Right: &TreeNode{
		// 	Val: 5,
		// 	Right: &TreeNode{
		// 		Val: 6,
		// 	},
		// },
	}

	flatten(&root)
}
