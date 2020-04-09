// go run 6.go
package main

import (
	"fmt"
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

	kevin := User{
		Name: "Kevin Bacon",
		Age:  60,
	}

	kevin.IncAge()
	fmt.Println(kevin.Age) // 61
}

// this code is nearly identical to 5.go
// The only difference is that I added "*" to indicate that IncAge function
// now accepts a pointer of the type User. This tells IncAge to not make a copy of
// 'kevin' instead get the memory location of kevin and use that instead.
// Notice we didn't have to explicitly provide the IncAge method with a pointer,
// that's because golang is smart enough to work out what you're trying to do.
