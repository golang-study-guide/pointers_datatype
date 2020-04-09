package main

import (
	"fmt"
)

func main() {

	// The "new" builtin function creates a box that holds the
	// specified datatype, and returns the memory location of that box.
	// (i.e. the name written on the box lid)
	fmt.Println(new(string)) // 0xc000010200
	fmt.Println(new(string)) // 0xc000010210
	fmt.Println(new(string)) // 0xc000010220
	fmt.Println(new(string)) // 0xc000010230
	fmt.Println(new(string)) // 0xc000010240

	// In the above example we're just printing the memory locations out to
	// the standard output. So can't really do much with them like this.
}
