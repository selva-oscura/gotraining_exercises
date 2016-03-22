// Declare three variables that are initalized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initalize the variable by
// converting the literal value of Pi (3.14).
package main

// add needed imports
import (
	"fmt"
	"math"
)

// main is the entry point for the application.
func main() {
	// Declare variables that are set to their zero value.
	var aFloat float64
	var aString string
	var anInt int
	var anotherInt int
	var aBool bool

	// Display the value of those variables.
	fmt.Println("\nDisplaying variables set to zero value:")
	fmt.Printf("var aFloat \t %T [%v]\n", aFloat, aFloat)
	fmt.Printf("var aString \t %T [%v]\n", aString, aString)
	fmt.Printf("var anInt \t %T [%v]\n", anInt, anInt)
	fmt.Printf("var anotherInt \t %T [%v]\n", anotherInt, anotherInt)
	fmt.Printf("var aBool \t %T [%v]\n", aBool, aBool)

	// Declare variables and initalize.
	// Using the short variable declaration operator.
	e := 2.718281828
	aBeginning := "Once upon a Time"
	meaningOfLife := 42
	goIsTyped := true

	// Display the value of those variables.
	fmt.Println("\nDisplaying variables declared and initialized:")
	fmt.Printf("var e \t %T [%v]\n", e, e)
	fmt.Printf("var aBeginning \t %T [%v]\n", aBeginning, aBeginning)
	fmt.Printf("var meaningOfLife \t %T [%v]\n", meaningOfLife, meaningOfLife)
	fmt.Printf("var goIsTyped \t %T [%v]\n", goIsTyped, goIsTyped)

	// Perform a type conversion.
	var pi float32
	pi = math.Pi

	// Display the value of that variable.
	fmt.Println("\nAnd only 8 days too late for Pi Day:")
	fmt.Printf("var pi \t %T [%v]\n\n", pi, pi)
}
