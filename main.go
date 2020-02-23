package main

import "fmt"
// https://medium.com/@kent.rancourt/go-pointers-when-to-use-pointers-4f29256ddff3
// https://dhdersch.github.io/golang/2016/01/23/golang-when-to-use-string-pointers.html
// http://piotrzurek.net/2013/09/20/pointers-in-go.html
// https://winterflower.github.io/2017/08/20/the-asterisk-and-the-ampersand/
func main() {

	// pointers are created using the asterisk notation
	var nameAlpha *string

	// at this stage this pointer is uninitialised.
	fmt.Println(nameAlpha) // this should output 'nil', this is go's way of saying the pointer isn't point to anything

	// since it is unitialised it means we can't story any data at the pointer's memory location. hence the following will fail:
	// *nameAlpha = Peter
	// Note we use the asterisk notation to indicate that we're working with a pointer datatype's memory location

	// This time we initialise it, which in earmarks some memory
	var nameBravo *string = new(string)
	// this outputs the memory location, e.g. something like 0x53x10
	fmt.Println(nameBravo)

	// to get the value stored at this memory location we do
	// the following, at this point it will print a blank line
	// since that location is empty.
	fmt.Println(*nameBravo)

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
