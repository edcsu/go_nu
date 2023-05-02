package main

import "fmt"

func main() {

	var firstName string = "Jane"
	var secondName = "Hadidi"
	var thirdName string

	fmt.Println(firstName, secondName, thirdName)

	thirdName = "Panenka"

	fmt.Println(firstName, secondName, thirdName)

	// cannot be used outside of a function
	age := 22
	name := "Doe"

	var anotherAge int = 20
	otherAge := 21

	fmt.Println(anotherAge, otherAge)

	// Print
	fmt.Print("Hello, ")
	fmt.Print("world! \n")

	fmt.Println("Hello, world!")

	fmt.Println("My age is", age, "and my name is", name)

	// Formatted string
	// %_ => format specifer
	fmt.Printf("My age is %v and my name is %v \n", age, name)
	fmt.Printf("My age is %v and my name is %q \n", age, name)
	fmt.Printf("My age is of type %T \n", age)
	fmt.Printf("This float is %f \n", 209.78)
	fmt.Printf("This float is %0.1f \n", 209.78)
	var savedString = fmt.Sprintf("My age is %v and my name is %q \n", age, name)
	fmt.Printf("The saved string is: %v", savedString)
}
