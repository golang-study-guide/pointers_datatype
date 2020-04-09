// go run 4.go
package main

import (
	"fmt"
	"reflect"
)

// User is just a simple struct
type User struct {
	Name string
	Age  int
}

// IncAge increases the User's age
func (u User) IncAge() {
	u.Age = u.Age + 1
	fmt.Println(u.Age) // 61
}

func main() {

	// We've looked at what are pointers. But why use pointers?
	// There's several reasons, the main one being memory efficiency.

	// Here's an example

	kevin := User{
		Name: "Kevin Bacon",
		Age:  60,
	}
	fmt.Println(reflect.TypeOf(kevin)) // main.User

	kevin.IncAge()

	fmt.Println(kevin.Age) // 60
	// This remain unchanged. That's because a copy of "kevin" was created for the
	// IncAge to use. Hence, to instances of 'kevin' datatype is created during the run.
	// hence double the memory.
	// if you only want one instance and you want it to be modified by your function, then
	// you need to do it using pointers. See 6.go for how the pointers version
	// of this code looks.
}
