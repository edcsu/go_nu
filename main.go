package main

import (
	"fmt"
)

// non-pointer types -> strings, ints, bools, floats, arrays, structs
// pointer types -> maps, slices, functions

func updateLocation(x *string) {
	*x = "updated"
}

// go run main.go greetings.go bye.go cyclenames.go circlearea.go getinitials.go

func main() {
	SayGreeting("Jane")

	SayBye("Doe")

	cycleNames([]string{"Messi", "Mbappe", "Neymar"}, SayGreeting)
	cycleNames([]string{"Messi", "Mbappe", "Neymar"}, SayBye)

	a1 := circleArea(5)
	fmt.Println(a1)
	fmt.Printf("Area for circle is: %0.4f \n", a1)

	fn1, sn1 := getInitials("jane doe")
	fmt.Println(fn1, sn1)

	// maps => dictionary in python or objects in JS
	menu := map[string]float64{
		"coffee":  5000.99,
		"tea":     2500.99,
		"samosa":  500.99,
		"colesaw": 2000.99,
	}

	fmt.Println(menu)
	fmt.Println(menu["samosa"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// int as key type

	phonebook := map[int]string{
		256772123456: "Jane",
		256702123456: "Pendo",
		256712123456: "Theron",
		256773123456: "Fonda",
		256774123456: "Alyson",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[256773123456])

	phonebook[256773123456] = "Anita"
	fmt.Println(phonebook)

	location := "Kampala"

	fmt.Println("Memory address of location: ", &location)

	memoryLocation := &location
	fmt.Println("Memory address of location: ", memoryLocation)
	fmt.Println("value of Memory address of location: ", *memoryLocation)

	updateLocation(memoryLocation)

	fmt.Println(location)
}
