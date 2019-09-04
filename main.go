package main

import "fmt"

func main() {

	// pointers are created using the asterisk notation
	var nameAlpha *string

	// at this stage this pointer is uninitialised.
	fmt.Println(nameAlpha) // this should output 'nil', this is go's way of saying the pointer isn't point to anything

	// singer it is unitialised it means we can't story any data at the pointer's memory location. hence the following will fail:
	// *nameAlpha = Peter
	// Note we use the asterisk notation to indicate that we're working with a pointer datatype

	// This time we initialise it, which in earmarks some memory
	var nameBravo *string = new(string)
	// this outputs the memory location, e.g. something like 0x53x10
	fmt.Println(nameBravo)

	// now we can store a variable value at our pointer's memory locaitoni
	*nameBravo = "Simon"
	fmt.Println(*nameBravo)

	// you can also create pointers to point to existing variable locations
	nameCharlie := "Chaplin"
	fmt.Println(nameCharlie)

	// here we use ampersand to get a variable's memory location
	myPointer := &nameCharlie
	// here we've essentially created a pointer without needing to use the asterisk
	fmt.Println(myPointer)  // this outputs the memory's location
	fmt.Println(*myPointer) // this outputs the content at that memory locaiton

	// Now lets change the content at that memory location
	nameCharlie = "Sheen"
	fmt.Println(myPointer)  //  this stays the same.
	fmt.Println(*myPointer) // this shows the new content.

}
