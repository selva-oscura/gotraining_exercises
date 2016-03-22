// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initalize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.

package main

// Add imports.

import "fmt"

// Add user type and provide comment.
type userStructLiteral struct {
	name	string
	email	string
	age		int
}

// main is the entry point for the application.
func main() {
	// Declare variable of type user and init using a struct literal.
	user1 := userStructLiteral{
		name:	"user mcnewby",
		email:	"user@mchnewby.org",
		age:	42,
	}

	// Display the field values.

	fmt.Printf("%+v\n", user1)
	fmt.Println("name", user1.name)
	fmt.Println("email", user1.email) 
	fmt.Println("age", user1.age)
 
	// Declare a variable using an anonymous struct.

    userAnonStruct := struct {
        name    string
        email   string
        age     int
    }{
        name:   "other newby",
        email:  "other@newby.org",
        age:    21,
    }

    // Display the field values.

    fmt.Printf("%+v\n", userAnonStruct)
    fmt.Println("name", userAnonStruct.name)
    fmt.Println("email", userAnonStruct.email)
    fmt.Println("age", userAnonStruct.age)

}
