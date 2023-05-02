package main

import (
	"fmt"
)

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
}
