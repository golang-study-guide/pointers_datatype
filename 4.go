// go run 4.go
package main

import (
	"fmt"
	"reflect"
)

// https://medium.com/@kent.rancourt/go-pointers-when-to-use-pointers-4f29256ddff3
// https://dhdersch.github.io/golang/2016/01/23/golang-when-to-use-string-pointers.html
// http://piotrzurek.net/2013/09/20/pointers-in-go.html
// https://winterflower.github.io/2017/08/20/the-asterisk-and-the-ampersand/
func main() {

	// This carries on from the 2-box-concept from - https://github.com/Sher-Chowdhury/gsg_package_level_variables/blob/482a49b9c9e869056dda76d562fb9778c25e442d/11_the_pointer_datatype.go

	// pointers are created using the asterisk notation
	var nameAlpha *string
	// This ends up only creating the first box.
	// This box can hold the memory address of a second box that's designed to hold a
	// string variable.
	fmt.Println(reflect.TypeOf(nameAlpha)) // *string

	// But we haven't created the second box yet.
	// Hence the first box's content is blank.

	// at this stage this pointer is uninitialised. i.e. the first box is empty.
	fmt.Println(nameAlpha) // <nil>

	// Sidenote, in case you need to, you can get the first box's memory location:
	fmt.Println(&nameAlpha) // 0xc00000e028

	// To create the 2nd box, and also pupulate the 1st box, you can create your pointer like this:

	var nameBravo *string = new(string)
	// or the shorthand form:
	nameCharlie := new(string)

	// In these cases, both boxes are created, the first box stores the memory address
	// of the second box. But the second box is empty.
	fmt.Println(reflect.TypeOf(nameBravo))   // *string
	fmt.Println(reflect.TypeOf(nameCharlie)) // *string

	fmt.Println(*nameBravo)   // ''
	fmt.Println(*nameCharlie) // ''

	// The asterisk is how we tell golang you want use the content of the first box as the lookup
	// address to view the content of the second box. This technique is referred to as "deferencing the pointer"

	// So how do we populate the second box? You might of doing something like this:
	/*
		var nameDelta *string = "hello delta"
	*/
	// but that will fail because this nameDelta pointer can only store memory location
	// data (e.g. something like "0xc000010270"). not a string variable itself.

	// The correct way is by using the dereferencing technique
	*nameBravo = "My name is Bravo"
	fmt.Println(*nameBravo) // 'My name is Bravo'

	// 	you can also create pointers to point to existing variable locations
	nameEddie := "Murphy"
	fmt.Println(nameEddie)  // Murphy
	fmt.Println(&nameEddie) // 0xc000010230

	myPointer := &nameEddie
	// Here Go uses inference to work out that myPointer is a pointer datatype, since we are asking it to store
	// a memory location

	fmt.Println(reflect.TypeOf(myPointer)) // *string
	fmt.Println(*myPointer)                // Murphy

	// here we've essentially created a pointer without needing to use the asterisk notation. You can now
	// edit the content of the second box by either of the following

	nameEddie = "Jordan"
	*myPointer = "Jordan"

	// Irrespective of the approach you take, creating a pointer and populating it
	// requires a minimum of 2 lines of code.
	// 1. create the pointer
	// 2. populate the second box.
}
