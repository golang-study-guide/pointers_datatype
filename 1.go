// go run 1.go
package main

import (
	"fmt"
	"reflect"
)

func main() {

	// This short variable declaration (colon-equal) does 3 things.
	// 1. initilize the variable (create a box)
	// 2. infer data type - set the type of box it is (e.g. this box can only store strings)
	// 3. declares a variable (populate the box with the data)
	// https://blog.learngoprogramming.com/golang-short-variable-declaration-rules-6df88c881ee
	nameBravo := "Simon"

	// let's explore these 3 steps.

	// 1. initialising - this box's lid has a memory location address. To view it
	// you need to prefix the variable with ampersand "&":
	fmt.Println(&nameBravo) // 0xc0000721e0

	// 2. Infer - To find the type of box this is, i.e. data type:
	fmt.Println(reflect.TypeOf(nameBravo)) // string

	// 3. Declare - To view the box's content:
	fmt.Println(nameBravo) // Simon
	// you can think of the variable name, "nameBravo" as an entry in the index
	// in the back of a book, where instead of a page number, it gives the 
	// memory location address, e.g. something like 0xc0000721e0

}
