// go run 7.go
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
func (u *User) IncAge() {
	u.Age = u.Age + 1
	fmt.Println(u.Age) // 61
}

// https://linuxacademy.com/blog/linux/demystifying-pointers-in-go/
func main() {

	// Here's how you create a struct based pointer in golang.
	// all you have to do is just add an ampersand!
	kevin := &User{
		Name: "Kevin Bacon",
		Age:  60,
	}
	fmt.Println(reflect.TypeOf(kevin)) // *main.User

	kevin.IncAge()
	fmt.Println(kevin.Age) // 61
}

// this is the same as 6.go, but this time we're actually feeding a *User pointer
// into the function. This will make the code run a fraction of a second faster,
// since golang has to do less working out.
