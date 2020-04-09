// go run 2.go
package main

import (
	"fmt"
)

func main() {
	// So far we saw
	nameBravo := "Simon"
	fmt.Println(&nameBravo) // 0xc0000721e0

	// You can also do this:
	fmt.Println(*&nameBravo) // Simon
	// This asterisk prefix tells go to treat "&nameBravo" as a memory address, and
	// get the box's content. Hence in a round-about way this equivalent to simply just:

	fmt.Println(nameBravo) // Simon
	// Where "nameBravo" is just a nickname to the memory location.

}
