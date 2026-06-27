// Understanding Variables Declarations in Go
package main

import "fmt"

func main() {
	// way1: Declare and assign on 2 different lines
	var mango string = "This is a big mango!"
	var weight int = 54

	// way2: Declare and assign on the same line
	var height int = 23

	fmt.Println("Mango:", mango)
	fmt.Println("weight:", weight)
	fmt.Println("height:", height)

	// way3 (shorthand)
	// shorthand notation :=
	// Type inference
	// var age = 23
	age := 54
	city := "Washington"
	fmt.Println("My age is:", age)
	fmt.Println("My city is:", city)

	// Multiple declaration and init. on same line
	var apples, oranges int = 23, 78
	fmt.Println("I have", apples, "apples and", oranges, "oranges")

	var fruits = apples + oranges
	fmt.Println("fruits:", fruits)

	var windows, mac, linux string = "Windows is ok\n", "Mac is meh\n", "Linux is the GOAT!\n"
	print(windows, mac, linux)

	// Static Type declaration
	var x float64 = 20.5
	fmt.Println(x)
	fmt.Printf("x is of type: %T\n", x)

	// Dynamic Type declaration
	y := 89
	fmt.Println(y)
	fmt.Printf("y is of type: %T\n", y)

	// Mixed Variable Declaration
	var a, b, c = 758.52, 8, "foobar"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a is of type: %T\n", a)
	fmt.Printf("b is of type: %T\n", b)
	fmt.Printf("c is of type: %T\n", c)

	// exported and unexported variables
	// exported variables start with a capital letter and can be accessed from other packages
	// unexported variables start with a lowercase letter and can only be accessed within the same package
	var ExportedVariable string = "I am exported"
	var unexportedVariable string = "I am unexported"
	fmt.Println(ExportedVariable)
	fmt.Println(unexportedVariable)

	// grouped declartion of variables
	var (
		// exported variable
		ExportedVariable2 string = "I am exported"
		// unexported variable
		unexportedVariable2 string = "I am unexported"
	)
	fmt.Println(ExportedVariable2)
	fmt.Println(unexportedVariable2)

	// type conversion
	// The format that we follow to convert a variable from one type to another is: <newType>(<variable>)
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Printf("i is of type: %T and value: %v\n", i, i)
	fmt.Printf("f is of type: %T and value: %v\n", f, f)
	fmt.Printf("u is of type: %T and value: %v\n", u, u)

	// scope of variables
	// The scope of a variable is the region of the program where the variable is defined and can be accessed.
	//  In Go, variables can have either package-level scope or function-level scope.
	// Package-level scope: A variable declared outside of any function has package-level scope and can be accessed from any function within the same package.
	// The short variable declaration (:=) can only be used within a function and cannot be used at the package level.
	// Function-level scope: A variable declared inside a function has function-level scope and can only be accessed within that function.

	var packageLevelVariable string = "I am a package-level variable"

	func() {
		var functionLevelVariable string = "I am a function-level variable"
		fmt.Println(packageLevelVariable)
		fmt.Println(functionLevelVariable)
	}()

	// Uncommenting the following line will cause an error because functionLevelVariable is not accessible outside its function
	// fmt.Println(functionLevelVariable)

	// variable shadowing
	// Variable shadowing occurs when a variable declared within a
	// certain scope (e.g., a function) has the same name as a variable declared in an outer scope
	// (e.g., package-level). In such cases, the inner variable "shadows" the outer variable, making the outer variable inaccessible within the inner scope.

	var shadowedVariable string = "I am the outer variable"

	func() {
		var shadowedVariable string = "I am the inner variable"
		fmt.Println(shadowedVariable) // This will print: I am the inner variable
	}()
	fmt.Println(shadowedVariable) // This will print: I am the outer variable

	// blank identifier
	// The blank identifier is a special identifier that is represented by an underscore (_).
	// It is used to discard values that are not needed.

	_, _, _ = 1, 2, 3 // Discard all three values

	// func returnMultipleValues() (int, int, int) {
	// 	return 1, 2, 3
	// }

	// a, _, c := returnMultipleValues() // Discard the second value
	// fmt.Println("a:", a)
	// fmt.Println("c:", c)

}
