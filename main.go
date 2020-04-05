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

    fmt.Println(new(string))  // 0xc000010200
    fmt.Println(new(string))  // 0xc000010210
    fmt.Println(new(string))  // 0xc000010220
    fmt.Println(new(string))  // 0xc000010230
    fmt.Println(new(string))  // 0xc000010240


	// pointers are created using the asterisk notation
	var nameAlpha *string
	// This only ends up creating the first box. Hence the first box's content is blank. so you cant follow the memory
	// locaiton using the *nameAplha notation. 

	// at this stage this pointer is uninitialised.
	fmt.Println(nameAlpha) // this should output 'nil', this is go's way of saying the pointer isn't point to anything

	// since it is unitialised it means we can't story any data at the pointer's memory location. 
	// hence the following will fail:
	// *nameAlpha = Peter
	// Note we use the asterisk notation to say we're working with the 2nd box's content. 

	// This time we initialise it, which in earmarks some memory
	var nameBravo *string = new(string)
	
	fmt.Println(nameBravo)   // 0xc000010240

	// this is how you tell golang you want use the content of the first box as the lookup 
	// address to view the content of the second box. 
	fmt.Println(*nameBravo)

	// now we can store content inside the second box. 
	*nameBravo = "Simon"
	fmt.Println(*nameBravo)

	// you can also create pointers to point to existing variable locations
	nameCharlie := "Chaplin"
	fmt.Println(nameCharlie)

	// here we use ampersand to get a variable's memory location
	myPointer := &nameCharlie
	fmt.Println("mypointer datatype is: ", reflect.TypeOf(myPointer))  // mypointer datatype is:  *string
	fmt.Println("*mypointer datatype is: ", reflect.TypeOf(*myPointer)) // *mypointer datatype is:  string  

	

	// here we've essentially created a pointer without needing to use the asterisk
	fmt.Println(myPointer)  // 0xc000010270
	fmt.Println(*myPointer) // Chaplin

	// Now lets change the content at that memory location
	nameCharlie = "Sheen" // this changes the content of the second box
	fmt.Println(myPointer)  //  0xc000010270
	fmt.Println(*myPointer) // this shows the new content.
	*myPointer = "chaplin"
	fmt.Println(nameCharlie) // chaplin

	// SO in effect, there's 2 ways to change the second box's content, by updating nameCharlie or *myPointer
}
